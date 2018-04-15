package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var INF int = 1 << 31

type pair struct {
	total int
	next  int
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

type PriorityQueue []pair

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].total <= pq[j].total
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

type edge struct {
	to   int
	cost int
}

var sc = bufio.NewScanner(os.Stdin)

func next() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i

}

func main() {
	sc.Split(bufio.ScanWords)

	var V, E, r int
	fmt.Scanf("%d %d %d", &V, &E, &r)

	edges := make([][]*edge, V)
	dist := make([]int, V)
	done := make([]bool, V)
	queue := make(PriorityQueue, 0, V)

	for i := 0; i < E; i++ {
		var s, t, cost int
		s, t, cost = next(), next(), next()
		edges[s] = append(edges[s], &edge{to: t, cost: cost})

	}

	for i := 0; i < V; i++ {
		dist[i] = INF
	}

	heap.Push(&queue, pair{total: 0, next: r})

	for queue.Len() > 0 {
		item := heap.Pop(&queue).(pair)
		if done[item.next] {
			continue
		}
		done[item.next] = true
		dist[item.next] = item.total

		for _, edge := range edges[item.next] {
			heap.Push(&queue, pair{
				total: edge.cost + item.total,
				next:  edge.to,
			})

		}

	}

	for _, node := range dist {
		if node >= INF {
			fmt.Printf("INF\n")
		} else {
			fmt.Printf("%d\n", node)
		}
	}

}
