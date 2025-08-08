package shell

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Session struct {
	Input       io.Reader
	Output, Err io.Writer
}

func NewSession(in io.Reader, out, errs io.Writer) *Session {
	return &Session{
		Input:  in,
		Output: out,
		Err:    errs,
	}
}

func (s *Session) Run() {
	stdout := io.MultiWriter(s.Output)
	stderr := io.MultiWriter(s.Err)
	input := bufio.NewReader(s.Input)
	fmt.Fprintf(stdout, "> Please enter a word to guess \n> ")
	contents, err := input.ReadString('\n')
	if err != nil {
		fmt.Fprintln(stderr, "error: ", err)
	}
	fmt.Fprintf(stdout, "read line: %s>", contents)
}

func Main() {
	session := NewSession(os.Stdin, os.Stdout, os.Stderr)
	session.Run()
}
