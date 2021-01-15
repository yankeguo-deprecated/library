package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/acicn/library"
)

func main() {
	var err error
	defer func(err *error) {
		if *err != nil {
			log.Println("exited with error:", (*err).Error())
			os.Exit(1)
		}
	}(&err)

	var names []string

	for _, task := range library.Mirrors {
		var subTasks []library.MirrorSubTask
		if subTasks, err = task.SubTasks(context.Background()); err != nil {
			return
		}
		for _, subTask := range subTasks {
			if err = subTask.Do(); err != nil {
				return
			}
			names = append(names, subTask.To)
		}
	}

	sort.Strings(names)
	if err = ioutil.WriteFile("MIRRORS.txt", []byte(strings.Join(names, "\n")), 0644); err != nil {
		return
	}
}
