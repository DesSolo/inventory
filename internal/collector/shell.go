package collector

import (
	"log"
	"os/exec"
	"strings"
)

type ShellCommand struct {
	Command   string
	Arguments string
}

func (s *ShellCommand) Execute() string {
	cmd := exec.Command(s.Command, strings.Split(s.Arguments, " ")...)
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatalln(err)
	}
	return strings.TrimSpace(string(stdout))
}
