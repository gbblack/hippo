package hangman_test

import (
	"hangman"
	"os"
	"testing"
)

func TestNewSession_CreateExpectedNewSession(t *testing.T) {
	t.Parallel()
	want := hangman.Session{
		Input:  os.Stdin,
		Output: os.Stdout,
		Err:    os.Stderr,
	}
	got := *hangman.NewSession(os.Stdin, os.Stdout, os.Stderr)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestGuess_CorrectGuess(t *testing.T) {
	t.Parallel()
	game := hangman.Game{
		Letters: []string{"h", "e", "l", "l", "o"},
		Tally:   0,
	}
	want := true
	got := game.Guess("h")
	if want != got {
		t.Errorf("want %t, got %t", want, got)
	}
}

func TestGuess_IncorrectGuess(t *testing.T) {
	t.Parallel()
	game := hangman.Game{
		Letters: []string{"h", "e", "l", "l", "o"},
		Tally:   0,
	}
	want := false
	got := game.Guess("a")
	if want != got {
		t.Errorf("want %t, got %t", want, got)
	}
}

func TestIncreaseTally(t *testing.T) {
	t.Parallel()
	game := hangman.Game{
		Letters: []string{"h", "e", "l", "l", "o"},
		Tally:   0,
	}
	want := 1
	result, err := hangman.IncreaseTally(game)
	if err != nil {
		t.Fatal(err)
	}
	got := result.Tally
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestGameOverCheck(t *testing.T) {
	t.Parallel()
	game := hangman.Game{
		Letters: []string{"h", "e", "l", "l", "o"},
		Tally:   6,
	}
	want := true
	got := game.GameOverCheck()
	if want != got {
		t.Errorf("want %t, got %t", want, got)
	}
}
