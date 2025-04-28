package main

import (
	"fmt"
	"io"
	"os"

	"github.com/mattemello/orgToTodoist/cmd/git"
	"github.com/mattemello/orgToTodoist/cmd/parse"
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

	file := parse.ParseAdd(strCont)

	if file != nil {
		fmt.Println(file)
	}

	//exec.Command("git", "add", ".")

}
