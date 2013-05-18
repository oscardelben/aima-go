package search

import (
  "container/heap"
  "errors"
  "math"
  "sort"
)

type Problem interface {
  // The initial state of the problem
  InitialState() interface{}
  GoalState(interface{}) bool
  // A serialize version of state to use for caching
  Hash(interface{}) string
  // The state that results from calling problem(state, action)
  Result(interface{}, interface{}) interface{}
  // The cost of calling problem(state, action)
  StepCost(interface{}, interface{}) float64
  // The actions that are possible from state
  Actions(interface{}) []interface{}
  // The estimated cost from the state to the solution
  H(interface{}) float64
}

type Node struct {
  State interface{}
  Parent *Node
  Action interface{}
  index int // used in the priority queue
  hash string // a string representation of state
  Cost float64 // also used as g()
  h float64 // estimated cost of the cheapest path from state to goal
  f float64 // Cost + h(), used for example in A* search
}

func ChildNode(problem Problem, parent *Node, action interface{}) *Node {
  state := problem.Result(parent.State, action)
  h := problem.H(state)
  cost := parent.Cost + problem.StepCost(parent.State, action)

  return &Node{
    State: state,
    Parent: parent,
    Action: action,
    Cost: cost,
    hash: problem.Hash(state),
    h: h,
    f: cost + h,
  }
}


func TreeSearch(problem Problem) *Node {
  node := &Node{ State: problem.InitialState() }
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

    for _, action := range problem.Actions(leaf.State) {
      frontier = append(frontier, ChildNode(problem, leaf, action))
    }
  }

  return nil
}


// General Graph Search Algorithm
func GraphSearch(problem Problem) *Node {
  state := problem.InitialState()
  h := problem.Hash(state)
  node := &Node{ State: state, hash: h }

  frontier := []*Node{ node } // used as a FIFO queue
  frontierCache := map[interface{}]bool{} // for fast lookup
  explored := map[interface{}]bool{}

  frontierCache[node.hash] = true

  for {
    if len(frontier) == 0 {
      break // no solution
    }

    leaf := frontier[0]
    frontier = frontier[1:]

    if problem.GoalState(leaf.State) {
      return leaf
    }

    explored[leaf.hash] = true

    for _, action := range problem.Actions(leaf.State) {
      node := ChildNode(problem, leaf, action)

      _, exp := explored[node.hash]
      _, front := frontierCache[node.hash]

      if !exp && !front {
        frontier = append(frontier, node)
        frontierCache[node.hash] = true
      }
    }
  }

  return nil
}

// Breadth First Search - very similar to a standard Graph Search Algorithm


func BreadthFirstSearch(problem Problem) *Node {
  state := problem.InitialState()
  h := problem.Hash(state)
  node := &Node{ State: state, hash: h }

  if problem.GoalState(state) {
    return node
  }

  frontier := []*Node{ node } // used as a FIFO queue
  frontierCache := map[interface{}]bool{} // for fast lookup
  explored := map[interface{}]bool{}

  frontierCache[node.hash] = true

  for {
    if len(frontier) == 0 {
      break // no solution
    }

    leaf := frontier[0]
    frontier = frontier[1:]

    explored[leaf.hash] = true

    for _, action := range problem.Actions(leaf.State) {
      node := ChildNode(problem, leaf, action)

      _, exp := explored[node.hash]
      _, front := frontierCache[node.hash]

      if !exp && !front {
        if problem.GoalState(node.State) {
          return node
        }

        frontier = append(frontier, node)
        frontierCache[node.hash] = true
      }
    }
  }

  return nil
}

// Expands nodes with lower cost first
// If your problem implements H() this is the same as A* Search
func UniformCostSearch(problem Problem) *Node {
  state := problem.InitialState()
  h := problem.Hash(state)
  node := &Node{ State: state, hash: h }

  frontier := &PriorityQueue{}
  heap.Init(frontier)
  heap.Push(frontier, node)

  frontierCache := map[interface{}]bool{} // for fast lookup
  explored := map[interface{}]bool{}

  frontierCache[node.hash] = true

  for {
    if frontier.Len() == 0 {
      break // no solution
    }

    leaf := heap.Pop(frontier).(*Node)

    if problem.GoalState(leaf.State) {
      return leaf
    }

    explored[leaf.hash] = true

    for _, action := range problem.Actions(leaf.State) {
      node := ChildNode(problem, leaf, action)

      _, exp := explored[node.hash]
      _, front := frontierCache[node.hash]

      if !exp && !front {
        heap.Push(frontier, node)
        frontierCache[node.hash] = true
      } else {
        frontier.SwapIfLowerCost(node)
      }
    }

  }

  return nil
}

// returns error if cut off
func DepthLimitedSearch(problem Problem, limit int) (*Node, error) {
  state := problem.InitialState()
  h := problem.Hash(state)
  node := &Node{ State: state, hash: h }

  return recursiveDLS(node, problem, limit)
}

func IterativeDeepeningSearch(problem Problem) *Node {
  depth := 0
  for {
    result, _ := DepthLimitedSearch(problem, depth)

    if result != nil {
      return result
    }

    depth += 1
  }

  return nil
}

func recursiveDLS(node *Node, problem Problem, limit int) (*Node, error) {
  if problem.GoalState(node.State) {
    return node, nil
  }

  if limit == 0 {
    return nil, errors.New("cutoff")
  }

  cutoff := false

  for _, action := range problem.Actions(node.State) {
    child := ChildNode(problem, node, action)
    result, err := recursiveDLS(child, problem, limit-1)

    if err != nil {
      cutoff = true
    } else if result != nil {
      return result, nil
    }
  }

  if cutoff {
    return nil, errors.New("cutoff")
  }

  return nil, nil // no result
}


func RecursiveBestFirstSearch(problem Problem) *Node {
  state := problem.InitialState()
  h := problem.H(state)
  node := &Node{
    State: state,
    hash: problem.Hash(state),
    Cost: 0,
    h: h,
    f: h, // Cost is 0
  }

  solution, _ := rbfs(problem, node, 9999) // todo use max value of float64
  return solution
}

type Nodes []*Node

// for sort

func (nodes Nodes) Len() int {
  return len(nodes)
}

func (nodes Nodes) Less(i, j int) bool {
  return nodes[i].f < nodes[j].f
}

func (nodes Nodes) Swap(i, j int) {
  nodes[i], nodes[j] = nodes[j], nodes[i]
}

func rbfs(problem Problem, node *Node, fLimit float64) (result *Node, newLimit float64) {
  if problem.GoalState(node.State) {
    return node, 0
  }

  successors := Nodes{}

  for _, action := range problem.Actions(node.State) {
    child := ChildNode(problem, node, action)
    successors = append(successors, child)
  }

  if len(successors) == 0 {
    return nil, 9999 // TODO: use max float
  }

  for i := 0; i < len(successors); i++ {
    s := successors[i]
    successors[i].f = math.Max(s.f, node.f)
  }

  for {

    sort.Sort(successors) // sort by f()

    best :=  successors[0]


    if best.f > fLimit {
      return nil, best.f
    }

    alternative := successors[1].f

    result, bestf := rbfs(problem, best, math.Min(fLimit, alternative))

    best.f = bestf

    if result != nil {
      return result, 0
    }

  }

  return nil, 0
}
