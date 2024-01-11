package main

import "fmt"

func main()  {
    fmt.Println(LinearSearch([]int{1,2,3}, 3))
}

// Al parecer no es bueno retornar adentro de un loop pero en algoritmos no hay tanto pedo buscar luego que show

func LinearSearch(haystack [] int, needle int) bool  {
    for _, v := range haystack {
        if (v == needle) {
            return true
        }
    }
    return false
}
