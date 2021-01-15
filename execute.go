package library

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func Execute(dir, name string, args ...string) (err error) {
	if dir == "" {
		log.Printf("Execute: %s %s", name, strings.Join(args, " "))
	} else {
		log.Printf("Execute: [%s] %s %s", dir, name, strings.Join(args, " "))
	}
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if ee, ok := err.(*exec.ExitError); ok {
		log.Printf("Completed, CODE = %d", ee.ExitCode())
	}
	return
}
