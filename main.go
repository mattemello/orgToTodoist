package main

import (
	"fmt"
	"io"
	"os"

	"github.com/mattemello/orgToTodoist/cmd/git"
	"github.com/mattemello/orgToTodoist/cmd/parse"
)

func main() {
	fmt.Println(git.ControlRemote())
	args := os.Args[1:]
	message := "todo"

	if len(args) > 0 {
		isMessage := false
		for _, elem := range args {
			if elem[0] == '-' {
				switch elem[1] {

				case 'm':
					isMessage = true
					break

				}
			}

			fmt.Println(elem[0])
			if (elem[0] == 101 || elem[0] == '\'') && isMessage {
				message = elem
			}
		}
	}

	fmt.Println(message)

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

		git.CommitChange(message)
		if git.ControlRemote() {
			git.Push()
		}
		fmt.Println(file)
	}

}
