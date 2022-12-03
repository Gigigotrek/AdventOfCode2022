package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

)

const abc = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getDuplicatedItems(firstItem string, secondItem string) (int) {
	var value int = 0
	for i := 0; i < len(firstItem); i++ {
		if strings.Contains(secondItem, string(firstItem[i])) {
			value = strings.Index(abc, string(firstItem[i])) + 1
		}
	}
	return value
}


func main() {
	var priorities int = 0
    tf, _ := os.ReadFile("./input")
    scanner := bufio.NewScanner(strings.NewReader(string(tf)))
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        firstItem := string(scanner.Text()[0:len(scanner.Text())/2])
		secondItem := string(scanner.Text()[(len(scanner.Text())/2):])
		
		priorities = priorities + getDuplicatedItems(firstItem, secondItem)
	}
	fmt.Println(priorities)
}