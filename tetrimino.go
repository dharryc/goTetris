package main

type Coord struct {
	X int
	Y int
}

var i_tetrimino = []Coord{
	{3, 0},
	{4, 0},
	{5, 0},
	{6, 0},
}
var j_tetrimino = []Coord{
	{3, 0},
	{3, 1},
	{4, 1},
	{5, 1},
}
var l_tetrimino = []Coord{
	{5, 0},
	{3, 1},
	{4, 1},
	{5, 1},
}
var o_tetrimino = []Coord{
	{4, 0},
	{5, 0},
	{4, 1},
	{5, 1},
}
var s_tetrimino = []Coord{
	{3, 0},
	{4, 0},
	{2, 1},
	{3, 1},
}
var t_tetrimino = []Coord{
	{4, 0},
	{3, 1},
	{4, 1},
	{5, 1},
}
var z_tetrimino = []Coord{
	{3, 0},
	{4, 0},
	{4, 1},
	{5, 1},
}

var tetriminos = [][]Coord{
	i_tetrimino,
	j_tetrimino,
	l_tetrimino,
	o_tetrimino,
	s_tetrimino,
	t_tetrimino,
	z_tetrimino,
}