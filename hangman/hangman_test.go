package hangman_test

import (
	"hippo/hangman"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_NewGame(t *testing.T) {
	t.Parallel()
	want := hangman.Game{
		Letters: []string{"h", "e", "l", "l", "o"},
		Tally:   0,
		Limit:   6,
		Current: []string{"_", "_", "_", "_", "_"},
	}
	got := *hangman.NewGame("hello")
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func Test_InitializeCurrent_Correct(t *testing.T) {
	t.Parallel()
	want := []string{"_", "_", "_"}
	got := hangman.InitializeCurrent(3)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func Test_IncreaseTally(t *testing.T) {
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

func Test_GameOverCheck(t *testing.T) {
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

func Test_SetCurent(t *testing.T) {
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

func Test_PlayerGuess(t *testing.T) {
	t.Parallel()
	game := hangman.Game{
		Letters: []string{"h", "e", "l", "l", "o"},
		Tally:   6,
		Current: []string{"_", "_", "_", "_", "_"},
		Guessed: []string{},
	}
	want := []string{"_", "_", "l", "l", "_"}
	err := game.PlayerGuess("l")
	if err != nil {
		t.Fatal(err)
	}
	got := game.Current
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func Test_PlayerGuess_Error(t *testing.T) {
	t.Parallel()
	game := hangman.Game{
		Letters: []string{"h", "e", "l", "l", "o"},
		Tally:   6,
		Current: []string{"_", "_", "_", "_", "_"},
		Guessed: []string{"l"},
	}
	err := game.PlayerGuess("l")
	if err == nil {
		t.Error("want error from guessing a same letter twice")
	}
}

func Test_AlreadyGuessed(t *testing.T) {
	t.Parallel()
	game := hangman.Game{
		Guessed: []string{"a", "b"},
	}
	want := true
	got := game.AlreadyGuessed("a")
	if want != got {
		t.Errorf("want %t, got %t", want, got)
	}
}

func Test_SliceFromFile(t *testing.T) {
	t.Parallel()
	want := []string{"never", "gonna", "give", "you", "up"}
	got, err := hangman.SliceFromFile("testdata/test_words.txt")
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func Test_AddGuess(t *testing.T) {
	t.Parallel()
	game := hangman.Game{
		Guessed: []string{},
	}
	want := []string{"a"}
	err := game.AddGuess("a")
	if err != nil {
		t.Fatal(err)
	}
	got := game.Guessed
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
