package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func reverseString(s string) string {
	var reversedString string = ""
	for _, c := range s {
		reversedString = string(c) + reversedString
	}
	return reversedString
}

func getFirstNumber(s string) int {
	for _, c := range s {
		if n, err := strconv.Atoi(string(c)); err == nil {
			return n
		}
	}
	return 0
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
		lineBackwards := reverseString(lineContent)
		firstNumber := getFirstNumber(lineContent)
		lastNumber := getFirstNumber(lineBackwards)

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
