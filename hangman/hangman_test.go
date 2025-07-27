package hangman_test

import (
	"hangman"
	"os"
	"testing"
)

func TestNewGame_CreateExpectedNewGame(t *testing.T) {
	t.Parallel()
	want := hangman.Game{
		Input:   os.Stdin,
		Output:  os.Stdout,
		Err:     os.Stderr,
		Word:    "hello",
		Guesses: 0,
	}
	got := *hangman.NewGame(os.Stdin, os.Stdout, os.Stderr)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestGuess_CorrectGuess(t *testing.T) {
	t.Parallel()
	game := hangman.NewGame(os.Stdin, os.Stdout, os.Stderr)
	want := true
	got := game.Guess("h")
	if want != got {
		t.Errorf("want %t, got %t", want, got)
	}
}

func TestGuess_IncorrectGuess(t *testing.T) {
	t.Parallel()
	game := hangman.NewGame(os.Stdin, os.Stdout, os.Stderr)
	want := false
	got := game.Guess("a")
	if want != got {
		t.Errorf("want %t, got %t", want, got)
	}
}
