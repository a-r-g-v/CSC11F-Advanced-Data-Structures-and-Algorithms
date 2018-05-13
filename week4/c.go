package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func search(in [16]byte) int {
	for i := 0; i < 16; i++ {
		if in[i] == byte(0) {
			return i
		}
	}
	panic("not found 0")
}

func up(in [16]byte) (bool, [16]byte) {
	pos := search(in)
	if pos > 3 {
		in[pos], in[pos-4] = in[pos-4], in[pos]

	} else {
		return false, in
	}
	return true, in

}

func down(in [16]byte) (bool, [16]byte) {
	pos := search(in)
	if pos < 12 {
		in[pos], in[pos+4] = in[pos+4], in[pos]
	} else {
		return false, in

	}
	return true, in

}

func left(in [16]byte) (bool, [16]byte) {
	pos := search(in)
	// mod(pos, 4) == 0
	if pos != 0 && pos != 4 && pos != 8 && pos != 12 {
		in[pos], in[pos-1] = in[pos-1], in[pos]
	} else {
		return false, in
	}
	return true, in

}
func right(in [16]byte) (bool, [16]byte) {
	pos := search(in)
	// mod(pos, 4) == 3
	if pos != 3 && pos != 7 && pos != 11 && pos != 15 {
		in[pos], in[pos+1] = in[pos+1], in[pos]
	} else {
		return false, in

	}
	return true, in

}

func next() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i

}

var goal = [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0}

func pp(s [16]byte) {

	for i := 0; i < 4; i++ {
		fmt.Printf("%2d %2d %2d %2d\n", s[0+4*i], s[1+4*i], s[2+4*i], s[3+4*i])
	}
	fmt.Println("----------------------")

}

func dls(limit int, now int, state [16]byte, parents [][16]byte) bool {

	for _, s := range parents {
		if s == state {
			return false
		}

	}

	if now == limit+1 {
		return false

	}

	if state == goal {
		return true
	}

	parents = append(parents, state)

	ok, ur := up(state)
	if ok {
		var p [][16]byte
		copy(p, parents)
		if dls(limit, now+1, ur, p) {
			return true

		}
	}
	ok, dr := down(state)
	if ok {
		var p [][16]byte
		copy(p, parents)
		if dls(limit, now+1, dr, p) {
			return true

		}
	}
	ok, lr := left(state)
	if ok {
		var p [][16]byte
		copy(p, parents)
		if dls(limit, now+1, lr, p) {
			return true
		}
	}
	ok, rr := right(state)
	if ok {
		var p [][16]byte
		copy(p, parents)
		if dls(limit, now+1, rr, p) {
			return true
		}
	}

	return false

}

func main() {
	sc.Split(bufio.ScanWords)
	var state [16]byte
	for i := 0; i < 16; i++ {
		state[i] = byte(next())
	}
	for i := 0; i < 45; i++ {
		if dls(i, 0, state, [][16]byte{}) {
			fmt.Println(i)
			return
		}
	}
}
