package main

import (
	"os"
	//"os/exec"

	"github.com/mattemello/orgToTodoist/cmd/git"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		// todo: here control the tags
	}

	git.ControlChange()
	//exec.Command("git", "add", ".")

}
