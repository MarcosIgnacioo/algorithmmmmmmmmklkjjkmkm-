package main

import (
	"fmt"
	"math"
)
func main()  {
    // fmt.Println(BinarySearch([]int{1,2,3,3,4,4,5,6}, 9))
    fmt.Println(Crystal([]bool{false,false,false,false,true,true,true,true}))
}

func BinarySearch(haystack [] int, needle int) bool  {
    low := 0
    high := len(haystack)
    for low < high {
        mid := int(math.Floor(float64(low + (high - low) / 2)))
        value := haystack[mid]
        if value == needle {
            return true
        } else if value > needle{
            high = mid
        } else {
            low = mid + 1
        }
    }
    return false
}


// Como funciona
// Partimos de tener un high y un low, que de base en la primera iteracion el low va a ser 0 mientras que el high sera la longitud del arreglo

// []int{1,2,3,3,4,5,5,6} Buscamos: 4

// En este ejemplo, durante nuestra primera iteracion los valores de low y high serian:

// Low = 0
// High = 8

// Mientras low sea menor a high, es decir mientras aun haya espacio buscable entre ambas haremos el ciclo

// Calculamos la mitad con la formula que es sumar Low + High, el resultado dividirlo entre 2 y como vamos a estar dividiend todo a la mitad ocupamos ajustar el valor a nuestra ordenada al origen que siempre sera Low

// Si el valor que hay en el indice de mid es el mismo que buscamos retornamos true, si el valor a la mitad es MAYOR del que estamos buscando, significa que ocupamos reducir nuestro High por lo que el High ahora sera la mitad
// Sin embargo si el valor a la mitad es menor al que estamos buscando significa que debemos de actualizar nuestro Low a que valga la mitad mas uno porque estariamos ya excluyendo el valor que habiamos encontrado antes (que es menor al que buscamos)

// Si el valor a la mitad es mayor que el needle ocupamos mover nuestro tope de rango de busqueda a la izquierda
// Si el valor a la mitad es menor que el needle ocupamos mover nuestra base de rango de busqueda a la derecha

// Dos bolas de cristal se caen y se rompen en un piso especifico

// Dado un arreglo de booleanos encuentra la primera aparicion de true

// Este problema se puede resolver con O N es decir con una busqueda lineal, pero se puede hacer tambien con una busqueda medio binaria rara theprimeagen style y es partiendo de la base del binary search, partiendo de la base de que el arreglo esta ordenado, recorreremos el arreglo con saltos que equivalen a la raiz cuadrada de N, cuando encontremos un elemento que sea true romperemos ese ciclo y guardaremos el indice en el que se encuentra.

// Ahora lo que hacemos es volver al punto inicial del ultimo intervalo de saltos, es decir el rango en el que se encuentra la primera aparicion del true. Ahora lo que hacemos es simplemente recorrer el arreglo de manera lineal hasta el punto en el que aparecio el true en el primer ciclo, y con el primero  que encuentre pues ya retorna ese indice

// Es como una busqueda lineal pero boosteada

func Crystal(breaks [] bool) int {
    jumpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))
    i := jumpAmount
    for ; i < len(breaks); i+=jumpAmount{
        if breaks[i] {
            break;
        }
    }

    i -= jumpAmount

    for j := i; j < len(breaks); j++ {
        if breaks[j] {
            return j
        } 
    }

    return -1
}
