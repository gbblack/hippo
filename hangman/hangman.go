package hangman

import (
	"bufio"
	"errors"
	"math/rand"
	"os"
	"slices"
	"strings"
)

type Game struct {
	Letters []string
	Tally   int
	Limit   int
	Current []string
	Guessed []string
}

func SliceFromFile(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return []string{}, errors.New("unable to open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	words := []string{}
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	return words, nil
}

func WordFromSlice(s []string) string {
	i := rand.Intn(len(s))
	word := s[i]
	return word
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

func IncreaseTally(g Game) (Game, error) {
	g.Tally++
	return g, nil
}

func (g Game) GameOverCheck() bool {
	t := g.Tally
	return t > 5
}

func (g *Game) SetCurrent(letter string, index int) error {
	g.Current[index] = letter
	return nil
}

func (g Game) AlreadyGuessed(l string) bool {
	return slices.Contains(g.Guessed, l)
}

func (g Game) PlayerTurn(l string) error {
	ok := g.AlreadyGuessed(l)
	if ok {
		return errors.New("you've already guessed this letter")
	}
	for i, letter := range g.Letters {
		if l == letter {
			g.SetCurrent(l, i)
		}
	}
	return nil
}
