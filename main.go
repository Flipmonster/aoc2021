package main

import (
	"strconv"
	"strings"

	"github.com/shadowrizla/aoc2021/cmd"
)

func main() {
	cmd.Execute()
}

//input parse for use by some challenges
func Input(line string) (out []int) {
	for _, raw := range strings.Split(line, ",") {
		val, _ := strconv.Atoi(raw)
		out = append(out, val)
	}

	return out
}
