package hangman

import (
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

func Guess(guess, letter string) bool {
	return guess == letter
}

func IncreaseTally(g Game) (Game, error) {
	g.Tally++
	return g, nil
}

func (g Game) GameOverCheck() bool {
	t := g.Tally
	return t > 5
}

func (g *Game) SetCurrent(letter string, index int) {
	g.Current[index] = letter
}
