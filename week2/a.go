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
	rank int
}

func (n *node) root() *node {
	var parent *node
	for parent = n; parent.prev != nil; parent = parent.prev {
		continue
	}
	if parent != n {
		n.prev = parent
	}
	return parent
}

func issame(nodeX, nodeY *node) bool {
	return nodeX.root() == nodeY.root()

}

func merge(nodeX, nodeY *node) {
	parentX := nodeX.root()
	parentY := nodeY.root()
	if parentX.item == parentY.item {
		return
	}

	if nodeX.rank < nodeY.rank {
		nodeX, nodeY = nodeY, nodeX
	}

	if nodeX.rank == nodeY.rank {
		nodeX.rank += 1
	}

	if parentX.prev != nil || parentY.prev != nil {
		panic(fmt.Sprintf("%v, %v", parentX, parentY))
	}

	parentY.prev = parentX

}

func main() {
	sc.Split(bufio.ScanWords)
	var n, q int
	fmt.Scanf("%d %d", &n, &q)

	groups := make([]*node, n)
	for i := 0; i < n; i++ {
		groups[i] = &node{item: i, prev: nil, rank: 0}
	}

	for i := 0; i < q; i++ {
		com, x, y := next(), next(), next()
		if com == 0 {
			merge(groups[x], groups[y])

		} else {
			if issame(groups[x], groups[y]) {
				fmt.Println("1")
			} else {
				fmt.Println("0")
			}
		}
	}

}
