package collector

import (
	"log"
	"os/exec"
	"strings"
)

func isUniq(newItem string, items []string) bool {
	for _, item := range items {
		if newItem == item {
			return false
		}
	}
	return true
}

type ShellCommand struct {
	command string
	arguments string
}

func RunShell(c ShellCommand) string {
	cmd := exec.Command(c.command, strings.Split(c.arguments, " ")...)
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatalln(err)
	}
	return strings.TrimSpace(string(stdout))
}
