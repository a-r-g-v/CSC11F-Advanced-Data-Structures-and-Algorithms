package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var intmax = 2147483647

func correct(i int) (x, y int) {
	if i == 0 {
		return 3, 3
	}

	return (i - 1) / 4, int(math.Mod(float64(i-1), 4))

}

func manhattan(in [16]byte) (sum int) {

	for i := 0; i < 16; i++ {
		var nx, ny int
		if i == 15 {
			nx, ny = 3, 3
		} else {
			nx, ny = (i)/4, int(math.Mod(float64(i), 4))
		}

		cx, cy := correct(int(in[i]))
		r := int(math.Abs(float64(nx-cx))) + int(math.Abs(float64(ny-cy)))
		//fmt.Printf("%d: (%d, %d), (%d, %d) = %d \n", in[i], nx, ny, cx, cy, r)
		sum += r
	}

	return
}

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

func ulen(l [][16]byte) int {
	c := 0
	e := map[[16]byte]bool{}
	for i := 0; i < len(l); i++ {
		if !e[l[i]] {
			e[l[i]] = true
			c++
		}
	}
	return c
}

func dls(limit int, now int, state [16]byte, parents [][16]byte) (bool, int) {

	if state == goal {
		return true, ulen(parents)
	}

	if manhattan(state)+now > limit {
		return false, 0
	}

	parents = append(parents, state)

	ok, ur := up(state)
	if ok && !(len(parents) >= 1 && parents[len(parents)-1] == ur) {
		p := make([][16]byte, len(parents))
		copy(p, parents)
		ok, ans := dls(limit, now+1, ur, p)
		if ok {
			return true, ans
		}
	}
	ok, dr := down(state)
	if ok && !(len(parents) >= 1 && parents[len(parents)-1] == dr) {
		p := make([][16]byte, len(parents))
		copy(p, parents)
		ok, ans := dls(limit, now+1, dr, p)
		if ok {
			return true, ans
		}
	}
	ok, lr := left(state)
	if ok && !(len(parents) >= 1 && parents[len(parents)-1] == lr) {
		p := make([][16]byte, len(parents))
		copy(p, parents)
		ok, ans := dls(limit, now+1, lr, p)
		if ok {
			return true, ans
		}
	}
	ok, rr := right(state)
	if ok && !(len(parents) >= 1 && parents[len(parents)-1] == rr) {
		p := make([][16]byte, len(parents))
		copy(p, parents)
		ok, ans := dls(limit, now+1, rr, p)
		if ok {
			return true, ans
		}
	}

	return false, 0

}

func main() {
	sc.Split(bufio.ScanWords)
	var state [16]byte
	for i := 0; i < 16; i++ {
		state[i] = byte(next())
	}

	for bound := manhattan(state); bound < 45; bound++ {
		found, cost := dls(bound, 0, state, [][16]byte{})
		if found {
			fmt.Println(cost)
			return
		}
	}

	panic("error")
}
