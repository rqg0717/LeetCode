package main

import "fmt"

var dirX = []int{0, 1, 0, -1, 1, 1, -1, -1}
var dirY = []int{1, 0, -1, 0, 1, -1, 1, -1}

//directions = [(0, 1), (0, -1), (-1, -1), (-1, 0), (-1, 1), (1, 0), (1, -1), (1, 1)]

func updateBoard(board [][]byte, click []int) [][]byte {
	x, y := click[0], click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X'
	} else {
		game(board, x, y)
	}
	return board
}

func game(board [][]byte, x, y int) {
	count := 0
	for i := 0; i < 8; i++ {
		tx, ty := x+dirX[i], y+dirY[i]
		if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) {
			continue
		}
		if board[tx][ty] == 'M' {
			count++
		}
	}
	if count > 0 {
		board[x][y] = byte(count + '0')
	} else {
		board[x][y] = 'B'
		for i := 0; i < 8; i++ {
			tx, ty := x+dirX[i], y+dirY[i]
			if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) || board[tx][ty] != 'E' {
				continue
			}
			game(board, tx, ty)
		}
	}
}

func main() {
	board := [][]byte{{'E', 'M', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'M', 'M', 'E', 'M'},
		{'E', 'E', 'E', 'E', 'E'}}
	click := []int{1, 2}
	board = updateBoard(board, click)
	fmt.Println("Output: ")
	for _, array := range board {
		fmt.Printf("%s\n", string(array))
	}

}
