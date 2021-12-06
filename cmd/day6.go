package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(day6Cmd)

}

var day6Cmd = &cobra.Command{
	Use: "day6",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter input: ")
		Values, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		Vis := Input(Values)
		Mapepe := make(map[int]int)

		for _, visse := range Vis {
			Mapepe[visse]++
		}

		defer println("Een: ", Calc(Mapepe, 80))
		defer println("Twee: ", Calc(Mapepe, 256))
	},
}
var Mapepe = make(map[string]int)

func Calc(vis map[int]int, days int) int {
	var counter int

	for vis, values := range vis {
		counter += (growthCalc(vis, days, 1) * values)
	}

	return counter
}
func growthCalc(vis, dae, counter int) int {
	vis--
	dae--

	if vis < 0 && dae >= 0 {
		vis = 6

		key := fmt.Sprintf("%d:%d", dae, 8)

		if Mapepe[key] != 0 {
			counter += Mapepe[key]
		} else {
			growth := growthCalc(8, dae, 1)
			counter += growth
			Mapepe[key] = growth
		}

	}

	if dae < 0 {
		return counter
	}

	return growthCalc(vis, dae, counter)
}
