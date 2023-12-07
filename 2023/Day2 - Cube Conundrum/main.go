package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	inputFile = "./inputDay02.txt"
	Red       = "red"
	Green     = "green"
	Blue      = "blue"
	maxRed    = 12
	maxGreen  = 13
	maxBlue   = 14
)

type Game struct {
	ID     int
	Power  int
	Colors [][]byte
}

func main() {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal("failed to read input:", err)
	}

	games := processInput(input)

	totalID := 0
	totalPower := 0
	for _, game := range games {
		totalPower += game.Power
		if isPossible(game.Colors) {
			totalID += game.ID
		}
	}

	fmt.Println("Total gameIDs:", totalID, "and Total game Powers:", totalPower)
}

func processColors(games [][]byte, maxColors map[string]int) {
	for _, game := range games {
		for _, pick := range bytes.Split(game, []byte(",")) {
			pick = bytes.TrimSpace(pick)
			parts := bytes.Split(pick, []byte(" "))
			num, err := strconv.Atoi(string(parts[0]))
			if err != nil {
				log.Fatal("failed to get colors:", err)
			}
			color := string(parts[1])
			if num > maxColors[color] {
				maxColors[color] = num
			}
		}
	}
}

func processInput(input []byte) []Game {
	var games []Game
	headerLen := len("Game ")
	for _, line := range bytes.Split(input, []byte("\n")) {
		colon := bytes.IndexByte(line, ':')
		if colon == -1 {
			continue
		}

		gameID, err := strconv.Atoi(string(line[headerLen:colon]))
		if err != nil {
			log.Fatal("failed to get gameID:", err)
		}

		colors := bytes.Split(line[colon+1:], []byte(";"))
		maxColors := map[string]int{Red: 1, Green: 1, Blue: 1}
		processColors(colors, maxColors)

		power := maxColors[Red] * maxColors[Green] * maxColors[Blue]

		games = append(games, Game{
			ID:     gameID,
			Power:  power,
			Colors: colors,
		})
	}
	return games
}

func isPossible(games [][]byte) bool {
	maxColors := map[string]int{Red: maxRed, Green: maxGreen, Blue: maxBlue}
	processColors(games, maxColors)

	return maxColors[Red] <= maxRed && maxColors[Green] <= maxGreen && maxColors[Blue] <= maxBlue
}
