# Harry's tetris in Go

start: go run . ->

print menu ->
- view leaderboard
- new game

view leaderboard ->
- read from file scores, names
- format and print score, name

new game ->
- initialize board, [][]int
- initialize moving piece map[int]int (coords)
- start game loop
- put new piece on board (check collision -> game over)
- listen for key presses for the duration of 1 game loop
- rotate / move left / move right / move down according to inputs
- advance board down 1
- check if moving piece would move into a set piece
- if moving piece has no room, {
- board = board + 1 @ moving piece coords
- check if any slice is full
- if slice(s) are full {
- board above moves down by one (animate?)
- incriment score according to # of rows gone }
- add moving piece = new random piece}


game over
- game of life on death screen


Insights about how to get this to actually work
- 3 threads
- 1 listens to keystrokes and updates board
- 1 moves living tetrimino down 1 every {gametick}
- 1 prints board anytime there's a change


## Credits
Colored board: https://github.com/fatih/color?tab=readme-ov-file
Clearing the console: https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go