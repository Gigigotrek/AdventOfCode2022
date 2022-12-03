package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

)

const abc = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getBadges(firstRucksack string, secondRucksack string, thirdRucksack string) (int) {
	var value int = 0
	for i := 0; i < len(firstRucksack); i++ {
		if strings.Contains(secondRucksack, string(firstRucksack[i])) {
			if strings.Contains(thirdRucksack, string(firstRucksack[i])) {
				value = strings.Index(abc, string(firstRucksack[i])) + 1
			}
		}
	}
	return value
}


func main() {
	var priorities int = 0
	var text []string
    tf, _ := os.ReadFile("./input")
    scanner := bufio.NewScanner(strings.NewReader(string(tf)))
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	for i := 0; i < len(text); i+=3 {
		var firstRucksack string = string(text[i])
		var secondRucksack string = string(text[i+1])
		var thirdRucksack string = string(text[i+2])

		getBadges(firstRucksack, secondRucksack, thirdRucksack)
		priorities = priorities + getBadges(firstRucksack, secondRucksack, thirdRucksack)
	}

	fmt.Println(priorities)
}