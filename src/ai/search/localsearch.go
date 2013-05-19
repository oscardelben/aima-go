package search

import (
  "math"
  "math/rand"
  "time"
  "fmt"
)

// Returns a state that is a local maximum. It uses problem.H() to calculat the value of a node  (the lower the better)
func HillClimbing(problem Problem) interface{} {
  state := problem.InitialState()
	current := &Node{State: state, h: problem.H(state)}

	for {
		var neighbor *Node

		for _, action := range problem.Actions(current.State) {
      node := ChildNode(problem, current, action)

      if neighbor == nil { // first child
        neighbor = node
      } else if node.h < neighbor.h { // exchange if node is ranked better
        neighbor = node
      }
    }

		if neighbor.h >= current.h {
			return current.State
		}

    current = neighbor
	}

  return nil
}

// Untested
func SimulatedAnnealing(problem Problem, schedule func(int) float64) *Node {
  state := problem.InitialState()
  current := &Node{State: state, h: problem.H(state)}

  r := rand.New(rand.NewSource(time.Now().UnixNano()))

  for t := 1;; t++ {
    temperature := schedule(t)
    if temperature == 0 {
      return current
    }

    successors := []*Node{}
    for _, action := range problem.Actions(current.State) {
      node := ChildNode(problem, current, action)
      successors = append(successors, node)
    }

    successor := successors[r.Intn(len(successors))]

    delta := current.f - successor.f

    if delta > 0 {
      current = successor
    } else {
      r := r.Float32()
      probability := math.Exp(delta / temperature)

      if float64(r) <= probability {
        current = successor
      }
    }
  }

  return nil
}
