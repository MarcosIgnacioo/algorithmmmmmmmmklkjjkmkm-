package main

import (
	"fmt"
	"marcos.com/algorithms/src/recursionyay/arraybuffer"
)

func main() {
	maze := [][]string{
		//        0    1    2    3
		/* 0 */ {"x", " ", "x", " ", " "},
		/* 1 */ {" ", " ", " ", " ", " "},
		/* 2 */ {"x", "x", "x", "x", " "},
	}

	// Define start and end points
	start := NewPoint(1, 1)
	end := NewPoint(1, 3)
	asdf := MazeSolver(maze, "x", start, end)
	fmt.Println(*asdf.Array...)
}

var dir = []Point{
	NewPoint(0, 1),
	NewPoint(1, 0),
	NewPoint(0, -1),
	NewPoint(-1, 0),
}

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("[%v, %v]", p.Y, p.X)
}

func NewPoint(x int, y int) Point {
	return Point{X: x, Y: y}
}

func Walk(maze [][]string, wall string, curr Point, end Point, seen [50][50]bool, path arraybuffer.ArrayBuffer) bool {
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
	seen[curr.Y][curr.X] = true
	path.Push(curr)

	for i := 0; i < len(dir); i++ {
		x := dir[i].X + curr.X
		y := dir[i].Y + curr.Y
		if Walk(maze, wall, NewPoint(x, y), end, seen, path) {
			return true
		}
	}
	return false
}

func MazeSolver(maze [][]string, wall string, start Point, end Point) *arraybuffer.ArrayBuffer {
	path := arraybuffer.NewArrayBuffer(67)
	var seen [50][50]bool
	for i := 0; i < len(seen[0]); i++ {
		for j := 0; j < i; j++ {
			seen[i][j] = false
		}
	}
	Walk(maze, wall, start, end, seen, path)
	return &path
}
