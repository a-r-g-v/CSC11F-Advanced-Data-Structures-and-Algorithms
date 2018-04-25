package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var MIN int = 0

var sc = bufio.NewScanner(os.Stdin)

func next() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i

}

type RMQ struct {
	data []int
	size int
}

func NewRMQ(size int) *RMQ {
	s := 1
	for s < size {
		s *= 2

	}
	r := make([]int, s+2)
	return &RMQ{data: r, size: s}
}

func (r *RMQ) update(k, a int) {
	k += r.size - 1
	r.data[k] += a
	for k > 0 {
		k = (k - 1) / 2
		r.data[k] = r.data[k*2+1] + r.data[k*2+2]
		//r.data[k] += a
	}

}

/// [a,b]
func (r *RMQ) findSum(a, b int) int {
	return r.query(a, b+1, 0, 0, r.size)
}

// [a,b)
func (rmq *RMQ) query(a, b, k, l, r int) int {
	// case1
	if r <= a || b <= l {
		return MIN
	}
	if a <= l && r <= b {
		return rmq.data[k]
	}

	vl := rmq.query(a, b, k*2+1, l, (l+r)/2)
	vr := rmq.query(a, b, k*2+2, (l+r)/2, r)
	return vl + vr
}

func main() {
	sc.Split(bufio.ScanWords)

	n, q := next(), next()

	rmq := NewRMQ(n)

	for i := 0; i < q; i++ {
		com, x, y := next(), next(), next()

		if com == 0 {
			x -= 1
			rmq.update(x, y)

		} else {
			x -= 1
			y -= 1
			fmt.Println(rmq.findSum(x, y))

		}
	}

}
