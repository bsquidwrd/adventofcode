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
	var games []Game

	filePath, err := filepath.Abs("./input.txt")
	checkErr(err)

	inputFile, err := os.Open(filePath)
	checkErr(err)
	defer inputFile.Close()
	inputScanner := bufio.NewScanner(inputFile)

	for inputScanner.Scan() {
		lineContent := inputScanner.Text()
		game := assembleGame(lineContent)
		games = append(games, game)
	}

	part1(games)
	part2(games)
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

func (game Game) isGameViable(totalRed int, totalGreen int, totalBlue int) bool {
	for _, draw := range game.Draws {
		if draw.red == 0 && draw.blue == 0 && draw.green == 0 {
			return false
		}
		if draw.red > totalRed || draw.green > totalGreen || draw.blue > totalBlue {
			return false
		}
	}
	return true
}

func parseColorCount(s string, color string) int {
	splitString := strings.Fields(s)
	if strings.ToLower(splitString[1]) == strings.ToLower(color) {
		foundNumber, err := strconv.Atoi(splitString[0])
		if err != nil {
			foundNumber = 0
		}
		return foundNumber
	}
	return 0
}

func (game Game) getMinumumViableDrawing() Drawing {
	var result Drawing

	for _, drawing := range game.Draws {
		if drawing.red > result.red {
			result.red = drawing.red
		}
		if drawing.blue > result.blue {
			result.blue = drawing.blue
		}
		if drawing.green > result.green {
			result.green = drawing.green
		}
	}

	return result
}

func (drawing Drawing) getDrawingPower() int {
	return drawing.red * drawing.green * drawing.blue
}

func assembleGame(lineContent string) Game {
	gameSplit := strings.Split(lineContent, ":")
	trimmedGameNumber := strings.TrimSpace(strings.Split(gameSplit[0], " ")[1])
	gameNumber, err := strconv.Atoi(trimmedGameNumber)
	checkErr(err)

	game := Game{
		Number: gameNumber,
	}

	for _, rawDrawing := range strings.Split(gameSplit[1], ";") {
		var draw Drawing
		for _, rawColor := range strings.Split(rawDrawing, ",") {
			draw.red += parseColorCount(rawColor, "red")
			draw.blue += parseColorCount(rawColor, "blue")
			draw.green += parseColorCount(rawColor, "green")
		}
		game.Draws = append(game.Draws, draw)
	}
	return game
}

func part1(games []Game) {
	totalRed := 12
	totalGreen := 13
	totalBlue := 14

	sumViableGames := 0
	for _, game := range games {
		viableGame := game.isGameViable(totalRed, totalGreen, totalBlue)
		if viableGame {
			sumViableGames += game.Number
		}
	}

	fmt.Println("Part 1 answer:", sumViableGames)
}

func part2(games []Game) {
	totalDrawingPower := 0
	for _, game := range games {
		totalDrawingPower += game.getMinumumViableDrawing().getDrawingPower()
	}

	fmt.Println("Part 2 answer:", totalDrawingPower)
}
