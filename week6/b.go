package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

var MIN int = 0

var sc = bufio.NewScanner(os.Stdin)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

type Node struct {
	key, pri    int
	left, rigth *Node
}

func next() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i

}

func rigthRotate(t *Node) *Node {
	s := t.left
	t.left = s.rigth
	s.rigth = t
	return s
}

func leftRotate(t *Node) *Node {
	s := t.rigth
	t.rigth = s.left
	s.left = t
	return s
}

func insert(t *Node, key int, pri int) *Node {
	if t == nil {
		return &Node{key: key, pri: pri}
	}
	if key == t.key {
		return t
	}
	if t.key > key {

		t.left = insert(t.left, key, pri)
		if t.pri < t.left.pri {
			t = rigthRotate(t)
		}
	} else {
		t.rigth = insert(t.rigth, key, pri)
		if t.pri < t.rigth.pri {
			t = leftRotate(t)
		}

	}
	return t
}

func find(t *Node, key int) bool {
	if t == nil {
		return false
	}

	if t.key == key {
		return true
	}

	if t.key > key {
		return find(t.left, key)

	} else if t.key < key {
		return find(t.rigth, key)

	}
	panic("error")
}

func erase(t *Node, key int) *Node {
	if t == nil {
		return nil
	}
	if key == t.key {
		if t.left == nil && t.rigth == nil {
			return nil

		} else if t.left == nil {
			t = leftRotate(t)
		} else if t.rigth == nil {
			t = rigthRotate(t)

		} else {
			if t.left.pri > t.rigth.pri {
				t = rigthRotate(t)

			} else {

				t = leftRotate(t)

			}

		}
		return erase(t, key)

	}
	if key < t.key {
		t.left = erase(t.left, key)

	} else {
		t.rigth = erase(t.rigth, key)

	}
	return t
}

func getMin(t *Node) int {
	if t.left != nil {
		return getMin(t.left)
	}
	return t.key

}

func main() {
	defer writer.Flush()

	sc.Split(bufio.ScanWords)

	var head *Node

	N, L := next(), next()

	items := make([]int, N)
	memo := make(map[int]int, N)

	for i := 0; i < N; i++ {
		items[i] = next()
	}

	for j := 0; j < L-1; j++ {
		head = insert(head, items[j], rand.Int())
		memo[items[j]] += 1
	}

	for i := L - 1; i < N; i++ {
		head = insert(head, items[i], rand.Int())
		memo[items[i]] += 1
		if i == N-1 {
			printf("%d\n", getMin(head))
		} else {
			printf("%d ", getMin(head))

		}

		r, ok := memo[items[i-L+1]]
		if ok {
			if r <= 1 {
				head = erase(head, items[i-L+1])
			}
			memo[items[i-L+1]] -= 1
		}
	}
}
