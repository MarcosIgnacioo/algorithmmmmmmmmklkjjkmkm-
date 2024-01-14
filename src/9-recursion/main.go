package main

import "fmt"

func main() {
	fmt.Println("hola")
}

type Point struct {
	X int
	Y int
}

func newPoint(x int, y int) Point {
	return Point{X: x, Y: y}
}

func Walk(maze [][]string, wall string, current Point, end Point, seen [][]bool) bool {
	// Base case 1 Not out of bounds
	if current.X < 0 || current.X > len(maze[0]) || current.Y < 0 || current.Y > len(maze) {
		return false
	}

	// Base case 2 Is this a wall
	if maze[current.Y][current.X] == wall {
		return false
	}

	// Base case 3 Have we seen it
	if seen[current.Y][current.X] == true {
		return false
	}

	// Base case 4 Is this the end
	if current.X == end.X && current.Y == end.Y {
		return true
	}

	seen[current.Y][current.X] = true

	return true
}

func MazeSolver(maze [][]string, wall string, start Point, end Point) []Point {

}
