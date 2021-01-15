package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/acicn/library"
)

const (
	all = "all"
)

var (
	optPull  bool
	optBuild string
	optPush  bool
)

func main() {
	var err error
	defer func(err *error) {
		if *err != nil {
			log.Println("exited with error:", (*err).Error())
			os.Exit(1)
		}
	}(&err)

	flag.BoolVar(&optPull, "pull", false, "pull upstreams")
	flag.StringVar(&optBuild, "build", all, "build repo")
	flag.BoolVar(&optPush, "push", false, "push after build")
	flag.Parse()

	// update image names
	var names []string
	for _, task := range library.Builds {
		for _, name := range task.Names {
			names = append(names, name)
		}
	}
	sort.Strings(names)
	if err = ioutil.WriteFile("IMAGES.txt", []byte(strings.Join(names, "\n")), 0644); err != nil {
		return
	}

	// pulls
	if optPull {
		for _, task := range library.Pulls {
			if err = task.Do(); err != nil {
				return
			}
		}
	}

	// builds
	for _, task := range library.Builds {
		if task.Repo != optBuild && optBuild != all {
			continue
		}
		if err = task.Do(optPush); err != nil {
			return
		}
	}
}
