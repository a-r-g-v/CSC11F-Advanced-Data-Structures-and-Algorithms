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

var (
	edges   [][]int
	visited []int
)

func main() {
	sc.Split(bufio.ScanWords)

	var V, E int
	fmt.Scanf("%d %d %d", &V, &E)

	edges = make([][]int, V)
	visited = make([]int, V)

	for i := 0; i < E; i++ {
		s, e := next(), next()
		edges[s] = append(edges[s], e)
	}

	for i := 0; i < V; i++ {
		if DFS(i) {
			fmt.Println("1")
			return
		}
	}

	fmt.Println("0")
}

func DFS(v int) bool {
	var ret bool
	visited[v] = 1
	for _, e := range edges[v] {
		if visited[e] == 1 {
			ret = true
		}
		if visited[e] == 0 {
			ret = ret || DFS(e)
		}

	}
	visited[v] = 2
	return ret
}
