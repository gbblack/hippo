package hangman

import (
	"io"
	"strings"
)

type Game struct {
	Input       io.Reader
	Output, Err io.Writer
	Word        string
	Guesses     int
}

func NewGame(in io.Reader, out, errs io.Writer) *Game {
	return &Game{
		Input:   in,
		Output:  out,
		Err:     errs,
		Word:    "hello",
		Guesses: 0,
	}
}

func (g Game) Guess(guess string) bool {
	return strings.Contains(g.Word, guess)
}
