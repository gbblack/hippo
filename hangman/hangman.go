package hangman

import (
	"bufio"
	"errors"
	"io"
	"math/rand"
	"os"
	"slices"
	"strings"
)

var wordfile = "words.txt"

type Game struct {
	Letters []string
	Tally   int
	Limit   int
	Current []string
	Guessed []string
}

func NewGame(word string) *Game {
	letters := strings.Split(word, "")
	word_len := len(letters)
	current := InitializeCurrent(word_len)
	return &Game{
		Letters: letters,
		Tally:   0,
		Limit:   6,
		Current: current,
	}
}

func InitializeCurrent(length int) []string {
	current := make([]string, length)
	for i := range current {
		current[i] = "_"
	}
	return current
}

func (g Game) PlayerGuess(l string) error {
	ok := g.AlreadyGuessed(l)
	if ok {
		return errors.New("you've already guessed this letter")
	}
	g.AddGuess(l)
	for i, letter := range g.Letters {
		if l == letter {
			g.SetCurrent(l, i)
		}
	}
	return nil
}

func (g Game) AlreadyGuessed(l string) bool {
	return slices.Contains(g.Guessed, l)
}

func IncreaseTally(g Game) (Game, error) {
	g.Tally++
	return g, nil
}

// WordsFromSlice return a single string element from a string slice, chose at random.
func WordFromSlice(s []string) (string, error) {
	if len(s) == 0 {
		return "", errors.New("slice is empty")
	}
	i := rand.Intn(len(s))
	word := s[i]
	return word, nil
}

func (g *Game) SetCurrent(letter string, index int) error {
	g.Current[index] = letter
	return nil
}

func (g *Game) AddGuess(l string) error {
	g.Guessed = append(g.Guessed, l)
	return nil
}

func (g Game) GameOverCheck() bool {
	t := g.Tally
	return t > 5
}

func ReadWords(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var words = []string{}
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		words = append(words, word)
	}
	if err := scanner.Err(); err != nil {
		return []string{}, err
	}
	return words, nil
}

func ReadWordFile(pathname string) ([]string, error) {
	f, err := os.Open(pathname)
	if err != nil {
		return []string{}, err
	}
	defer f.Close()
	words, err := ReadWords(f)
	if err != nil {
		return []string{}, err
	}
	return words, nil
}

func PickWord() (string, error) {
	words, err := ReadWordFile(wordfile)
	if err != nil {
		return "", err
	}
	word, err := WordFromSlice(words)
	if err != nil {
		return "", err
	}
	return word, nil
}