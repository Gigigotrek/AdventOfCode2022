package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func handResult(firstShape string, secondShape string) (int) {
	var points int = 0
	switch {
		case firstShape == "A" && secondShape == "X":
			points = 3
		case firstShape == "A" && secondShape == "Y":
			points = 4
		case firstShape == "A" && secondShape == "Z":
			points = 8
		case firstShape == "B" && secondShape == "X":
			points = 1
		case firstShape == "B" && secondShape == "Y":
			points = 5
		case firstShape == "B" && secondShape == "Z":
			points = 9
		case firstShape == "C" && secondShape == "X":
			points = 2
		case firstShape == "C" && secondShape == "Y":
			points = 6
		case firstShape == "C" && secondShape == "Z":
			points = 7
	}

	return points
}

func main() {

	var totalPoints int = 0
    tf, _ := os.ReadFile("./input")
    scanner := bufio.NewScanner(strings.NewReader(string(tf)))
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        play := scanner.Text()
		totalPoints = totalPoints + handResult(string(play[0]), string(play[2]))
	}
	fmt.Println(totalPoints)
}