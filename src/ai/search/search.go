package search

import (
  "container/heap"
)

// TODO:  USE POINTERS

type Problem interface {
  // The initial state of the problem
  InitialState() interface{}
  GoalState(interface{}) bool
  // A serialize version of state to use for caching
  Hash(interface{}) string
  // The state that results from calling problem(state, action)
  Result(interface{}, interface{}) interface{}
  // The cost of calling problem(state, action)
  StepCost(interface{}, interface{}) int
  // The actions that are possible from state
  Actions(interface{}) []interface{}
}

type Node struct {
  State interface{}
  Parent *Node
  Action interface{}
  Cost int
  index int // used in a priority queue
  hash string
}


func ChildNode(problem Problem, parent *Node, action interface{}) *Node {
  state := problem.Result(parent.State, action)
  return &Node{
    State: state,
    Parent: parent,
    Action: action,
    Cost: parent.Cost + problem.StepCost(parent.State, action),
    hash: problem.Hash(state),
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
