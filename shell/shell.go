package shell

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
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
	fmt.Fprintf(stdout, "> Make a guess \n> ")
	contents, err := input.ReadString('\n')
	if err != nil {
		fmt.Fprintln(stderr, "error: ", err)
	}
	guess, err := HandleUserInput(contents)
	if err != nil {
		fmt.Fprintln(stderr, "error: ", err)
	}
	fmt.Fprintf(stdout, "%c", guess)
}

func HandleUserInput(i string) (rune, error) {
	i = strings.TrimSpace(i)
	letters := []rune(i)
	if len(letters) > 1 {
		return rune(0), errors.New("input is too long")
	}
	guess := letters[0]
	ok := unicode.IsLetter(guess)
	if !ok {
		return rune(0), errors.New("input is not a letter")
	}
	return guess, nil
}

func Main() {
	session := NewSession(os.Stdin, os.Stdout, os.Stderr)
	session.Run()
}
