package search

type Problem interface {
  State() interface{}
  GoalState(interface{}) bool
  Expand(*Node) []*Node
  Hash(interface{}) string // this is only really needed for GraphSearch
}

type Node struct {
  State interface{}
  Parent *Node
  Action interface{}
  Cost int
}

func TreeSearch(problem Problem) *Node {
  node := &Node{ State: problem.State() }
  frontier := []*Node{ node }

  for {
    if len(frontier) == 0 {
      break // no solution
    }

    leaf := frontier[0]
    frontier = frontier[1:]

    if problem.GoalState(leaf.State) {
      return leaf
    }

    nodes := problem.Expand(leaf)
    frontier = append(frontier, nodes...)
  }

  return nil
}


// Add 8 puzzle test
func GraphSearch(problem Problem) *Node {
  state := problem.State()
  node := &Node{ State: state }

  frontier := []*Node{ node }
  frontierCache := map[interface{}]bool{} // for fast lookup
  explored := map[interface{}]bool{}

  h := problem.Hash(state)
  frontierCache[h] = true

  for {
    if len(frontier) == 0 {
      break // no solution
    }

    leaf := frontier[0]
    frontier = frontier[1:]

    if problem.GoalState(leaf.State) {
      return leaf
    }

    h := problem.Hash(leaf.State)
    explored[h] = true

    for _, node := range problem.Expand(leaf) {
      h := problem.Hash(node.State)

      _, exp := explored[h]
      _, front := frontierCache[h]

      if !exp && !front {
        frontier = append(frontier, node)
      }
    }
  }

  return nil
}
