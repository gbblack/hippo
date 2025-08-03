package hangman

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Session struct {
	Input       io.Reader
	Output, Err io.Writer
}

type Game struct {
	Letters []string
	Tally   int
	Limit   int
	Current []string
}

func NewSession(in io.Reader, out, errs io.Writer) *Session {
	return &Session{
		Input:  in,
		Output: out,
		Err:    errs,
	}
}

func NewGame(word string) *Game {
	letters := strings.Split(word, "")
	word_len := len(letters)
	current := InitializeCurrent(word_len)
	return &Game{
		Letters: letters,
		Tally:   0,
		Limit:   6,
		Current: current,
	}
}

func InitializeCurrent(length int) []string {
	current := make([]string, length)
	for i := range current {
		current[i] = "_"
	}
	return current
}

func IncreaseTally(g Game) (Game, error) {
	g.Tally++
	return g, nil
}

func (g Game) GameOverCheck() bool {
	t := g.Tally
	return t > 5
}

func (g *Game) SetCurrent(letter string, index int) error {
	g.Current[index] = letter
	return nil
}

func (g Game) PlayerTurn(l string) error {
	for i, letter := range g.Letters {
		if l == letter {
			g.SetCurrent(l, i)
		}
	}
	return nil
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
