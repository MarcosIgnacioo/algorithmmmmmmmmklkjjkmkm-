package graphs

import (
	"marcos.com/algorithms/src/12-trees/arraylist"
	"marcos.com/algorithms/src/12-trees/queues"
)

type Matrix [][]int

func BFSMatrix(graph Matrix, source int, needle int) arraylist.ArrayList {
	queue := queues.NewQueue()
	seen := arraylist.NewArrayList(uint(len(graph[0])))
	path := arraylist.NewArrayList(uint(len(graph[0])))
	seen.Fill(false)
	path.Fill(-1)

	seen.ArrayList[source] = true
	queue.Enqueue(source)

	curr := queue.Dequeue().(int)
	for queue.Length > 0 {
		if curr == needle {
			break
		}
		for i := 0; i < len(graph[curr]); i++ {
			if graph[curr][i] == 0 {
				continue
			}
			if seen.ArrayList[i] == true {
				continue
			}
			seen.ArrayList[i] = true
			path.Push(curr)
			queue.Enqueue(graph[curr][i])
		}
		seen.ArrayList[curr] = true
	}

	curr = needle
	out := arraylist.NewArrayList(10)

	for path.ArrayList[curr] != -1 {
		out.Push(path.ArrayList[curr])
		curr = path.ArrayList[curr].(int)
	}

	return out
}
