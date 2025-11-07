package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"

	color "github.com/fatih/color"
)

func main() {
	// clear the terminal screen
	CallClear()

	// print welcome message
	fmt.Println("Welcome to Harry's Tetris in Go!")

	// make channel buffered so sender doesn't block
	var keyPressChannel = make(chan string, 8)

	// channel to keep track of game speed
	var gameSpeedChannel = make (chan int, 1)

	// channel to keep track of game running bool
	var gameRunningChannel = make (chan bool, 1)

	// set initial game speed
	gameSpeedChannel <- 1000 // milliseconds per tick

	// set the game to running
	var isGameRunning = true
	gameRunningChannel <- isGameRunning

	// board with mutex for safe concurrent access
	var (
		boardKey sync.RWMutex
		board    = make([][]int, 20)
	)
	boardKey.Lock()

	// start initial board (filled with 0s)
	for i := range board {
		board[i] = make([]int, 10)
	}

	boardKey.Unlock()

	go captureKeyPresses(keyPressChannel)
	go gameTick(gameSpeedChannel, gameRunningChanne)
}

func printTetris(board [][]int) {
	var text string
	for _, row := range board {
		for _, cell := range row {
			if cell == 0 {
				text += color.BgRGB(36, 36, 36).Sprintf(". ")
			} else {
				text += color.BgRGB(0, 0, 255).Sprintf("# ")
			}
		}
		text += "\r\n"
	}
	fmt.Print(text)
}

func get_tetrimino() []Coord {
	var random_tetrimino = rand.Intn(7)
	return tetriminos[random_tetrimino]
}
