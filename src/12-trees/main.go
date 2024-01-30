package main

import (
	"fmt"

	"marcos.com/algorithms/src/12-trees/tries"
)

func main() {
	t := tries.NewTrie()
	t.Insert("hola")
	t.Insert("holograma")
	t.Insert("holagamer")
	t.Autocompletion("hol")
	fmt.Println()
}
