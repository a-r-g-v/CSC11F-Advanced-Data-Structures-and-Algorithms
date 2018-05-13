package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

var board [8][8]bool

func next() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i

}

func printBoard() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] {
				fmt.Print("Q")
			} else {
				fmt.Print(".")
			}

		}
		fmt.Print("\n")
	}

}

func check(x, y int) bool {
	for i := 0; i < 8; i++ {
	}
	for i := 0; i < 8; i++ {
	}
}

func dfs(N int) bool {
	if N == 8 {
		return true
	}

	for i := 0; i < 8; i++ {
		if board[N][i] {
			return dfs(N + 1)
		}
	}

	for i := 0; i < 8; i++ {
		if check(N, i) {
			fill(N, i)
			if !dfs(N + 1) {
				unfill(N, i)
			} else {
				break
			}
		}

	}

	return false

}

func unfill(x, y int) {
	board[x][y] = false
}

func fill(x, y int) {
	board[x][y] = true

}

func main() {
	sc.Split(bufio.ScanWords)

	queens := next()
	for n := 0; n < queens; n++ {
		x, y := next(), next()
		fill(x, y)
	}

	dfs(0)

	printBoard()
}
