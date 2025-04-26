package git

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
)

func ControlChange() bool {
	comm := exec.Command("git", "status")
	fmt.Println(comm)

	var a bytes.Buffer
	var c bytes.Buffer

	var r = bufio.NewReader(&c)
	var b = bufio.NewWriter(&a)

	var rb = bufio.NewReadWriter(r, b)

	comm.Stdout = rb

	err := comm.Run()
	fmt.Println(err)

	var prova = make([]byte, 5)
	rb.Read(prova)
	fmt.Println(prova)

	return false
}
