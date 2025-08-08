package shell_test

import (
	"bytes"
	"io"
	"os"
	"shell"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestSessionRun(t *testing.T) {
	t.Parallel()
	in := strings.NewReader("")
	out := new(bytes.Buffer)
	session := shell.NewSession(in, out, io.Discard)
	session.Run()
	want := "> Please enter a word to guess \n> read line: >"
	got := out.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
