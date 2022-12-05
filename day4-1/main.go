package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getContainingSections(firstPair string, secondPair string) (bool) {
	var parsedFirstPair = strings.Split(firstPair, "-")
	var parsedSecondPair = strings.Split(secondPair, "-")
	var firstLocationFirstPair, _ = strconv.Atoi(parsedFirstPair[0])
	var secondLocationFirstPair, _ = strconv.Atoi(parsedFirstPair[1])
	var firstLocationSecondPair, _ = strconv.Atoi(parsedSecondPair[0])
	var secondLocationSecondPair, _ = strconv.Atoi(parsedSecondPair[1])
	if (secondLocationFirstPair - firstLocationFirstPair) >= (secondLocationSecondPair - firstLocationSecondPair) {
		if firstLocationFirstPair <= firstLocationSecondPair && secondLocationFirstPair >= secondLocationSecondPair {
			return true
		}
	} else {
		if firstLocationFirstPair >= firstLocationSecondPair && secondLocationFirstPair <= secondLocationSecondPair {
			return true
		}
	}
	return false
}

func main() {
    tf, _ := os.ReadFile("./input")
    scanner := bufio.NewScanner(strings.NewReader(string(tf)))
    scanner.Split(bufio.ScanLines)
	var count int = 0
    for scanner.Scan() {
		firstPair := strings.Split(string(scanner.Text()), ",")[0]
		secondPair := strings.Split(string(scanner.Text()), ",")[1]
		if getContainingSections(firstPair, secondPair) {
			count++
		}
	}
	fmt.Println(count)
}