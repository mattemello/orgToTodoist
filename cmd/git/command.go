package git

import (
	// "bufio"
	"bytes"
	"fmt"
	"os/exec"
)

func ControlChange() bytes.Buffer {
	comm := exec.Command("git", "status")
	fmt.Println(comm)

	var buff = bytes.Buffer{}

	comm.Stdout = &buff

	err := comm.Run()
	fmt.Println(err)

	return buff
}
