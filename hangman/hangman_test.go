package hangman_test

import (
	"bytes"
	"hangman"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	testscript.Main(m, map[string]func(){
		"hangman": hangman.Main,
	})
}

func TestNewGame_CreateNewGame(t *testing.T) {
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
		Guessed: []string{},
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
	want := "> Please enter a word to guess \n> read line: >"
	got := out.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
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

func TestPlayerTurn_AlreadyGuessed(t *testing.T) {
	t.Parallel()
	game := hangman.Game{
		Letters: []string{"h", "e", "l", "l", "o"},
		Tally:   6,
		Current: []string{"_", "_", "_", "_", "_"},
		Guessed: []string{"l"},
	}
	err := game.PlayerTurn("l")
	if err == nil {
		t.Error("want error from guessing a same letter twice")
	}
}

// func TestPickWord_GetsWordFromFile(t *testing.T) {
// 	t.Parallel()
// 	slice := []string{"good", "great", "grievous"}
// 	got := hangman.WordFromSlice(slice)
// 	fmt.Printf("chosen word: %s", got)
// }

func TestSliceFromFile(t *testing.T) {
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
