package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
	"golang.org/x/term"
)

var clear map[string]func() //create a map for storing clear funcs
var gamespeed = 1000        //milliseconds

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

func main() {
	// make channel buffered so sender doesn't block
	var something = make(chan string, 8)
	go doSomething(something)
    
	for {
        board := make([][]int, 20)
        for i := range board {
            board[i] = make([]int, 10)
        }
		printTetris(put_tetrimino_on_board(board, get_tetrimino(), 3,0))
		select {
		case msg := <-something:
			if msg == "\x03" {
				return
			}
			fmt.Printf("key: %q\n", msg)
		default:
			time.Sleep(time.Duration(gamespeed) * time.Millisecond)
		}
		CallClear()
	}
}

func doSomething(something chan string) {
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
			something <- string(buf[:n])
		}
	}
}

func printTetris(board [][]int) {
	for _, row := range board {
		var line string
		for _, cell := range row {
			if cell == 0 {
				line += ". "
			} else {
				line += "# "
			}
		}
		fmt.Print(line + "\r\n")
	}
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
    for i := 0; i < len(tetrimino); i++ {
        for j := 0; j < len(tetrimino[i]); j++ {
            if tetrimino[i][j] == 1 {
                board[y+i][x+j] = 1
            }
        }
    }
    return board
}
