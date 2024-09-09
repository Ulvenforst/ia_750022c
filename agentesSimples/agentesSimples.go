package agentesSimples

import (
  "fmt"
  "time"
)

const (
  wall = iota + 1
  goal
)

type SimpleAgent struct{
  x, y int
  knowledge map[[5]int]func(a *SimpleAgent)
  ambientPerception [5]int
}
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

func NewSimpleAgent(x,y int, knowledge map[[5]int]func(a *SimpleAgent)) *SimpleAgent {
  return &SimpleAgent{
    x,
    y,
    knowledge,
    [5]int{},
  }
}

type Enviroment struct {
  matrix [][]int
}
func NewEnviroment(matrix [][]int) *Enviroment {
  return &Enviroment{
    matrix,
  }
}

func (a *SimpleAgent) GeneratePerception(Enviroment Enviroment) {
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
    Current
  )

  // Se genera una percepción
  if moveUp >= 0 && Enviroment.matrix[moveUp][a.y] != wall {
    a.ambientPerception[Up] = 1 
  } else {
    a.ambientPerception[Up] = 0
  }
  if moveRight < len(Enviroment.matrix[0]) && Enviroment.matrix[a.x][moveRight] != wall {
    a.ambientPerception[Right] = 1
  } else {
    a.ambientPerception[Right] = 0
  }
  if moveDown < len(Enviroment.matrix) && Enviroment.matrix[moveDown][a.y] != wall {
    a.ambientPerception[Down] = 1 
  } else {
    a.ambientPerception[Down] = 0
  }
  if moveLeft >= 0 && Enviroment.matrix[a.x][moveLeft] != wall {
    a.ambientPerception[Left] = 1 
  } else {
    a.ambientPerception[Left] = 0
  }
  if Enviroment.matrix[a.x][a.y] == goal {
    a.ambientPerception[Current] = 1
    fmt.Println("Goal!")
  } 
}


func (a *SimpleAgent) GenerateAction(env Enviroment) {
  // Por medio de los actuadores se generó una acción usando el comportamiento
  if action, exists := a.knowledge[a.ambientPerception]; exists {
    action(a)
  }
}

func (a *SimpleAgent) LookForGoal(env Enviroment) bool {
  counter, maxIterations := 0, 20

  for ; a.ambientPerception[4] != 1 && counter < maxIterations; counter++ {
    a.GeneratePerception(env)
    a.GenerateAction(env)
  }

  if counter == maxIterations {
    fmt.Println("No se encontró la meta")
    return false
  }

  return true
}

// * Todo lo que dice "Visualización" no es necesario para el funcionamiento del agente
func (a *SimpleAgent) VisualizePath(env Enviroment) bool {
  counter, maxIterations := 0, 20
  env.matrix[a.x][a.y] = 3 // Visualización: se pinta la posición inicial

  for ; a.ambientPerception[4] != 1 && counter < maxIterations; counter++ {
    if a.ambientPerception[4] == 1 {
      fmt.Println("Goal!")
      break
    }
    fmt.Print("\033[H\033[2J") // Visualización: limpia la consola
    fmt.Println("Iteración:", counter) // Visualización: imprime la iteración
    for _, row := range env.matrix { // Visualización: imprime la matriz
      fmt.Println(row)
    }

    a.GeneratePerception(env)
    env.matrix[a.x][a.y] = 0 // Visualización: se borra la posición actual
    a.GenerateAction(env)
    fmt.Println("Percepción:", a.ambientPerception) // Visualización: imprime la percepción
    if env.matrix[a.x][a.y] != goal { // Visualización: se pinta la nueva posición
      env.matrix[a.x][a.y] = 3   
    }
    time.Sleep(1 * time.Second) // Visualización: pausa de 1 segundo
  }

  if counter == maxIterations {
    fmt.Println("No se encontró la meta")
    return false
  }

  return true
}
