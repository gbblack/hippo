package hangman

import (
	"fmt"
	"io"
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

func NewGame(letters []string) *Game {
	return &Game{
		Letters: letters,
		Tally:   0,
	}
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
	fmt.Fprintf(stdout, "Hello")
}
