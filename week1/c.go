package main

import (
	"container/heap"
	"fmt"
)

const INF = 9999999999999999

type node struct {
	me    int
	total int
	index int
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

type PriorityQueue []*node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].total < pq[j].total
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*node)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

type edge struct {
	to   int
	cost int
}

func main() {
	var V, E, r int
	fmt.Scanf("%d %d %d", &V, &E, &r)

	edges := make([][]edge, V)
	nodes := make([]*node, V)
	queue := make(PriorityQueue, V)

	for i := 0; i < E; i++ {
		var s, t, cost int
		n, _ := fmt.Scanf("%d %d %d", &s, &t, &cost)
		if n != 3 {
			panic(fmt.Sprintf("scaned string want 3 but %d", n))
		}
		edges[s] = append(edges[s], edge{to: t, cost: cost})

	}

	for i := 0; i < V; i++ {
		if r == i {
			nodes[i] = &node{me: i, total: 0}
			queue[i] = nodes[i]
		} else {
			nodes[i] = &node{me: i, total: INF}
			queue[i] = nodes[i]

		}
	}
	heap.Init(&queue)

	for queue.Len() > 0 {

		node := heap.Pop(&queue).(*node)

		for _, edge := range edges[node.me] {
			if nodes[edge.to].index == -1 {
				continue

			}
			alt := edge.cost + node.total

			if alt < nodes[edge.to].total {
				nodes[edge.to].total = alt
				heap.Fix(&queue, nodes[edge.to].index)

			}

		}

	}

	for _, node := range nodes {
		if node.total == INF {
			fmt.Printf("INF\n")
		} else {
			fmt.Printf("%d\n", node.total)
		}
	}

}
