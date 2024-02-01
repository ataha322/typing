package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ataha322/typing/cmd"
	"golang.org/x/term"
)

func main() {
	//raw-dog this terminal
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatal(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

    //change cursor to a blinking underline
	fmt.Printf("\x1b[3 q")
	defer fmt.Printf("\x1b[0 q")
    
	cmd.Execute()
}
