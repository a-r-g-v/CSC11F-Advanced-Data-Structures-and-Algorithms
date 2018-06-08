package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type pair struct {
	cost      int
	src, dest int
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

type PriorityQueue []pair

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost <= pq[j].cost
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(pair)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

var sc = bufio.NewScanner(os.Stdin)

func next() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i

}

type Tree struct {
	me   int
	prev *Tree
}

func getTop(t *Tree) *Tree {
	if t.prev != nil {
		r := getTop(t.prev)
		t.prev = r
		return r

	}
	return t

}

func same(a *Tree, b *Tree) bool {
	t1 := getTop(a)
	t2 := getTop(b)
	return t1.me == t2.me
}

func unite(a *Tree, b *Tree) {
	getTop(a).prev = b
}

func main() {
	sc.Split(bufio.ScanWords)

	vertexs, edges := next(), next()

	// heap queue
	queue := make(PriorityQueue, 0, edges)

	// union-find tree
	nodes := make([]*Tree, vertexs)

	// O(V)
	for i := 0; i < vertexs; i++ {
		nodes[i] = &Tree{me: i, prev: nil}

	}

	// O(E log E)
	for i := 0; i < edges; i++ {

		src, dest, cost := next(), next(), next()

		if src > vertexs || dest > vertexs {
			panic("over")
		}

		p := pair{src: src, dest: dest, cost: cost}
		// O(log E)
		heap.Push(&queue, p)
	}

	results := 0

	for queue.Len() > 0 {
		item := heap.Pop(&queue).(pair)

		if !same(nodes[item.src], nodes[item.dest]) {
			results += item.cost
			unite(nodes[item.src], nodes[item.dest])
		}
	}

	fmt.Println(results)

}
