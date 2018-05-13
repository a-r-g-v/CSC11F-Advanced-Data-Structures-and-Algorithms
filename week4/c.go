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

type item struct {
	d *[16]byte
	c int
}

var queue = []item{}
var closed = map[[16]byte]bool{}

func bfs(cost int) int {
	for {
		top := queue[0]
		queue = queue[1:]

		if *top.d == goal {
			return top.c
		}
		closed[*top.d] = true

		ok, ur := up(*top.d)
		ok2, _ := closed[ur]
		if ok && !ok2 {
			queue = append(queue, item{&ur, top.c + 1})
		}
		ok, dr := down(*top.d)
		ok2, _ = closed[dr]
		if ok && !ok2 {
			queue = append(queue, item{&dr, top.c + 1})
		}
		ok, lr := left(*top.d)
		ok2, _ = closed[lr]
		if ok && !ok2 {
			queue = append(queue, item{&lr, top.c + 1})
		}
		ok, rr := right(*top.d)
		ok2, _ = closed[rr]
		if ok && !ok2 {
			queue = append(queue, item{&rr, top.c + 1})
		}
	}

}

func main() {
	sc.Split(bufio.ScanWords)
	var state [16]byte
	for i := 0; i < 16; i++ {
		state[i] = byte(next())
	}
	queue = append(queue, item{&state, 0})
	cost := bfs(0)
	fmt.Println(cost)
}
