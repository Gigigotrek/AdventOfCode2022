package main

import (
    "os"
    "bufio"
    "fmt"
    "io"
    "strconv"
    "strings"
)

func findMax(calsArrayByElf []int) (max int) {
	max = calsArrayByElf[0]
	for _, value := range calsArrayByElf {
		if value > max {
			max = value
		}
	}
	return max
}

func ReadCals(r io.Reader) ([]int, error) {
    scanner := bufio.NewScanner(r)
    scanner.Split(bufio.ScanLines)
	var singleElfCalsArray []int
	var elfsCalsArray []int
    for scanner.Scan() {
        cal, err := strconv.Atoi(scanner.Text())
        if err != nil {
			singleElfTotalCals := 0
			for _, v := range singleElfCalsArray {  
				singleElfTotalCals += v  
			}
			elfsCalsArray = append(elfsCalsArray, singleElfTotalCals)
			singleElfCalsArray = nil
        }
        singleElfCalsArray = append(singleElfCalsArray, cal)
    }
    return elfsCalsArray, scanner.Err()
}

func main() {
    tf, _ := os.ReadFile("./input")
    globalCalsArray, _ := ReadCals(strings.NewReader(string(tf)))
	var maxCals = findMax(globalCalsArray)
    fmt.Println(maxCals)
}