package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

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
	sort.Ints(globalCalsArray)
    fmt.Println(globalCalsArray[len(globalCalsArray)-1] + globalCalsArray[len(globalCalsArray)-2] +globalCalsArray[len(globalCalsArray)-3] )
}