package hangman

import (
	"io"
	"strings"
)

type Game struct {
	Input       io.Reader
	Output, Err io.Writer
	Word        string
	Tally       int
}

func NewGame(in io.Reader, out, errs io.Writer) *Game {
	return &Game{
		Input:  in,
		Output: out,
		Err:    errs,
		Word:   "hello",
		Tally:  0,
	}
}

func (g Game) Guess(guess string) bool {
	return strings.Contains(g.Word, guess)
}

func Tally(g Game) (Game, error) {
	g.Tally++
	return g, nil
}
