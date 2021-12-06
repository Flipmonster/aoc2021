package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var day1Cmd = &cobra.Command{
	Use: "day1",

	Run: func(cmd *cobra.Command, args []string) {

		input := readFile("cmd/input/day1.txt")
		fmt.Println(len(input))
		defer println("Twee: ", twee(input))
		defer println("Een: ", een(input))

	},
}

func init() {
	rootCmd.AddCommand(day1Cmd)

}
func een(input []int) int {
	holdprev := 0
	counter := 0
	for _, val := range input {
		if holdprev != 0 && val > holdprev {
			counter++
		}
		holdprev = val
	}
	return counter
}

func twee(input []int) int {
	counter := 0
	for i := 1; i+2 < len(input); i = i + 1 {
		countprev := input[i-1] + input[i] + input[i+1]
		hold := input[i] + input[i+1] + input[i+2]
		if hold > countprev {
			counter++
		}
	}
	return counter
}
func readFile(filePath string) (numbers []int) {
	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", filePath, err))
	}
	defer fd.Close()
	var line int
	for {

		_, err := fmt.Fscanf(fd, "%d\n", &line)

		if err != nil {
			fmt.Println(err)
			if err == io.EOF {
				return
			}
			panic(fmt.Sprintf("Scan Failed %s: %v", filePath, err))

		}
		numbers = append(numbers, line)
	}

}
