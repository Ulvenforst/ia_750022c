package main

import (
  agents "github.com/Ulvenforst/ia_750022c/agentes"
)

func main() {
  testSimpleAgent()
}

// En un agente simple a cada percepción se le asocia una acción. 
func testSimpleAgent() {
  // 1 es un obstáculo y 2 es el objetivo.
  envMatrix := [][]int{
    {2,2,0,0},
    {1,0,0,1},
    {0,0,0,0},
    {0,0,1,0},
  }
  enviroment := agents.NewEnviroment(envMatrix, -1)
  // El conocimiento se representa mediante una tabla. 
  // {Up, Right, Down, Left, Goal}
  knowledge := map[[5]int]func(a *agents.SimpleAgent){
    {1, 1, 1, 1}: func(a *agents.SimpleAgent) { a.MoveUp() },
    {1, 1, 0, 1}: func(a *agents.SimpleAgent) { a.MoveUp() },
    {1, 0, 1, 1}: func(a *agents.SimpleAgent) { a.MoveUp() },
    {1, 0, 0, 1}: func(a *agents.SimpleAgent) { a.MoveLeft() },
    {0, 1, 1, 1}: func(a *agents.SimpleAgent) { a.MoveLeft() },
    {0, 1, 0, 1}: func(a *agents.SimpleAgent) { a.MoveRight() },
    {0, 0, 1, 1}: func(a *agents.SimpleAgent) { a.MoveLeft() },
    {0, 0, 0, 1}: func(a *agents.SimpleAgent) { a.MoveLeft() },
    {1, 1, 1, 0}: func(a *agents.SimpleAgent) { a.MoveUp() },
    {1, 1, 0, 0}: func(a *agents.SimpleAgent) { a.MoveRight() },
    {1, 0, 1, 0}: func(a *agents.SimpleAgent) { a.MoveDown() },
    {1, 0, 0, 0}: func(a *agents.SimpleAgent) { a.MoveUp() },
    {0, 1, 1, 0}: func(a *agents.SimpleAgent) { a.MoveRight() },
    {0, 1, 0, 0}: func(a *agents.SimpleAgent) { a.MoveRight() },
    {0, 0, 1, 0}: func(a *agents.SimpleAgent) { a.MoveDown() },
  }
  agente := agents.NewSimpleAgent(3,3, knowledge)

  agente.LookForGoal(*enviroment, true) 
}
