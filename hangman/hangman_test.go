package hangman_test

import (
	"bufio"
	"hippo/hangman"
	// "slices"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewGame(t *testing.T) {
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

func TestInitializeCurrent_Correct(t *testing.T) {
	t.Parallel()
	want := []string{"_", "_", "_"}
	got := hangman.InitializeCurrent(3)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
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

func TestPlayerGuess(t *testing.T) {
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

func TestPlayerGuess_Error(t *testing.T) {
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

func TestAlreadyGuessed(t *testing.T) {
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

func TestAddGuess(t *testing.T) {
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

func TestWordFromSlice_EmptySliceReturnsError(t *testing.T) {
	t.Parallel()
	_, err := hangman.WordFromSlice([]string{})
	if err == nil {
		t.Fatal("expected empty slice to error")
	}
}

func TestReadWords_Correct(t *testing.T) {
	t.Parallel()
	in := strings.NewReader("never gonna give you up")
	want := []string{"never", "gonna", "give", "you", "up"}
	got, err := hangman.ReadWords(in)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestReadWords_ScannerError(t *testing.T) {
	t.Parallel()
	in := strings.Repeat("x", bufio.MaxScanTokenSize+1)
	_, err := hangman.ReadWords(strings.NewReader(in))
	if err == nil {
		t.Fatal("expected scanner to fail if token too long")
	}
}

func TestReadWordFile_Correct(t *testing.T) {
	t.Parallel()
	path := "testdata/test_words.txt"
	want := []string{"never", "gonna", "give", "you", "up"}
	got, err := hangman.ReadWordFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestReadWordFile_FileNotFound(t *testing.T) {
	t.Parallel()
	path := ""
	_, err := hangman.ReadWordFile(path)
	if err == nil {
		t.Error("expected error when file is not found")
	}
}