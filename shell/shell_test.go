package shell_test

import (
	"hippo/shell"
	"os"
	"testing"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	testscript.Main(m, map[string]func(){
		"shell": shell.Main,
	})
}

func TestNewSession_CreateExpectedNewSession(t *testing.T) {
	t.Parallel()
	want := shell.Session{
		Input:  os.Stdin,
		Output: os.Stdout,
		Err:    os.Stderr,
	}
	got := *shell.NewSession(os.Stdin, os.Stdout, os.Stderr)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestHandleUserInput_SingleLetter(t *testing.T) {
	t.Parallel()
	in := "a"
	want := 'a'
	got, err := shell.HandleUserInput(in)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("want %c, got %c", want, got)
	}
}

func TestHandleUserInput_TooManyLetter(t *testing.T) {
	t.Parallel()
	in := "abc"
	_, err := shell.HandleUserInput(in)
	if err == nil {
		t.Error("want error if input is too long")
	}
}

func TestHandleUserInput_NotALetter(t *testing.T) {
	t.Parallel()
	in := "7"
	_, err := shell.HandleUserInput(in)
	if err == nil {
		t.Error("want error if input is not a letter")
	}
}