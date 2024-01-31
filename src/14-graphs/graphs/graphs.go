package graphs

import (
	"fmt"

	"marcos.com/algorithms/src/12-trees/arraylist"
	"marcos.com/algorithms/src/12-trees/queues"
)

type Matrix [][]int

func BFSMatrix(graph Matrix, source int, needle int) arraylist.ArrayList {
	// Creamos nuestra cola que va a servir para recorrer los enlaces de cada grafo
	queue := queues.NewQueue()
	// Creamos nuestro arreglo para guardar los elementos que hayamos visto
	seen := arraylist.NewArrayList(uint(len(graph[0])))
	// Creamos nuestro arreglo para almacenar el path que vamos a tomar
	path := arraylist.NewArrayList(uint(len(graph[0])))
	// Llenamos nuestro arreglo de seen de falses
	seen.Fill(false)
	// Llenamos nuestro arreglo de path con -1
	path.Fill(-1)

	// Ponemos que el nodo en el que iniciamos el BDS dentro de los nodos visitados
	seen.ArrayList[source] = true
	// Metemos a la cola a nuestro source
	queue.Enqueue(source)

	// Creamos la variable para guardar los nodos que vamos a estar haciendoles dequeue durante el ciclo
	var curr int

	for queue.Length > 0 {
		// Hacemos dequeue a nuestro elemento en head actualmente
		curr = queue.Dequeue().(int)
		if curr == needle {
			// Si es igual al que estamos buscando ya hemos logrado nuestra busqueda por lo que rompemos el ciclo
			break
		}

		// En el caso de que no sea asi significa que tenemos que buscar en los enlaces del nodo actual
		// Hacemos un ciclo que va a recorrer los enlaces que tiene cada vertice de la siguiente manera
		//
		// [
		//          0  1  2  3  4
		//     v=0:[0, 1, 4, 5, 0],
		//     v=1:[1, 0, 0, 0, 0],
		//     v=2:[0, 0, 0, 2, 0],
		//     v=3:[0, 0, 0, 0, 5],
		//     v=4:[0, 0, 0, 0, 0],
		// ]
		for i := 0; i < len(graph[curr]); i++ {
			if graph[curr][i] == 0 {
				continue
			}
			if seen.ArrayList[i] == true {
				continue
			}
			// Por que seen es un array y no una matriz?
			// Porque el array es el representante de todos los nodos que hay, por lo que si ya visitaramos el 3 no ocupariamos visitarlo, por ejemplo este grafo tiene 5 nodos (0,1,2,3,4) y si empezaramos a recorrerlo desde 0, visitariamos los nodos 1,2,3, y marcariamos como vistos esos 3 nodos, luego se terminaria este ciclo e iniciaria el otro, sacariamos al head de la cola en ese punto que seria 1, entonces, entrariamos a este ciclo, y checariamos, 1 tiene solamente un enlace, con el nodo 0 con un peso de 1, pero como este ya se encuentra visitado (Recordemos que antes sde iniciar el ciclo hicimos esto seen.ArrayList[source] = true), por lo que el vamos a la siguiente iteracion, pero como no tenemos mas enlaces en ese array de la matriz se terminara el ciclo. Luego iremos a 2, el cual tiene 2 enlaces pero estos ya se encuentran en nuestro arreglo de seen, es decir, son true en la posicion de 0 y 3, entonces vamos al siguiente nodo, el 3, el cual tiene un enlace con 0 pero este ya se encuentra en nuestro arreglo de seen, por lo que nos saltamos hasta que haya un nodo, y llegamos al nodo 4, con el que tiene un enlace con un peso de 5 y este no se encuentra visto (seen[4] == false en ese punto), por lo que lo registramos (tambien en nuestro path y nuestra cola), y luego volvemos al ciclo principal, sacamos al head de nuestra cola actualmente que seria ya 4 pero como este no tiene ningun enlace pues vamos a simplemente hacer continue en cada ciclo hasta que se acabe
			//                        [true, true, true, true, false]
			//                     v = 0     1     2     3     4
			//     1
			// (0)---(1)
			//  | \
			//  |  \
			// 4|   \ 5
			//  |    \
			//  |     \
			// (2)----(3)------(4)
			//     2       5
			seen.ArrayList[i] = true
			fmt.Println(curr)
			path.ArrayList[i] = curr
			queue.Enqueue(i)
		}
	}
	curr = needle
	out := arraylist.NewArrayList(10)

	fmt.Println(seen.String())
	fmt.Println(path.String())

	if path.ArrayList[curr] != -1 {
		for path.ArrayList[curr] != -1 {
			out.Push(curr)
			curr = path.ArrayList[curr].(int)
		}
		out.Push(source)
	}

	return out
}

func TestBFSMatrix() {
	graph := Matrix{
		{0, 1, 4, 5, 0},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 2, 0},
		{0, 0, 0, 0, 5},
		{0, 0, 0, 0, 0},
	}

	source := 0
	needle := 4

	result := BFSMatrix(graph, source, needle)

	expected := arraylist.NewArrayList(10)
	expected.Push(0)
	expected.Push(1)
	expected.Push(4)

	if !result.Equals(&expected) {
		fmt.Printf("BFSMatrix(%v, %d, %d) = %v, expected %v", graph, source, needle, result, expected)
	}
}
