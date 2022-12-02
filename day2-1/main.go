package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

)

func shapePoints(shape string) (int) {
	var points int = 0
	switch shape {
		case "X":
			points = 1
		case "Y":
			points = 2 
		case "Z":
			points = 3
	}
	return points
}

func handResult(firstShape string, secondShape string) (int) {
	var points int = 0
	switch {
		case firstShape == "A" && secondShape == "Y":
			points = 6
		case firstShape == "A" && secondShape == "Z":
			points = 0
		case firstShape == "B" && secondShape == "X":
			points = 0
		case firstShape == "B" && secondShape == "Z":
			points = 6
		case firstShape == "C" && secondShape == "X":
			points = 6
		case firstShape == "C" && secondShape == "Y":
			points = 0
		default:
			points = 3
	}

	points = points + shapePoints(secondShape)

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