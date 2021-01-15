package library

import (
	"context"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
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

type BuildTask struct {
	Dir        string
	Names      []string
	Doc        string
	Dockerfile string
	Vars       map[string]interface{}
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
		tasks = append(tasks, MirrorSubTask{
			From: m.From + ":" + tag,
			To:   m.To + ":" + tag,
		})
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
