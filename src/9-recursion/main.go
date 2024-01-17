package main

import (
	"fmt"
	"math"
)

func main() {
	TestMazeSolver()
}

type Node struct {
	Value interface{}
	Prev  *Node
}

type Stack struct {
	Length uint
	Head   *Node
}

func Factory() Stack {
	return Stack{Length: 0, Head: nil}
}

// Metemos un nuevo elemento al stack, por lo que ahora la cabeza de nuestro stack va a ser el ultimo elemento que metimos

func (s *Stack) push(item interface{}) {

	// Aumentamos la Length de nuestro stack porque vamos a ingresarle un nuevo elemento

	s.Length++

	// Creamos el nodo del nuevo elemento y de momento su propiedad de prev no la enlazamos con nada

	node := Node{Value: item, Prev: nil}

	if s.Head != nil {
		// Si el head no es nulo, es decir, si existe al menos un elemento en nuestro stack, hacemos que el prev del nuevo nodo sea el head actual del stack
		node.Prev = s.Head
	}

	// Actualizamos el head del stack para que ahora apunte al nuevo nodo que ya  se encuentra enlazado con el anterior head
	s.Head = &node

}

func (s *Stack) pop() interface{} {
	s.Length = uint(math.Max(0, float64(s.Length)-1))
	// Aqui lo que hacemos es que retornamos el numero mas grande entre dos numeros

	// Si nuestra Length - 1 es menor a 0, la Length se mantendra en 0, pero si no es asi, es decir nuestra lenght menos uno es mayor a 0 pues el valor de nuestra length se actualiara a ese

	// Esto sirve para evitar tener que poner un if de que si nuestra lenght es 0 no se haga la operación

	// if (s.Length != 0) {
	//     s.Length--
	// }

	// Guardamos la head actual en una variable
	oldHead := s.Head

	//(A) <- (B) <- (C) <- (D) <- (E) <- (F)
	//                                   ^
	//                                   head

	// Si la Length despues de actualizarla es 0, signfica que vamos a vaciar nuestro stack, poir lo que hacemos que simplemente la head este apuntando a nil
	if s.Length == 0 {
		// (A)
		// ^
		// head

		// nil   (A)
		// ^
		// head
		s.Head = nil
	}

	// Si nuestra head no es nil, signfica que podemos sacar elementos de nuestro stack porque tiene elementos, por lo que, hacemos que nuestro head ahora apunte al elemento previo al head antiguo, y desenlazamos al head antiguo de nuestro stack
	if s.Head != nil {
		// Estado actual de nuestro stack
		//(A) <- (B) <- (C) <- (D) <- (E) <- (F)
		//                                   ^
		//                                   head
		s.Head = oldHead.Prev
		// Hacemos que head apunte ahora al elemento previo al que se le va a hacer pop
		//(A) <- (B) <- (C) <- (D) <- (E) <- (F)
		//                            ^
		//                            head
		oldHead.Prev = nil
		// Desenlazamos el ultimo nodo (el que queremos sacar)
		//(A) <- (B) <- (C) <- (D) <- (E)    (F)
		//                            ^
		//                            head
	}

	return oldHead.Value
	// Devolvemos el valor del elemento que sacamos
}

// Retornamos el valor contenido en el head de nuestro stack
func (s *Stack) peek() interface{} {
	return s.Head.Value
}
func TestMazeSolver() {
	// Define a sample maze
	maze := [][]string{
		{"#", "#", "#", "#", "#", "#", "#", "#"},
		{"#", " ", " ", " ", " ", " ", " ", "#"},
		{"#", " ", " ", " ", " ", " ", " ", "#"},
		{"#", " ", " ", " ", " ", " ", " ", "#"},
		{"#", " ", " ", " ", " ", " ", " ", "#"},
		{"#", " ", " ", " ", " ", " ", " ", "#"},
		{"#", " ", " ", " ", " ", " ", " ", "#"},
		{"#", "#", "#", "#", "#", "#", "#", "#"},
	}

	// Define start and end points
	start := NewPoint(2, 2)
	end := NewPoint(3, 3)

	// Define the expected path
	// expectedPath := []Point{
	// 	{0, 0}, {0, 1}, {0, 2}, {0, 3}, {1, 3}, {2, 3}, {2, 4}, {2, 5}, {2, 6}, {2, 7},
	// 	{3, 7}, {4, 7}, {5, 7}, {6, 7}, {7, 7},
	// }

	// Solve the maze
	result := MazeSolver(maze, "#", start, end)
	curr := result.Head
	for curr != nil {
		fmt.Println(curr)
		curr = curr.Prev
	}

}

type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) Point {
	return Point{X: x, Y: y}
}

var dir = []Point{
	NewPoint(0, 1),
	NewPoint(-1, 0),
	NewPoint(0, -1),
	NewPoint(1, 0),
}

func Walk(maze [][]string, wall string, curr Point, end Point, seen [50][50]bool, path *Stack) bool {

	if curr.X < 0 || curr.X >= len(maze[0]) || curr.Y < 0 || curr.Y >= len(maze) {
		return false
	}
	if maze[curr.Y][curr.X] == wall {
		return false
	}
	if curr.X == end.X && curr.Y == end.Y {
		return true
	}
	if seen[curr.Y][curr.X] {
		return false
	}
	seen[curr.X][curr.Y] = true
	path.push(curr)
	for i := 0; i < len(dir); i++ {
		x := dir[i].X
		y := dir[i].Y
		if Walk(maze, wall, NewPoint(curr.X+x, curr.Y+y), end, seen, path) {
			return true
		}
	}
	return false
}
func MazeSolver(maze [][]string, wall string, start Point, end Point) *Stack {
	var seen [50][50]bool
	for i := 0; i < len(seen[0]); i++ {
		for j := 0; j < i; j++ {
			seen[i][j] = false
		}
	}
	path := Factory()
	Walk(maze, wall, start, end, seen, &path)
	return &path
}
