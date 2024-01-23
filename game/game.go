package game

import (
	"fmt"
	"log"
	"os"
	"time"
	"unicode"

	"golang.org/x/term"
)

const (
	WHITE = 37
	GREEN = 32
	RED   = 31

	CTRLC     = 3
	BACKSPACE = 127
)

type Game struct {
	text   []rune
	width  int
	height int

	total_chars int
	typed_chars int
	word_count  int
	curr_index  int
	num_lines   int
	mistyped    int

	start time.Time
}

func (g *Game) init(text []rune) {
	g.text = text
	var err error
	g.width, g.height, err = term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatal(err)
	}
	g.num_lines = (len(g.text) + g.width - 1) / g.width

	g.total_chars = len(text)
	g.word_count = 0
	g.curr_index = 0
	g.mistyped = 0
	g.typed_chars = 0

	g.start = time.Now()

	fmt.Printf("\x1b[%dm%s\x1b[0m", WHITE, string(g.text))
	fmt.Println()
	fmt.Printf("\x1b[%dF", g.num_lines)
}

func isPrintable(char byte) bool {
	r := rune(char)
	return unicode.IsLetter(r) || unicode.IsDigit(r) ||
		unicode.IsPunct(r) || unicode.IsSpace(r)
}

func isBackspace(char byte) bool {
	return char == BACKSPACE
}

func isCtrlC(char byte) bool {
	return char == CTRLC
}

func (g *Game) printable(char byte) {
	r := rune(char)
	correct := g.text[g.curr_index]
	var color int
	if r == correct {
		color = GREEN
	} else {
		g.mistyped++
		color = RED
	}
	fmt.Printf("\x1b[%dm%s\x1b[0m", color, string(correct))
	if g.curr_index+1 == g.total_chars || g.text[g.curr_index+1] == ' ' {
		g.word_count++
	}
	g.curr_index++
	g.typed_chars++
}

func (g *Game) backspace() {
	fmt.Printf("\x1b[%dm%s\x1b[0m", WHITE, string(g.text[g.curr_index]))
	fmt.Printf("\x1b[2D")
	if g.curr_index > 0 {
		g.curr_index--
	}
}

func (g *Game) printResults() {
	elapsed := time.Now().Sub(g.start)
	var wpm float64
	var accuracy float64 = 1.0

	wpm = float64(g.word_count) / elapsed.Minutes()
    if g.typed_chars > 0 {
        accuracy = 1.0 - float64(g.mistyped)/float64(g.typed_chars)
    }

	fmt.Printf("\x1b[%dE", g.num_lines)
	fmt.Printf("WPM: %d\n", int(wpm))
	fmt.Printf("\x1b[0G")
	fmt.Printf("Accuracy: %d%%\n", int(accuracy*100))
	fmt.Printf("\x1b[0G")
	fmt.Printf("Characters typed: %d\n", g.typed_chars)
	fmt.Printf("\x1b[0G")
}
