package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

var validNumbers = map[string]int{
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getFirstNumber(inputString string) int {
	previousIndex := -1
	foundNumber := 0
	for key, value := range validNumbers {
		currentIndex := strings.Index(inputString, key)
		if currentIndex != -1 && (currentIndex <= previousIndex || previousIndex == -1) {
			previousIndex = currentIndex
			foundNumber = value
		}
	}
	return foundNumber
}

func getLastNumber(inputString string) int {
	previousIndex := -1
	foundNumber := 0
	for key, value := range validNumbers {
		currentIndex := strings.LastIndex(inputString, key)
		if currentIndex != -1 && (currentIndex >= previousIndex || previousIndex == -1) {
			previousIndex = currentIndex
			foundNumber = value
		}
	}
	return foundNumber
}

func main() {
	var finalNumber int = 0

	filePath, err := filepath.Abs("./advent/1/input.txt")
	checkErr(err)

	inputFile, err := os.Open(filePath)
	checkErr(err)
	defer inputFile.Close()

	inputScanner := bufio.NewScanner(inputFile)

	for inputScanner.Scan() {
		lineContent := inputScanner.Text()
		firstNumber := getFirstNumber(lineContent)
		lastNumber := getLastNumber(lineContent)

		concatNumbers := fmt.Sprintf("%v%v", firstNumber, lastNumber)
		calculatedNumber, err := strconv.ParseInt(concatNumbers, 10, 0)
		if err != nil {
			calculatedNumber = 0
		}

		finalNumber += int(calculatedNumber)
	}

	checkErr(inputScanner.Err())

	fmt.Println("The sum of the calculated numbers is:", finalNumber)
}
