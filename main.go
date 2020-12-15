package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

const (
	defaultDockerfile    = "Dockerfile"
	defaultDockerfileOut = "Dockerfile.out"
)

type manifestGlobal struct {
	Base      string                 `yaml:"base"`
	Doc       string                 `yaml:"doc"`
	Upstreams []string               `yaml:"upstreams"`
	Repos     []string               `yaml:"repos"`
	Vars      map[string]interface{} `yaml:"vars"`
	Mirrors   []string               `yaml:"mirrors"`
}

type manifestRepo struct {
	Name string                            `yaml:"string"`
	Desc string                            `yaml:"desc"`
	Tags []manifestTag                     `yaml:"tags"`
	Vars map[string]map[string]interface{} `yaml:"vars"`
}

type manifestTag struct {
	Name       string   `yaml:"name"`
	Also       []string `yaml:"also"`
	Dockerfile string   `yaml:"dockerfile"`
	Vars       []string `yaml:"vars"`
}

func exit(err *error) {
	if *err != nil {
		log.Println("exited with error:", (*err).Error())
		os.Exit(1)
	}
}

var (
	optOnly string
)

func main() {
	var err error
	defer exit(&err)

	flag.StringVar(&optOnly, "only", "", "只构建某个目录，用于调试")
	flag.Parse()

	var buf []byte
	if buf, err = ioutil.ReadFile("manifest.yml"); err != nil {
		return
	}

	var global manifestGlobal
	if err = yaml.Unmarshal(buf, &global); err != nil {
		return
	}

	if optOnly == "" {
		for _, upstream := range global.Upstreams {
			if err = execute("", "docker", "pull", upstream); err != nil {
				return
			}
		}
	}

	for _, dir := range global.Repos {
		if optOnly != "" && dir != optOnly {
			continue
		}
		log.Println("Dir:", dir)
		var repo manifestRepo
		if buf, err = ioutil.ReadFile(filepath.Join(dir, "manifest.yml")); err != nil {
			return
		}
		if err = yaml.Unmarshal(buf, &repo); err != nil {
			return
		}
		if repo.Name == "" {
			repo.Name = dir
		}
		log.Println("Repo:", repo.Name)

		for _, tag := range repo.Tags {
			if tag.Dockerfile == "" {
				tag.Dockerfile = defaultDockerfile
			}
			log.Println("Tag:", tag.Name)
			vars := map[string]interface{}{}
			for k, v := range global.Vars {
				vars[k] = v
			}
			for _, kg := range tag.Vars {
				if repo.Vars[kg] == nil {
					err = errors.New("missing vars group: " + kg)
					return
				}
				for k, v := range repo.Vars[kg] {
					vars[k] = v
				}
			}
			log.Printf("Vars: %v", vars)
			if err = build(optsBuild{
				doc:        global.Doc,
				dir:        dir,
				base:       global.Base,
				desc:       repo.Desc,
				repo:       repo.Name,
				tag:        tag.Name,
				also:       tag.Also,
				dockerfile: tag.Dockerfile,
				vars:       vars,
			}); err != nil {
				return
			}
		}
	}

	if optOnly != "" {
		return
	}

	for _, mirrorItem := range global.Mirrors {
		log.Println("Mirror:", mirrorItem)
		// decode
		mirrorSplits := strings.SplitN(mirrorItem, "=>", 2)
		if len(mirrorSplits) != 2 {
			err = fmt.Errorf("bad syntax for mirrors: %s", mirrorItem)
			return
		}
		src, dst := strings.TrimSpace(mirrorSplits[0]), strings.TrimSpace(mirrorSplits[1])
		if src == "" || dst == "" {
			err = fmt.Errorf("bad syntax for mirrors: %s", mirrorItem)
			return
		}

		// list tags
		var tags []string
		if tags, err = registryListTags(context.Background(), src); err != nil {
			return
		}

		for _, tag := range tags {
			srcImage := src + ":" + tag
			if err = execute("", "docker", "pull", srcImage); err != nil {
				return
			}
			dstImage := path.Join(global.Base, dst+":"+tag)
			if err = execute("", "docker", "tag", srcImage, dstImage); err != nil {
				return
			}
			if err = execute("", "docker", "push", dstImage); err != nil {
				return
			}
			if err = execute("", "docker", "rmi", dstImage); err != nil {
				return
			}
		}
	}
}

var tmplFuncs = template.FuncMap{}

