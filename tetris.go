package main

import (
	"fmt"
	"golang.org/x/term"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
	color "github.com/fatih/color"
)

var clear map[string]func() //create a map for storing clear funcs
var gamespeed = 1000        //milliseconds

func main() {
	color.BgRGB(1, 179, 19)
	CallClear()
	fmt.Println("Welcome to Harry's Tetris in Go!")
	// make channel buffered so sender doesn't block
	var keyPressChannel = make(chan string, 8)
	go doSomething(keyPressChannel)
	board := make([][]int, 20)
	for i := range board {
		board[i] = make([]int, 10)
	}

	for {
		printTetris(put_tetrimino_on_board(board, get_tetrimino(), 3, 0))
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

func doSomething(keyPressChannel chan string) {
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		// fallback: just exit goroutine on error
		return
	}
	defer term.Restore(fd, oldState)

	buf := make([]byte, 3) // 1..3 bytes to capture common escape sequences
	for {
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
	var text string;
	for _, row := range board {
		for _, cell := range row {
			if cell == 0 {
				text += color.BgRGB(36, 36, 36).Sprintf("  ")
			} else {
				text += color.BgRGB(0, 0, 255).Sprintf("  ")
			}
		}
		text += fmt.Sprint("\r\n")
	}
	fmt.Print(text)
}

func get_tetrimino() [][]int {
	i_tetrimino := [][]int{
		{1, 1, 1, 1},
		{0, 0, 0, 0},
	}
	j_tetrimino := [][]int{
		{1, 0, 0},
		{1, 1, 1},
	}
	l_tetrimino := [][]int{
		{0, 0, 1},
		{1, 1, 1},
	}
	o_tetrimino := [][]int{
		{1, 1},
		{1, 1},
	}
	s_tetrimino := [][]int{
		{0, 1, 1},
		{1, 1, 0},
	}
	t_tetrimino := [][]int{
		{0, 1, 0},
		{1, 1, 1},
	}
	z_tetrimino := [][]int{
		{1, 1, 0},
		{0, 1, 1},
	}
	var random_tetrimino = rand.Intn(7)
	switch random_tetrimino {
	case 0:
		return i_tetrimino
	case 1:
		return j_tetrimino
	case 2:
		return l_tetrimino
	case 3:
		return o_tetrimino
	case 4:
		return s_tetrimino
	case 5:
		return t_tetrimino
	case 6:
		return z_tetrimino
	}
	return nil
}

func put_tetrimino_on_board(board [][]int, tetrimino [][]int, x int, y int) [][]int {
  o_tetrimino := [][]int{
		{1, 1},
		{1, 1},
	}
  var is_o_tetrimino bool = (len(tetrimino) == len(o_tetrimino) && len(tetrimino[0]) == len(o_tetrimino[0]))
	for i := 0; i < len(tetrimino); i++ {
		for j := 0; j < len(tetrimino[i]); j++ {
			if tetrimino[i][j] == 1 {
				if is_o_tetrimino {
					board[y+i][x+j+1] = 1
				} else {
					board[y+i][x+j] = 1
				}
			}
		}
	}
	return board
}

// From stackoverflow, neat way to clear terminal for clean print
func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}