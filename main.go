package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mattemello/orgToTodoist/cmd/git"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		// todo: here control the tags
	}

	control := git.ControlChange()

	var err error
	err = nil

	var strCont string

	for err != io.EOF {
		var strTmp string
		strTmp, err = control.ReadString(byte('.'))
		if err != nil {
			if err != io.EOF {
				//ERROR
			}
		}

		strCont += strTmp
	}

	indexMod := strings.Index(strCont, "modified: ")
	if indexMod == -1 {
		// todo: there are no modified, so need to do something?
	}
	indexAdd := strings.Index(strCont, "new file: ")
	if indexMod == -1 {
		// todo: there are no modified, so need to do something?
	}

	var whereToTake int
	whereToTake = min(indexMod, indexAdd)

	usefulCont := strCont[whereToTake:]

	fmt.Println(usefulCont)

	//exec.Command("git", "add", ".")

}
