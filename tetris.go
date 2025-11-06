package main

import (
	"fmt"
	color "github.com/fatih/color"
	"golang.org/x/term"
	"math/rand"
	"os"
	"time"
)

func main() {
	color.BgRGB(1, 179, 19)
	CallClear()
	fmt.Println("Welcome to Harry's Tetris in Go!")
	// make channel buffered so sender doesn't block
	var keyPressChannel = make(chan string, 8)
	var gameSpeedChannel = make (chan int, 1000)
	var gameRunningChannel = make (chan bool, 1)
	go captureKeyPresses(keyPressChannel)
	board := make([][]int, 20)
	for i := range board {
		board[i] = make([]int, 10)
	}

	for {
		printTetris(board)
		select {
		case msg := <-keyPressChannel:
			if msg == "\x03" {
				return
			}
		default:
			time.Sleep(time.Duration(gamespeed) * time.Millisecond)
		}
		CallClear()
	}
}

func captureKeyPresses(keyPressChannel chan string) {
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		// fallback: just exit goroutine on error
		return
	}
	defer term.Restore(fd, oldState)
	var quit = false

	buf := make([]byte, 3) // 1..3 bytes to capture common escape sequences
	for !quit {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			return
		}
		if n > 0 {
			// send the raw bytes as a string; arrow keys / function keys are escape sequences
			keyPressChannel <- string(buf[:n])
		}
	}
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
