package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func next() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i

}

type node struct {
	item int
	prev *node
}

func main() {
	sc.Split(bufio.ScanWords)
	var n, q int
	fmt.Scanf("%d %d", &n, &q)

	groups := make([]*node, n)
	for i := 0; i < n; i++ {
		groups[i] = &node{item: i, prev: nil}
	}

	for i := 0; i < q; i++ {
		com, x, y := next(), next(), next()
		if com == 0 {
			var parentX, parentY *node
			for parentX = groups[x]; parentX.prev != nil; parentX = parentX.prev {
				continue
			}
			for parentY = groups[y]; parentY.prev != nil; parentY = parentY.prev {
				continue
			}
			if parentX.prev != nil || parentY.prev != nil {
				panic(fmt.Sprintf("%v, %v", parentX, parentY))
			}
			if parentX.item != parentY.item {
				parentX.prev = parentY

			}

		} else {
			var parentX, parentY *node
			for parentX = groups[x]; parentX.prev != nil; parentX = parentX.prev {
				continue
			}
			for parentY = groups[y]; parentY.prev != nil; parentY = parentY.prev {
				continue
			}
			if parentX.item == parentY.item {
				fmt.Println("1")
			} else {
				fmt.Println("0")
			}
		}
	}

}
