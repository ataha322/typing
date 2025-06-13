package game

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/ataha322/typing/res"
	"golang.org/x/term"
)

func StartWordsLoop(num int) int {
	//raw-dog this terminal
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatal(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	ret := 1 // return 0 to stop, 1 to play again

	fmt.Printf("\x1b[3 q")       //change cursor to a blinking underline
	defer fmt.Printf("\x1b[0 q") //restore default cursor when finished

	var g Game
	chars := pickRandomWords(num)
	g.init(chars)

	b := make([]byte, 1)
	for g.curr_index < g.total_chars {
		os.Stdin.Read(b)
		char := b[0]

		g.updateDimensions() //in case terminal resizes

		if g.curr_index == 0 {
			g.start = time.Now()
			g.word_count = 0
		}

		if isPrintable(char) {
			g.printable(char)
		} else if isBackspace(char) {
			g.backspace()
		} else if isCtrlC(char) {
			ret = 0
			break
		} else if isCtrlR(char) {
			ret = 1
			break
		} else {
			//noop
		}
	}

	g.printResults()
	return ret
}

func pickRandomWords(num int) []rune {
	chosenWords := make([]string, num)
	for i := 0; i < num; i++ {
		randNum := rand.Intn(len(res.WordArr))
		chosenWords[i] = res.WordArr[randNum]
	}

	chars := make([]rune, 0)
	for _, word := range chosenWords {
		for _, char := range word {
			chars = append(chars, char)
		}
		chars = append(chars, ' ')
	}
	chars = chars[:len(chars)-1]
	return chars
}
