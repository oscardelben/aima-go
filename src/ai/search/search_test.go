package search

import (
  "testing"
  "strconv"
)

// TREE SEARCH

type Sum struct{}

func (p Sum) State() interface{} {
  return 1
}

func (p Sum) GoalState(x interface{}) bool {
  return x.(int) == 5
}

func (p Sum) Expand(node *Node) []*Node {
  newNode := &Node{
    Parent: node,
    Cost: node.Cost + 1,
    State: node.State.(int) + 2,
    Action: "sum 2",
  }

  return []*Node{ newNode }
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

func (p SlidePuzzle) State() interface{} {
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

func (p SlidePuzzle) Expand(node *Node) []*Node {

  state := node.State.([][]int)

  var x, y int
  // get zero position
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      if state[i][j] == 0 {
        x = i
        y = j
      }
    }
  }

  newState := func() [][]int {
    s := [][]int{
      []int{0, 0, 0},
      []int{0, 0, 0},
      []int{0, 0, 0},
    }

    for i := 0; i < 3; i++ {
      for j := 0; j < 3; j++ {
        s[i][j] = state[i][j]
      }
    }

    return s
  }

  nodes := []*Node{}

  // up
  if x > 0 {
    s := newState()
    s[x-1][y], s[x][y] = s[x][y], s[x-1][y]

    n := &Node{
      Parent: node,
      Cost: node.Cost + 1,
      State: s,
      Action: "up",
    }

    nodes = append(nodes, n)
  }

  // down
  if x < 2 {
    s := newState()
    s[x+1][y], s[x][y] = s[x][y], s[x+1][y]

    n := &Node{
      Parent: node,
      Cost: node.Cost + 1,
      State: s,
      Action: "down",
    }

    nodes = append(nodes, n)
  }

  // left
  if y > 0 {
    s := newState()
    s[x][y], s[x][y-1] = s[x][y-1], s[x][y]

    n := &Node{
      Parent: node,
      Cost: node.Cost + 1,
      State: s,
      Action: "left",
    }

    nodes = append(nodes, n)
  }

  // right
  if y < 2 {
    s := newState()
    s[x][y], s[x][y+1] = s[x][y+1], s[x][y]

    n := &Node{
      Parent: node,
      Cost: node.Cost + 1,
      State: s,
      Action: "right",
    }

    nodes = append(nodes, n)
  }


  return nodes
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

  if problem.Hash(solution.State) != "012345678" {
    t.Fail()
  }

  if solution.Cost != 26 {
    t.Errorf("expected %d, got %d", 26, solution.Cost)
  }
}
