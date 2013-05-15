package search

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
}


func ChildNode(problem Problem, parent *Node, action interface{}) *Node {
  return &Node{
    State: problem.Result(parent.State, action),
    Parent: parent,
    Action: action,
    Cost: parent.Cost + problem.StepCost(parent.State, action),
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


func GraphSearch(problem Problem) *Node {
  state := problem.InitialState()
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

    for _, action := range problem.Actions(leaf.State) {
      node := ChildNode(problem, leaf, action)
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
