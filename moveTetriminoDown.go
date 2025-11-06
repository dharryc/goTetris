package main

import "time"

func moveTetriminoDown(tetrimino *[]Coord) {
	for i := range *tetrimino {
		(*tetrimino)[i].Y -= 1
	}
}

func gameTick(gameSpeedChannel chan int, gameRunningChannel chan bool, currentTetrimino *[]Coord) {
	var gameRunning = true
	for gameRunning {
		time.Sleep(time.Duration(<-gameSpeedChannel) * time.Millisecond)
		select {
		case gameRunning = <-gameRunningChannel:
		default:
			moveTetriminoDown(currentTetrimino)
		}
	}
}