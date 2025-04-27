package git

import (
	"bytes"
	"log"
	"os/exec"
)

func ControlChange() bytes.Buffer {
	comm := exec.Command("git", "status")

	var buff = bytes.Buffer{}

	comm.Stdout = &buff

	err := comm.Run()
	if err != nil {
		log.Fatal("Error in the Run of the command 'git status'", err)
	}

	return buff
}
