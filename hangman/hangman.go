package hangman

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"slices"
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
	Guessed []string
}

func NewSession(in io.Reader, out, errs io.Writer) *Session {
	return &Session{
		Input:  in,
		Output: out,
		Err:    errs,
	}
}

func WordFromSlice(s []string) string {
	i := rand.Intn(len(s))
	word := s[i]
	return word
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

func (g Game) AlreadyGuessed(l string) bool {
	return slices.Contains(g.Guessed, l)
}

func (g Game) PlayerTurn(l string) error {
	ok := g.AlreadyGuessed(l)
	if ok {
		return errors.New("you've already guessed this letter")
	}
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
