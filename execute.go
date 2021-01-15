package library

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func execute(dir, name string, args ...string) (err error) {
	log.Printf("Execute: [%s] %s %s", dir, name, strings.Join(args, " "))
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
