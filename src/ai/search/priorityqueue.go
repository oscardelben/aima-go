package search

import (
  "container/heap"
)

// Nodes with lower cost have higher priority

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
  return pq[i].h < pq[j].h
}

func (pq PriorityQueue) Swap(i, j int) {
  pq[i], pq[j] = pq[j], pq[i]
  pq[i].index = i
  pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
  n := len(*pq)
  item := x.(*Node)
  item.index = n
  *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
  old := *pq
  n := len(old)
  item := old[n-1]
  item.index = -1 // for safety
  *pq = old[0 : n-1]
  return item
}

func (pq *PriorityQueue) SwapIfLowerCost(x interface{}) {
  n := len(*pq)
  item := x.(*Node)

  for i := 0; i < n; i++ {
    cur := (*pq)[i]
    if cur.hash == item.hash && cur.h > item.h {
      heap.Remove(pq, cur.index)
      heap.Push(pq, item)
      return
    }
  }
}
