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
	item   int
	prev   *node
	rank   int
	weight int
}

func (n *node) root() *node {
	if n.prev == nil {
		return n
	}
	r := n.prev.root()
	n.weight += n.prev.weight
	n.prev = r
	return r
}

func diff(nodeX, nodeY *node) int {
	return nodeY.Weight() - nodeX.Weight()
}

func issame(nodeX, nodeY *node) bool {
	return nodeX.root() == nodeY.root()

}

func (n *node) Weight() int {
	n.root()
	return n.weight
}

func merge(nodeX, nodeY *node, w int) {
	w += nodeX.Weight()
	w -= nodeY.Weight()

	nodeX = nodeX.root()
	nodeY = nodeY.root()
	if nodeX.item == nodeY.item {
		return
	}

	if nodeX.rank < nodeY.rank {
		nodeX, nodeY = nodeY, nodeX
		w *= -1
	}

	if nodeX.rank == nodeY.rank {
		nodeX.rank += 1
	}

	if nodeX.prev != nil || nodeY.prev != nil {
		panic(fmt.Sprintf("%v, %v", nodeX, nodeY))
	}

	nodeY.prev = nodeX
	nodeY.weight = w

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
			z := next()
			merge(groups[x], groups[y], z)

		} else {
			if issame(groups[x], groups[y]) {
				fmt.Println(diff(groups[x], groups[y]))
			} else {
				fmt.Println("?")
			}
		}
	}

}
