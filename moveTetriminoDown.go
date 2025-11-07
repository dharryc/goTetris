package main

import "time"

func moveTetriminoDown(tetrimino *[]Coord) {
	for i := range *tetrimino {
		(*tetrimino)[i].Y -= 1
	}
}

func gameTick(gameSpeedChannel chan int, gameRunningChannel chan bool, currentTetrimino *[]Coord) {
	var gameRunning = <-gameRunningChannel
	var tickSpeed = <-gameSpeedChannel
	for gameRunning {
		time.Sleep(time.Duration(tickSpeed) * time.Millisecond)
		select {
		case gameRunning = <-gameRunningChannel:
		case tickSpeed = <-gameSpeedChannel:
		default:
			moveTetriminoDown(currentTetrimino)
		}
	}
}