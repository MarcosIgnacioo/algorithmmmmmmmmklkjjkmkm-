package main

import "fmt"

func main()  {
    array := []int {543,1,3,4,2,5,6,7,2,3,5}
    fmt.Println(BubbleSort(array))
}

func BubbleSort(array []int) []int  {
    N := len(array)
    for i := 0; i < N; i++ {
        for j := 0; j < N - 1 - i; j++ {
            if array[j] > array[j+1] {
                buffer := array[j]
                array[j] = array[j+1]
                array[j+1] = buffer
            }
        }
    }
    return array
}

// Primero pensemos como podemos cambiar el elemento que se encuentra al lado si el elemento en el indice es mayor al elemento en el indice + 1

            // if array[j] > array[j+1] {
            //     buffer := array[j]
            //     array[j] = array[j+1]
            //     array[j+1] = buffer
            // }
// De esta manera lo hacemos

// Ahora metemos esto a un ciclo y voala, nuestro elemento mas grande esta ahora al final de nuestro arreglo, sin embargo queremos que todo el arreglo se acomode, por lo que tenemos que hacer esta operacion N cantidad de veces es decir tendremos que hacerlo con respecto al tamaño del arreglo, que en este caso puede ser 11, porque pues tiene 11 elementos que deben de ser ordenados

        // for j := 0; j < N ; j++ {
        //     if array[j] > array[j+1] {
        //         buffer := array[j]
        //         array[j] = array[j+1]
        //         array[j+1] = buffer
        //     }
        // }

// Entonces esto lo tendremos que poner dentro de otro ciclo haciendo de nuestro algoritmo un ciclo anidado


    // for i := 0; i < N; i++ {
    //     for j := 0; j < N ; j++ {
    //         if array[j] > array[j+1] {
    //             buffer := array[j]
    //             array[j] = array[j+1]
    //             array[j+1] = buffer
    //         }
    //     }
    // }

// Bien pero estamos teniendo ahora un pequeño problema, estamos yendonos out of bounds con el arreglo porque el ultimo elemento de nuestro arreglo (indice de 10 y recordemos que la regla es j < N asi que seria j < 11 por lo que 10 esta todavia incluido) bueno no existe un indice 11 por lo que cuando llegue al indice 10 va a intentar ver un elemento al lado pero como el indice 10 es el ultimo dara error out of bounds arreglemos eso facilmente añadiendo una resta de 1 a nuestra condicional 

    // for i := 0; i < N; i++ {
    //     for j := 0; j < N - 1 ; j++ {
    //         if array[j] > array[j+1] {
    //             buffer := array[j]
    //             array[j] = array[j+1]
    //             array[j+1] = buffer
    //         }
    //     }
    // }

// Y ya casi estamos listos pero como se habia dicho antes el BubbleSort tiene la propiedad interesante de que despues de cada iteracion  el ultimo elemento pues si es el mas grande por lo que estamos haciendo iteraciones que realmente estamos desperdiciando porque esos elementos ya han sido acomodados, en la primera iteracion de la operacion del cambio de lugares (cuando i vale 0) pues ningun elemento esta  acomodado asi que todo bien con revisar el arreglo entero, pero una vez se complete la revision entera del arreglo y se hagan los cambios que se pueden hacer en esa primera iteracion, esta la garantia de que el ultimo elemento del arreglo es el más grande de todos, por lo que, la manera de saltearnos este es recortando hasta donde puede iterar el for de j, simplemente restaremos a nuestra N, aparte del 1 que nos servia para evitar el out of bounds, tambien i, porque asi evitamos intentar hacer el shifting con el ultimo elemento y los elementos que ya hayamos puesto en su lugar correspondniente no tienen que ser revisados de nuevo

    // for i := 0; i < N; i++ {
    //     for j := 0; j < N - 1 - i ; j++ {
    //         if array[j] > array[j+1] {
    //             buffer := array[j]
    //             array[j] = array[j+1]
    //             array[j+1] = buffer
    //         }
    //     }
    // }

// Y ahora si ya quedo nuestro BubbleSort
// Repaso

// El ciclo de i sirve para hacer los cambios N veces porque es las veces que se ocupa leer el arreglo para acomodarlo

// El ciclo de j son los pasos internos, es decir, leer el arreglo desde el inicio hasta el final cambiando elemento por elemento si la condicion se cumple

// Restamos 1 a la N en el ciclo de las j porque no se puede hacer la operacion con el ultimo elemento puesto que no tiene un elemento siguiente con el compararlo

// Restamos i a la N porque queremos evitar revisar los ultimos elementos que por la regla de este algoritmo que por cada vez que se haga una iteracion completa de i el ultimo elemento correspondiente va a estar perfectamente acomodado


