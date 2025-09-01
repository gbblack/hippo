package shell

import (
	// "bufio"
	"errors"
	"flag"
	"fmt"
	"hippo/hangman"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

/*
you need to be able to:
Make a new session
? Start Menu for what game to play
TODO: Be able to select and then run a game
TODO: Gracefully exit at any point
TODO: Test all behaviors, happy and unhappy path
* Create a way to intereact with terminal using terminal inputs
? Use Doc Comments for documentation
? Save state for these games?
*/
type Session struct {
	Input       io.Reader
	Output, Err io.Writer
}

func NewSession(in io.Reader, out, errs io.Writer) *Session {
	return &Session{
		Input:  in,
		Output: out,
		Err:    errs,
	}
}

func HandleUserInput(i string) (rune, error) {
	i = strings.TrimSpace(i)
	letters := []rune(i)
	if len(letters) > 1 {
		return rune(0), errors.New("input is too long")
	}
	guess := letters[0]
	ok := unicode.IsLetter(guess)
	if !ok {
		return rune(0), errors.New("input is not a letter")
	}
	return guess, nil
}

func (s *Session) PlayHangman() {
	var wordfile string
	flag.StringVar(&wordfile, "words_file", "hangman/words.txt", "Words file")
	flag.Parse()
	if _, err := os.Stat(wordfile); err != nil {
		log.Fatalf("Could not open the words file: %s\n", err)
	}
	word, err := hangman.PickWord(wordfile)
	if err != nil {
		panic(err)
	}
	game := hangman.NewGame(word)
	fmt.Println(game)
}

func Main() {
	session := NewSession(os.Stdin, os.Stdout, os.Stderr)
	session.PlayHangman()
}
