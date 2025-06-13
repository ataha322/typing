package cmd

import (
	"github.com/ataha322/typing/game"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

const (
	defaultWordNum = 20
)

// wordsCmd represents the words command
var wordsCmd = &cobra.Command{
	Use:   "words [number of words]",
	Short: "Random words",
	Long: "Type N random words generated using " +
	"a list of commonly used english words. " +
	"N=20 given no argument.",

	Run: func(cmd *cobra.Command, args []string) {
		var numWords int
		switch len(args) {
		case 0:
			numWords = defaultWordNum
		case 1:
			num, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal(err)
			}
			if num <= 0 || num > 500 {
				log.Fatal("Error: need 0 < N <= 500")
			}
			numWords = num
		default:
			log.Fatal("Invalid number of arguments")
		}

		for game.StartWordsLoop(numWords) == 1 {}
	},
}

func init() {
	rootCmd.AddCommand(wordsCmd)
}
