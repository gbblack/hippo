package hangman_test

import (
	"bytes"
	"hangman"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestSetCurent(t *testing.T) {
	t.Parallel()
	game := hangman.Game{
		Letters: []string{"h", "e", "l", "l", "o"},
		Tally:   6,
		Current: []string{"_", "_", "_", "l", "_"},
	}
	want := []string{"_", "_", "l", "l", "_"}
	err := game.SetCurrent("l", 2)
	if err != nil {
		t.Fatal(err)
	}
	got := game.Current
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestPlayerTurn(t *testing.T) {
	t.Parallel()
	game := hangman.Game{
		Letters: []string{"h", "e", "l", "l", "o"},
		Tally:   6,
		Current: []string{"_", "_", "_", "_", "_"},
	}
	want := []string{"_", "_", "l", "l", "_"}
	err := game.PlayerTurn("l")
	if err != nil {
		t.Fatal(err)
	}
	got := game.Current
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestSessionRun(t *testing.T) {
	t.Parallel()
	in := strings.NewReader("")
	out := new(bytes.Buffer)
	session := hangman.NewSession(in, out, io.Discard)
	session.Run()
	want := "Hello"
	got := out.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