type optsBuild struct {
	doc        string
	dir        string
	base       string
	repo       string
	desc       string
	tag        string
	also       []string
	dockerfile string
	vars       map[string]interface{}
}

func build(opts optsBuild) (err error) {
	if err = ioutil.WriteFile(
		filepath.Join(opts.dir, "banner.minit.txt"),
		[]byte(fmt.Sprintf(
			"本镜像基于 %s/%s:%s 制作，详细信息参阅 %s\n",
			opts.base,
			opts.repo,
			opts.tag,
			strings.ReplaceAll(opts.doc, "{{.Repo}}", opts.repo),
		)+opts.desc),
		0644); err != nil {
		return
	}
	var buf []byte
	if buf, err = ioutil.ReadFile(filepath.Join(opts.dir, opts.dockerfile)); err != nil {
		return
	}
	var tmpl *template.Template
	if tmpl, err = template.New("__main__").Option("missingkey=zero").Funcs(tmplFuncs).Parse(string(buf)); err != nil {
		return
	}
	out := &bytes.Buffer{}
	if err = tmpl.Execute(out, opts.vars); err != nil {
		return
	}
	out.WriteString("\nADD banner.minit.txt /etc/banner.minit.txt")
	if err = ioutil.WriteFile(filepath.Join(opts.dir, defaultDockerfileOut), sanitize(out.Bytes()), 0640); err != nil {
		return
	}
	canonicalName := fmt.Sprintf("%s/%s:%s", opts.base, opts.repo, opts.tag)
	log.Println("Build:", canonicalName)
	if err = execute(opts.dir, "docker", "build", "-t", canonicalName, "-f", defaultDockerfileOut, "."); err != nil {
		return
	}
	for _, alt := range opts.also {
		altCanonicalName := fmt.Sprintf("%s/%s:%s", opts.base, opts.repo, alt)
		log.Println("Tag:", canonicalName, altCanonicalName)
		if err = execute(opts.dir, "docker", "tag", canonicalName, altCanonicalName); err != nil {
			return
		}
	}
	if optOnly != "" {
		return
	}
	log.Println("Push:", canonicalName)
	if err = execute(opts.dir, "docker", "push", canonicalName); err != nil {
		return
	}
	for _, alt := range opts.also {
		altCanonicalName := fmt.Sprintf("%s/%s:%s", opts.base, opts.repo, alt)
		log.Println("Push:", canonicalName, altCanonicalName)
		if err = execute(opts.dir, "docker", "push", altCanonicalName); err != nil {
			return
		}
	}
	return
}

func sanitize(buf []byte) []byte {
	lines := bytes.Split(buf, []byte{'\n'})
	out := make([][]byte, 0, len(lines))
	for _, line := range lines {
		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		out = append(out, line)
	}
	return bytes.Join(out, []byte{'\n'})
}

type registryListTagResponse struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

type dockerHubTagItem struct {
	Layer string `json:"layer"`
	Name  string `json:"name"`
}

func registryListTags(ctx context.Context, repo string) (tags []string, err error) {
	comps := strings.Split(repo, "/")
	if len(comps) < 2 {
		err = fmt.Errorf("bad format for docker repository: %s", repo)
		return
	}
	isHub := !strings.Contains(comps[0], ".")
	var urlList string
	if isHub {
		urlList = "https://registry.hub.docker.com/v1/repositories/" + strings.Join(comps, "/") + "/tags"
	} else {
		urlList = "https://" + comps[0] + "/v2/" + strings.Join(comps[1:], "/") + "/tags/list"
	}
	var req *http.Request
	var res *http.Response
	if req, err = http.NewRequest(http.MethodGet, urlList, nil); err != nil {
		return
	}
	if res, err = http.DefaultClient.Do(req); err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad response: %s", res.Status)
		return
	}
	if isHub {
		var body []dockerHubTagItem
		if err = json.NewDecoder(res.Body).Decode(&body); err != nil {
			return
		}
		for _, item := range body {
			tags = append(tags, item.Name)
		}
	} else {
		var body registryListTagResponse
		if err = json.NewDecoder(res.Body).Decode(&body); err != nil {
			return
		}
		tags = body.Tags
	}
	return
}

func execute(dir, name string, args ...string) (err error) {
	log.Printf("Execute: %s %s", name, strings.Join(args, " "))
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if ee, ok := err.(*exec.ExitError); ok {
		log.Printf("completed, code(%d)", ee.ExitCode())
	}
	return
}
