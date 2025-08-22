package shell

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	// "flag"
	"hippo/hangman"
	"strings"
	"unicode"
)

/*
you need to be able to:
Make a new session
TODO: Start Menu for what game to play
TODO: Be able to select and then run a game
TODO: Gracefully exit at any point
TODO: Test all behaviors, happy and unhappy path
* Create a way to intereact with terminal using terminal inputs
? Use Doc Comments for documentation
? Save state for these games?
*/
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
	fmt.Fprintf(stdout, "> Pick a game \n> ")
	contents, err := input.ReadString('\n')
	if err != nil {
		fmt.Fprintln(stderr, "error: ", err)
	}
	contents = strings.TrimSpace(contents)
	if contents == "hangman" {
		s.PlayHangman()
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

func (s *Session) PlayHangman() {
	game := hangman.NewGame("hello")
	_, err := game.ReadWordFile("shell/words.txt")
	if err != nil {
		panic(err)
	}
}

func Main() {
	session := NewSession(os.Stdin, os.Stdout, os.Stderr)
	session.Run()
}
