package git

import (
	"bytes"
	"log"
	"os"
	"os/exec"
)

func Push() {
	commAdd := exec.Command("git", "push")
	err := commAdd.Run()
	if err != nil {
		log.Fatal("Error in the Run of the command 'git add'", err)
	}
}

func ControlChange() bytes.Buffer {
	commAdd := exec.Command("git", "add", ".")
	err := commAdd.Run()
	if err != nil {
		log.Fatal("Error in the Run of the command 'git add'", err)
	}
	comm := exec.Command("git", "status")
	var buff = bytes.Buffer{}

	comm.Stdout = &buff

	err = comm.Run()
	if err != nil {
		log.Fatal("Error in the Run of the command 'git status'", err)
	}

	return buff
}

func CommitChange(message string) {
	commComm := exec.Command("git", "commit", "-m", message)
	err := commComm.Run()
	if err != nil {
		log.Fatal("Error in the Run of the command 'git add'", err)
	}
}

func ControlRemote() bool {
	direc, err := os.Getwd()
	if err != nil {
		return false
	}
	_, err = os.Open(direc + "/.git/logs/refs/remotes")
	return err == nil
}
