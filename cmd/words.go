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
		"a list of popular english words. " +
		"N=20 given no argument.",

	Run: func(cmd *cobra.Command, args []string) {
        switch len(args) {
        case 0:
			game.StartWordsLoop(defaultWordNum)
        case 1:
			num, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal(err)
			}
            if num <= 0 || num > 500 {
                log.Fatal("Error: need 0 < N <= 500")
            }
			game.StartWordsLoop(num)
        default:
			log.Fatal("Invalid number of arguments")
		}
	},
}

func init() {
	rootCmd.AddCommand(wordsCmd)
}
