package git

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
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
		log.Fatal("Error in the Run of the command 'git commit'", err)
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

func TakeChange(file string) []string {

	var changes []string

	commShow := exec.Command("git", "--no-pager", "show", file)
	var buff = bytes.Buffer{}

	commShow.Stdout = &buff
	err := commShow.Run()
	if err != nil {
		log.Fatal("Error in the Run of the command 'git show'", err)
	}

	var strCont string

	for err != io.EOF {
		var strTmp string
		strTmp, err = buff.ReadString(byte('.'))
		if err != nil {
			if err != io.EOF {
				//ERROR
			}
		}

		strCont += strTmp
	}

	strAll := strings.Split(strCont, "@@")

	for elem := range strings.SplitSeq(strAll[len(strAll)-1], "\n") {
		if elem != "" {
			if elem[0] == '+' || elem[0] == '-' {
				changes = append(changes, elem)
			}
		}
	}

	return changes

}
