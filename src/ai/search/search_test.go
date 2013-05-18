package search

import (
  "testing"
  "strconv"
  "math"
)

// PROBLEM DEFINITION

type SlidePuzzle struct{}

func (p SlidePuzzle) InitialState() interface{} {
  state := [][]int{
    []int{ 1, 4, 2 },
    []int{ 3, 0, 5 },
    []int{ 6, 7, 8 },
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

func (p SlidePuzzle) StepCost(state interface{}, action interface{}) float64 {
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

var expected = map[int][]int{
    0: []int{0, 0},
    1: []int{0, 1},
    2: []int{0, 2},
    3: []int{1, 0},
    4: []int{1, 1},
    5: []int{1, 2},
    6: []int{2, 0},
    7: []int{2, 1},
    8: []int{2, 2},
  }

func (p SlidePuzzle) H(x interface{}) float64 {
  // manhattan distance
  var total float64 = 0

  state := x.([][]int)

  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      if state[i][j] != 0 {
        exp := expected[state[i][j]]
        total += math.Abs(float64(i - exp[0]))
        total += math.Abs(float64(j - exp[1]))
      }
    }
  }

  return total
}


// END PROBLEM DEFINITION

func TestTreeSearch(t *testing.T) {
  problem := SlidePuzzle{}
  solution := TreeSearch(problem)

  testSolution(t, solution)
}


func TestGraphSearch(t *testing.T) {
  problem := SlidePuzzle{}
  solution := GraphSearch(problem)

  testSolution(t, solution)
}

func TestBreadthFirstSearch(t *testing.T) {
  problem := SlidePuzzle{}
  solution := BreadthFirstSearch(problem)

  testSolution(t, solution)
}

func TestUniformCostSearch(t *testing.T) {
  problem := SlidePuzzle{}
  solution := UniformCostSearch(problem)

  testSolution(t, solution)
}

func TestDepthLimitSearch(t *testing.T) {
  problem := SlidePuzzle{}

  _, err := DepthLimitedSearch(problem, 1)
  if err == nil {
    t.Fail()
  }

  solution, _ := DepthLimitedSearch(problem, 2)

  testSolution(t, solution)
}

func TestIterativeDeepeningSearch(t *testing.T) {
  problem := SlidePuzzle{}
  solution := IterativeDeepeningSearch(problem)

  testSolution(t, solution)
}

func TestRecursiveBestFirstSearch(t *testing.T) {
  problem := SlidePuzzle{}
  solution := RecursiveBestFirstSearch(problem)

  testSolution(t, solution)
}

func testSolution(t *testing.T, solution *Node) {
  if solution == nil {
    t.FailNow()
  }

  if solution.hash != "012345678" {
    t.Fail()
  }

  if solution.Cost != 2 {
    t.Errorf("expected %d, got %v", 2, solution.Cost)
  }
}
