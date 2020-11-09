package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"gopkg.in/yaml.v3"
)

const (
	defaultDockerfile    = "Dockerfile"
	defaultDockerfileOut = "Dockerfile.out"
)

type manifestGlobal struct {
	Base  string                 `yaml:"base"`
	Repos []string               `yaml:"repos"`
	Vars  map[string]interface{} `yaml:"vars"`
}

type manifestRepo struct {
	Name string                 `yaml:"string"`
	Tags []manifestTag          `yaml:"tags"`
	Vars map[string]interface{} `yaml:"vars"`
}

type manifestTag struct {
	Name       string                 `yaml:"name"`
	Dockerfile string                 `yaml:"dockerfile"`
	Vars       map[string]interface{} `yaml:"vars"`
}

func exit(err *error) {
	if *err != nil {
		log.Println("exited with error:", (*err).Error())
		os.Exit(1)
	}
}

func main() {
	var err error
	defer exit(&err)

	var buf []byte
	if buf, err = ioutil.ReadFile("manifest.yml"); err != nil {
		return
	}

	var global manifestGlobal
	if err = yaml.Unmarshal(buf, &global); err != nil {
		return
	}

	for _, dir := range global.Repos {
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
			for k, v := range repo.Vars {
				vars[k] = v
			}
			for k, v := range tag.Vars {
				vars[k] = v
			}
			log.Printf("Vars: %v", vars)
			if err = build(optsBuild{
				dir:        dir,
				base:       global.Base,
				repo:       repo.Name,
				tag:        tag.Name,
				dockerfile: tag.Dockerfile,
				vars:       vars,
			}); err != nil {
				return
			}
		}
	}
}

var tmplFuncs = template.FuncMap{}

type optsBuild struct {
	dir        string
	base       string
	repo       string
	tag        string
	dockerfile string
	vars       map[string]interface{}
}

func build(opts optsBuild) (err error) {
	var buf []byte
	if buf, err = ioutil.ReadFile(filepath.Join(opts.dir, opts.dockerfile)); err != nil {
		return
	}
	var tmpl *template.Template
	if tmpl, err = template.New("__main__").Option("missingkey=error").Funcs(tmplFuncs).Parse(string(buf)); err != nil {
		return
	}
	out := &bytes.Buffer{}
	if err = tmpl.Execute(out, opts.vars); err != nil {
		return
	}
	if err = ioutil.WriteFile(filepath.Join(opts.dir, defaultDockerfileOut), sanitize(out.Bytes()), 0640); err != nil {
		return
	}
	canonicalName := fmt.Sprintf("%s/%s:%s", opts.base, opts.repo, opts.tag)
	log.Println("Build:", canonicalName)
	cmd := exec.Command("docker", "build", "-t", canonicalName, "-f", defaultDockerfileOut, ".")
	cmd.Dir = opts.dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); err != nil {
		return
	}
	log.Println("Push:", canonicalName)
	cmd = exec.Command("docker", "push", canonicalName)
	cmd.Dir = opts.dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); err != nil {
		return
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
