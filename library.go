package library

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

const (
	dockerfileOut = "Dockerfile.out"
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
	Name string                            `yaml:"name"`
	Tags []manifestTag                     `yaml:"tags"`
	Vars map[string]map[string]interface{} `yaml:"vars"`
}

type manifestTag struct {
	Name       string   `yaml:"name"`
	Also       []string `yaml:"also"`
	Dockerfile string   `yaml:"dockerfile"`
	Vars       []string `yaml:"vars"`
}

type PullTask struct {
	Name string
}

func (p PullTask) Do() error {
	return Execute("", "docker", "pull", p.Name)
}

type BuildTask struct {
	Dir        string
	Repo       string
	Names      []string
	Doc        string
	Dockerfile string
	Vars       map[string]interface{}
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

func (b BuildTask) Do(push bool) (err error) {
	if err = ioutil.WriteFile(
		filepath.Join(b.Dir, "banner.minit.txt"),
		[]byte(fmt.Sprintf(
			"本镜像基于 %s 制作，详细信息参阅 %s",
			b.Names[0],
			strings.ReplaceAll(b.Doc, "{{.Repo}}", b.Repo),
		)),
		0644); err != nil {
		return
	}
	var buf []byte
	if buf, err = ioutil.ReadFile(filepath.Join(b.Dir, b.Dockerfile)); err != nil {
		return
	}
	var tmpl *template.Template
	if tmpl, err = template.New("__main__").Option("missingkey=zero").Parse(string(buf)); err != nil {
		return
	}
	out := &bytes.Buffer{}
	if err = tmpl.Execute(out, b.Vars); err != nil {
		return
	}
	out.WriteString("\nADD banner.minit.txt /etc/banner.minit.txt")
	if err = ioutil.WriteFile(filepath.Join(b.Dir, dockerfileOut), sanitize(out.Bytes()), 0640); err != nil {
		return
	}
	if err = Execute(b.Dir, "docker", "build", "-t", b.Names[0], "-f", dockerfileOut, "."); err != nil {
		return
	}
	for _, altName := range b.Names[1:] {
		if err = Execute(b.Dir, "docker", "tag", b.Names[0], altName); err != nil {
			return
		}
	}
	if push {
		if err = Execute(b.Dir, "docker", "push", b.Names[0]); err != nil {
			return
		}
		for _, altName := range b.Names[1:] {
			if err = Execute(b.Dir, "docker", "push", altName); err != nil {
				return
			}
		}
	}
	return
}

type MirrorTask struct {
	From string
	To   string
}

type MirrorSubTask MirrorTask

func (m MirrorTask) SubTasks(ctx context.Context) (tasks []MirrorSubTask, err error) {
	var tags []string
	if tags, err = RegistryListTags(ctx, m.From); err != nil {
		return
	}
	for _, tag := range tags {
		if strings.Contains(strings.ToLower(tag), "windows") {
			continue
		}
		tasks = append(tasks, MirrorSubTask{
			From: m.From + ":" + tag,
			To:   m.To + ":" + tag,
		})
	}
	return
}

func (s MirrorSubTask) Do() (err error) {
	if err = Execute("", "docker", "pull", s.From); err != nil {
		return
	}
	if err = Execute("", "docker", "tag", s.From, s.To); err != nil {
		return
	}
	if err = Execute("", "docker", "push", s.To); err != nil {
		return
	}
	return
}

var (
	Pulls   []PullTask
	Builds  []BuildTask
	Mirrors []MirrorTask
)

func init() {
	var err error
	defer func(err *error) {
		if *err != nil {
			panic(*err)
		}
	}(&err)
	var buf []byte
	if buf, err = ioutil.ReadFile("manifest.yml"); err != nil {
		return
	}
	var global manifestGlobal
	if err = yaml.Unmarshal(buf, &global); err != nil {
		return
	}
	// Pulls
	for _, upstream := range global.Upstreams {
		Pulls = append(Pulls, PullTask{
			Name: upstream,
		})
	}
	// Builds
	for _, dir := range global.Repos {
		var buf []byte
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
		for _, tag := range repo.Tags {
			if tag.Dockerfile == "" {
				tag.Dockerfile = "Dockerfile"
			}
			if tag.Name == "" {
				err = fmt.Errorf("missing name for tag in repo: %s", repo.Name)
				return
			}
			vars := map[string]interface{}{}
			for k, v := range global.Vars {
				vars[k] = v
			}
			for _, kg := range tag.Vars {
				if repo.Vars[kg] == nil {
					err = fmt.Errorf("missing vars group: %s in repo: %s ", kg, repo.Name)
					return
				}
				for k, v := range repo.Vars[kg] {
					vars[k] = v
				}
			}
			names := []string{path.Join(global.Base, repo.Name+":"+tag.Name)}
			for _, alt := range tag.Also {
				names = append(names, path.Join(global.Base, repo.Name+":"+alt))
			}
			Builds = append(Builds, BuildTask{
				Dir:        dir,
				Names:      names,
				Repo:       repo.Name,
				Doc:        global.Doc,
				Dockerfile: tag.Dockerfile,
				Vars:       vars,
			})
		}
	}
	// Mirrors
	for _, mirrorItem := range global.Mirrors {
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
		Mirrors = append(Mirrors, MirrorTask{
			From: src,
			To:   filepath.Join(global.Base, dst),
		})
	}
}
