// Correspondiente a clases de agentes
package agentes

import (
  "fmt"
  "time"
)

const (
  wall = iota + 1
  goal
)

// SimpleAgent representa un agente simple
type SimpleAgent struct{
  x, y int
  knowledge map[[4]int]func(a *SimpleAgent)
  ambientPerception [4]int
  currentGoal int
}

// NewSimpleAgent crea un nuevo agente simple
func NewSimpleAgent(x,y int, knowledge map[[4]int]func(a *SimpleAgent)) *SimpleAgent {
  return &SimpleAgent{
    x,
    y,
    knowledge,
    [4]int{},
    0,
  }
}

// MoveUp, MoveRight, MoveDown y MoveLeft son los actuadores del agente
func (a *SimpleAgent) MoveUp() {
  a.x--
}
func (a *SimpleAgent) MoveRight() {
  a.y++
}
func (a *SimpleAgent) MoveDown() {
  a.x++
}
func (a *SimpleAgent) MoveLeft() {
  a.y--
}

// generatePerception genera una percepción basada en el entorno
func (a *SimpleAgent) generatePerception(env enviroment) {
  // Simulan los sensores
  moveUp := a.x - 1
  moveRight := a.y + 1
  moveDown := a.x + 1
  moveLeft := a.y - 1

  const (
    Up = iota
    Right
    Down
    Left
  )

  // Se genera una percepción
  if moveUp >= 0 && env.matrix[moveUp][a.y] != wall {
    a.ambientPerception[Up] = 1 
  } else {
    a.ambientPerception[Up] = 0
  }
  if moveRight < len(env.matrix[0]) && env.matrix[a.x][moveRight] != wall {
    a.ambientPerception[Right] = 1
  } else {
    a.ambientPerception[Right] = 0
  }
  if moveDown < len(env.matrix) && env.matrix[moveDown][a.y] != wall {
    a.ambientPerception[Down] = 1 
  } else {
    a.ambientPerception[Down] = 0
  }
  if moveLeft >= 0 && env.matrix[a.x][moveLeft] != wall {
    a.ambientPerception[Left] = 1 
  } else {
    a.ambientPerception[Left] = 0
  }
  if env.matrix[a.x][a.y] == goal {
    a.currentGoal++
  } 
}

// generateAction genera una acción basada en la percepción actual
func (a *SimpleAgent) generateAction(env enviroment) {
  // Por medio de los actuadores se generó una acción usando el comportamiento
  if action, exists := a.knowledge[a.ambientPerception]; exists {
    env.matrix[a.x][a.y] = 0 
    action(a)
  } else {
    fmt.Printf("No se encontró una acción para la percepción: %v\n", a.ambientPerception)
  }
}

// LookForGoal busca la meta en el entorno
func (a *SimpleAgent) LookForGoal(env enviroment, display bool) bool {
  counter, maxIterations := 0, 10

  if display {
      env.printPath(a, &counter, &maxIterations)
  } else {
    for ; a.currentGoal < env.totalGoal && counter < maxIterations; counter++ {
      a.generatePerception(env)
      a.generateAction(env)
    }
  }

  if a.currentGoal == env.totalGoal {
    fmt.Println("¡Se han encotrado todos los obejtivos!")
    return true
  } else if counter == maxIterations {
    fmt.Printf("No se encontró la meta; posición final: (%d, %d)\n", a.x, a.y)
    return false
  }

  return false
}

// enviroment representa el entorno del agente
type enviroment struct {
  matrix [][]int
  totalGoal int
}

// NewEnviroment crea un nuevo entorno
func NewEnviroment(matrix [][]int, totalGoal int) *enviroment {
  if totalGoal == -1 {
    totalGoal++
    for _, row := range matrix {
      for _, value := range row {
        if value == goal {
          totalGoal++
        }
      }
    }
  }
  return &enviroment{
    matrix,
    totalGoal,
  }
}

// printPath imprime el entorno y la percepción del agente
func (env *enviroment) printPath(a *SimpleAgent, counter, maxIterations *int) {
  // Iteración inicial
  a.generatePerception(*env)
  env.matrix[a.x][a.y] = 3
  fmt.Print("\033[H\033[2J") 
  fmt.Println("Iteración:", *counter)
  for _, row := range env.matrix {
    fmt.Println(row)
  }
  fmt.Println("Percepción:", a.ambientPerception) 
  time.Sleep(1 * time.Second)

  for ; a.currentGoal < env.totalGoal && *counter < *maxIterations; *counter++ {
    a.generateAction(*env)
    a.generatePerception(*env)
    env.matrix[a.x][a.y] = 3

    fmt.Print("\033[H\033[2J") 
    fmt.Println("Iteración:", *counter)
    for _, row := range env.matrix {
      fmt.Println(row)
    }
    fmt.Println("Percepción:", a.ambientPerception) 
    time.Sleep(1 * time.Second)
  }
}
