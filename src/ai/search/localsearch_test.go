package search

import (
  "testing"
  "strconv"
)


  // state[0][4] = 1
  // state[1][5] = 1
  // state[2][6] = 1
  // state[3][3] = 1
  // state[4][4] = 1
  // state[5][5] = 1
  // state[6][6] = 1
  // state[7][5] = 1


type Queen struct{}

func (problem Queen) InitialState() interface{} {
  state := [][]int{}

  for i := 0; i < 8; i++ {
    state = append(state, make([]int, 8))
  }

  state[0][7] = 1
  state[1][2] = 1
  state[2][6] = 1
  state[3][3] = 1
  state[4][0] = 1
  state[5][4] = 1
  state[6][0] = 1
  state[7][5] = 1

  return state
}

func (problem Queen) GoalState(x interface{}) bool {
  return problem.Hash(x) == "83742516"
}

func (problem Queen) Hash(x interface{}) string {
  state := x.([][]int)

  h := ""

  for i := 0; i < 8; i++ {
    for j := 0; j < 8; j++ {
      if state[i][j] == 1 {
        h += strconv.Itoa(j+1)
        break
      }
    }
  }

  return h
}

func (problem Queen) Result(x interface{}, action interface{}) interface{} {
  // action is passed as a hash representing the new formation. Ex: "81574274"

  state := [][]int{}

  for i := 0; i < 8; i++ {
    state = append(state, make([]int, 8))
  }

  for i, c := range action.(string) {
    n, _ := strconv.Atoi(string(c))
    state[i][n - 1] = 1
  }

  return state
}

func (problem Queen) StepCost(state interface{}, x interface{}) float64 {
  return 0 // we don't care
}

func (problem Queen) Actions(state interface{}) []interface{} {
  h := problem.Hash(state)

  permutations := make([]interface{}, 0)

  // generate permutations
  for i, c := range h {
    n, _ := strconv.Atoi(string(c))

    if n > 1 {
      s := []byte(h)
      s[i] = s[i] - 1
      permutations = append(permutations, string(s))
    }

    if n < 8 {
      s := []byte(h)
      s[i] = s[i] + 1
      permutations = append(permutations, string(s))
    }
  }

  return permutations
}

func (problem Queen) H(x interface{}) float64 {
  pairs := map[string]bool{}

  state := x.([][]int)

  // this algorithm only search down, right, up-right and down-right diagonal
  // returns the number of pairs of attacked queens
  for i := 0; i < 8; i++ {
    for j := 0; j < 8; j++ {
      if state[i][j] == 1 {
        // search down
        for cur := j + 1; cur < 8; cur++ {
          if state[i][cur] == 1 {
            pair_str := strconv.Itoa(i) + strconv.Itoa(j) + "-" + strconv.Itoa(i) + strconv.Itoa(cur)
            pairs[pair_str] = true
          }
        }

        // search right
        for cur := i + 1; cur < 8; cur++ {
          if state[cur][j] == 1 {
            pair_str := strconv.Itoa(i) + strconv.Itoa(j) + "-" + strconv.Itoa(cur) + strconv.Itoa(j)
            pairs[pair_str] = true
          }
        }

        // search up right
        for curI, curJ := i+1, j-1; curI < 8 && curJ >= 0; curI, curJ = curI+1, curJ-1 {
          if state[curI][curJ] == 1 {
            pair_str := strconv.Itoa(i) + strconv.Itoa(j) + "-" + strconv.Itoa(curI) + strconv.Itoa(curJ)
            pairs[pair_str] = true
          }
        }

        // search down right
        for curI, curJ := i+1, j+1; curI < 8 && curJ < 8; curI, curJ = curI+1, curJ+1 {
          if state[curI][curJ] == 1 {
            pair_str := strconv.Itoa(i) + strconv.Itoa(j) + "-" + strconv.Itoa(curI) + strconv.Itoa(curJ)
            pairs[pair_str] = true
          }
        }
      }
    }
  }

  n := 0

  for _ = range pairs {
    n += 1
  }

  return float64(n)
}

func TestHillClimbing(t *testing.T) {
  problem := Queen{}
  solution := HillClimbing(problem)

  if !problem.GoalState(solution) {
    t.Fail()
  }
}
