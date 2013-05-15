package search

import (
  "testing"
  "strconv"
  "container/heap"
)

// TREE SEARCH

type Sum struct{}

func (p Sum) InitialState() interface{} {
  return 1
}

func (p Sum) GoalState(x interface{}) bool {
  return x.(int) == 5
}

func (p Sum) Result(state interface{}, action interface{}) interface{} {
  return state.(int) + 2
}

func (p Sum) StepCost(state interface{}, action interface{}) int {
  return 1
}

func (p Sum) Actions(state interface{}) []interface{} {
  a := make([]interface{}, 0)
  a = append(a, "sum 2")
  return a
}

func (p Sum) Hash(x interface{}) string {
  return ""
}


func TestTreeSearch(t *testing.T) {
  problem := Sum{}
  solution := TreeSearch(problem)

  if solution == nil {
    t.FailNow()
  }

  if solution.State.(int) != 5 {
    t.Errorf("expected %d, got %d", 5, solution.State.(int))
  }

  if solution.Cost != 2 {
    t.Errorf("expected %d, got %d", 2, solution.Cost)
  }
}

// GRAPH SEARCH


type SlidePuzzle struct{}

func (p SlidePuzzle) InitialState() interface{} {
  state := [][]int{
    []int{ 7, 2, 4 },
    []int{ 5, 0, 6 },
    []int{ 8, 3, 1 },
  }

  return state
}

func (p SlidePuzzle) GoalState(x interface{}) bool {
  return p.Hash(x) == "012345678"
}

func (p SlidePuzzle) Result(state interface{}, action interface{}) interface{} {
  st := state.([][]int)

  s := [][]int{
    []int{0, 0, 0},
    []int{0, 0, 0},
    []int{0, 0, 0},
  }

  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      s[i][j] = st[i][j]
    }
  }

   var x, y int
  // get zero position
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      if st[i][j] == 0 {
        x = i
        y = j
      }
    }
  }

  switch action.(string) {
  case "up":
    s[x-1][y], s[x][y] = s[x][y], s[x-1][y]
  case "down":
    s[x+1][y], s[x][y] = s[x][y], s[x+1][y]
  case "left":
    s[x][y], s[x][y-1] = s[x][y-1], s[x][y]
  case "right":
    s[x][y], s[x][y+1] = s[x][y+1], s[x][y]
  }
  return s
}

func (p SlidePuzzle) StepCost(state interface{}, action interface{}) int {
  return 1
}

func (p SlidePuzzle) Actions(state interface{}) []interface{} {
  actions := make([]interface{}, 0)

  st := state.([][]int)

  var x, y int
  // get zero position
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      if st[i][j] == 0 {
        x = i
        y = j
      }
    }
  }

  if x > 0 { actions = append(actions, "up") }
  if x < 2 { actions = append(actions, "down") }
  if y > 0 { actions = append(actions, "left") }
  if y < 2 { actions = append(actions, "right") }

  return actions
}


func (p SlidePuzzle) Hash(x interface{}) string {
  str := ""
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      val := x.([][]int)[i][j]
      str += strconv.Itoa(val)
    }
  }

  return str
}


func TestGraphSearch(t *testing.T) {
  problem := SlidePuzzle{}
  solution := GraphSearch(problem)

  if solution == nil {
    t.FailNow()
  }

  if solution.hash != "012345678" {
    t.Fail()
  }

  if solution.Cost != 26 {
    t.Errorf("expected %d, got %d", 26, solution.Cost)
  }
}

func TestBreadthFirstSearch(t *testing.T) {
  problem := SlidePuzzle{}
  solution := BreadthFirstSearch(problem)

  if solution == nil {
    t.FailNow()
  }

  if solution.hash != "012345678" {
    t.Fail()
  }

  if solution.Cost != 26 {
    t.Errorf("expected %d, got %d", 26, solution.Cost)
  }
}



// Priority Queue

func TestPriorityQueue(t *testing.T) {
  p := &PriorityQueue{}
  heap.Init(p)

  heap.Push(p, &Node{ Cost: 2 })
  heap.Push(p, &Node{ Cost: 5 })
  heap.Push(p, &Node{ Cost: 1 })
  heap.Push(p, &Node{ Cost: 3 })

  var n *Node

  n = heap.Pop(p).(*Node)
  if n.Cost != 1 { t.FailNow() }

  n = heap.Pop(p).(*Node)
  if n.Cost != 2 { t.FailNow() }

  n = heap.Pop(p).(*Node)
  if n.Cost != 3 { t.FailNow() }

  n = heap.Pop(p).(*Node)
  if n.Cost != 5 { t.FailNow() }

}


func TestPriorityQueueSwap(t *testing.T) {
  p := &PriorityQueue{}
  heap.Init(p)

  heap.Push(p, &Node{ Cost: 2, hash: "abc" })
  heap.Push(p, &Node{ Cost: 5, hash: "cba" })

  p.SwapIfLowerCost(&Node{ Cost: 4, hash: "cba" })

  var n *Node

  n = heap.Pop(p).(*Node)
  if n.Cost != 2 { t.FailNow() }

  n = heap.Pop(p).(*Node)
  if n.Cost != 4 { t.FailNow() }

  if p.Len() != 0 { t.FailNow() }

}
