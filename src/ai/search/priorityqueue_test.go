package search

import (
  "testing"
  "container/heap"
)

func TestPriorityQueue(t *testing.T) {
  p := &PriorityQueue{}
  heap.Init(p)

  heap.Push(p, &Node{ h: 2 })
  heap.Push(p, &Node{ h: 5 })
  heap.Push(p, &Node{ h: 1 })
  heap.Push(p, &Node{ h: 3 })

  var n *Node

  n = heap.Pop(p).(*Node)
  if n.h != 1 { t.FailNow() }

  n = heap.Pop(p).(*Node)
  if n.h != 2 { t.FailNow() }

  n = heap.Pop(p).(*Node)
  if n.h != 3 { t.FailNow() }

  n = heap.Pop(p).(*Node)
  if n.h != 5 { t.FailNow() }

}


func TestPriorityQueueSwap(t *testing.T) {
  p := &PriorityQueue{}
  heap.Init(p)

  heap.Push(p, &Node{ h: 2, hash: "abc" })
  heap.Push(p, &Node{ h: 5, hash: "cba" })

  p.SwapIfLowerCost(&Node{ h: 4, hash: "cba" })

  var n *Node

  n = heap.Pop(p).(*Node)
  if n.h != 2 { t.FailNow() }

  n = heap.Pop(p).(*Node)
  if n.h != 4 { t.FailNow() }

  if p.Len() != 0 { t.FailNow() }

}
