package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	part1()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Drawing struct {
	red   int
	blue  int
	green int
}

type Game struct {
	Number int
	Draws  []Drawing
}

func isGameViable(game Game, totalRed int, totalGreen int, totalBlue int) bool {
	for _, draw := range game.Draws {
		if draw.red > totalRed || draw.green > totalGreen || draw.blue > totalBlue {
			return false
		}
	}
	return true
}

func parseColorCount(s string, color string) int {
	result := 0
	if strings.Contains(s, color) {
		foundNumber, err := strconv.Atoi(strings.TrimSpace(strings.Replace(s, color, "", 0)))
		if err != nil {
			foundNumber = 0
		}
		result = foundNumber
	}
	return result
}

func part1() {
	totalRed := 12
	totalGreen := 13
	totalBlue := 14
	var games []Game

	filePath, err := filepath.Abs("./test_input.txt")
	checkErr(err)

	inputFile, err := os.Open(filePath)
	checkErr(err)
	defer inputFile.Close()
	inputScanner := bufio.NewScanner(inputFile)

	for inputScanner.Scan() {
		lineContent := inputScanner.Text()
		gameSplit := strings.Split(lineContent, ":")
		gameNumber, err := strconv.Atoi(strings.TrimSpace(strings.Replace(gameSplit[0], "Game", "", 0)))
		checkErr(err)

		game := Game{
			Number: gameNumber,
		}

		for _, rawDrawing := range strings.Split(gameSplit[1], ";") {
			var drawing Drawing
			for _, rawColor := range strings.Split(rawDrawing, ",") {
				drawing.red = parseColorCount(rawColor, "red")
				drawing.blue = parseColorCount(rawColor, "blue")
				drawing.green = parseColorCount(rawColor, "green")
			}
		}

		games = append(games, game)

	}

	sumViableGames := 0
	for _, game := range games {
		viableGame := isGameViable(game, totalRed, totalGreen, totalBlue)
		if viableGame {
			sumViableGames += game.Number
		}
	}

	fmt.Println("Viable game sum:", sumViableGames)
}
