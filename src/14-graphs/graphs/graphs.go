package graphs

import (
	"fmt"

	"marcos.com/algorithms/src/12-trees/arraylist"
)

type Matrix [][]int
type List [][]Edge

type Edge struct {
	Value  int
	To     int
	Weight int
}

func walk(graph List, curr int, needle int, path *arraylist.ArrayList, seen *arraylist.ArrayList) bool {
	// Esta es la funcion donde vamos a hacer la recursividad, es decir en la que vamos a visitar los nodos
	// Aqui planteamos nuestros base cases
	//  - Si encontramos el needle
	//  - Si ya habiamos visitado ese nodo

	if curr == needle {
		return true
	}
	if curr == seen.ArrayList[curr] {
		return false
	}

	// Recurse

	seen.ArrayList[curr] = true
	// Prev
	path.Push(curr)
	// Recurse
	// Nuestro grafo es una matriz pero no matriz es una lista de nuestro Edges de nuestros grafos
	list := graph[curr]
	for i := 0; i < len(list); i++ {
		edge := list[i]
		if walk(graph, edge.To, needle, seen, path) {
			return true
		}
	}
	// Post
	path.Pop()

	return false
}

func DFSAdjacentList(graph List, source int, needle int) arraylist.ArrayList {
	seen := arraylist.NewArrayList(10)
	path := arraylist.NewArrayList(10)
	walk(graph, source, needle, &seen, &path)
	return path
}

func BFSMatrix(graph Matrix, source int, needle int) arraylist.ArrayList {

	// queue := queues.NewQueue()
	queue := arraylist.NewArrayList(10)
	seen := arraylist.NewArrayList(7)
	prev := arraylist.NewArrayList(20)
	seen.Fill(false)
	prev.Fill(-1)

	seen.ArrayList[source] = true
	queue.Enqueue(source)

	var curr int

	//  NODOS
	//  0  1  2  3  4  5  6
	// 0 {0, 3, 1, 0, 0, 0, 0},
	// 1 {0, 0, 0, 0, 1, 0, 0},
	// 2 {0, 0, 7, 0, 0, 0, 0},
	// 3 {0, 0, 0, 0, 0, 0, 0},
	// 4 {0, 1, 0, 5, 0, 2, 0},
	// 5 {0, 0, 18, 0, 0, 0, 1},
	// 6 {0, 0, 0, 1, 0, 0, 1},

	for queue.Length > 0 {

		curr = queue.Pop()

		if curr == needle {
			break
		}

		adjs := graph[curr]

		for i := 0; i < len(adjs); i++ {

			if adjs[i] == 0 {
				continue
			}

			if seen.ArrayList[i] == true {
				continue
			}

			// Aqui registramos los nodos que ya hemos visitado
			seen.ArrayList[i] = true
			// En el arreglo de prev pondremos los nodos padres en los que aparecen los nodos, es decir, si en el nodo 5 aparece el nodo 6, para registrarlo lo que se hara sera poner en el indice 6 el nodo 5, si el nodo 4 5 aparece en el nodo 4 lo qeu hacemos es poner en el indice 5 el nodo 4, y asi
			// Eso funciona porque lo que hacemos en este ciclo es recorrerlo n veces donde n es la cantidad de nodos que tenemos, por lo que podemos hacer el bookkeeping asignado a cada nodo de esa manera, de forma que podemos asignarle a los indices (nodos) el valor de los nodos donde aparecen por primera vez
			// Aqui lo que hacemos poner en nuestra lista los Edges(enlaces) a otros nodos y los ponemos en el indice del nodo que aparecen, por ejemplo
			// Este es nuestro prev
			//
			// -1 0 0 4 1 4 5
			// 0  1 2 3 4 5 6
			prev.ArrayList[i] = curr

			// Como no hemos visto el nodo en el que estamos iterando ahora mismo pues lo metemos a la cola
			queue.Push(i)
		}
	}

	curr = needle
	out := arraylist.NewArrayList(50)

	for prev.ArrayList[curr] != -1 {
		out.Push(curr)
		curr = prev.ArrayList[curr].(int)
	}
	return out.Reverse()
}

func TestBFSMatrix() {
	//
	// 0 0 4 1 4 5->6
	//           ^
	//

	// -1 0 0 4 1 4 5
	// 0  1 2 3 4 5 6
	graph := Matrix{
		//   0  1  2  3  4  5  6
		// Primera iteracion, nuestro source es el nodo 0 y nuestro needle es el nodo 6, por lo que vamos a ver los nodos con los que tiene conexiones
		// {0, 3, 1, 0, 0, 0, 0},  // 0
		// Vemos que el nodo 0 esta enlazado al nodo 1 y 2, por lo que a los nodos 1 y 2 los marcaremos como su padre en nodo 0, y marcaremos a los nodos 1 y 2 como visitados ya
		// [f,t,t,f,f,f,f]
		//  0 1 2 3 4 5 6
		// [-1, 0, 0,-1,-1,-1,-1]
		//  0   1  2  3  4  5  6
		// Metemos a la cola a los nodos que visitamos (1 y 2)
		// Ahora siguiente iteracion, que en este caso le tocara al nodo 1 porque es el que sigue en la cola

		// {0, 0, 0, 0, 1, 0, 0},  // 1
		//  0  1  2  3  4  5  6
		// Recorremos sus enlaces y vemos que solamente tiene un enlace con otro nodo que es el nodo 4, checamos en nuestro arraylist de seen si el nodo 4 (que es el indice 4) esta visitado
		//          v
		// [f,t,t,f,f,f,f]
		//  0 1 2 3 4 5 6
		// Vemos que no esta visitado por lo que podemos agregarlo a la cola y asignarle que su padre fue el nodo 1
		// [f,t,t,f,t,f,f]
		//  0 1 2 3 4 5 6
		// [-1, 0, 0,-1, 1,-1,-1]
		//  0   1  2  3  4  5  6
		// Ahora metemos el nodo 4 a nuestra cola
		// Siguiente iteracion 2, recordemos que estaba el nodo 2 antes que el 4
		// {0, 0, 7, 0, 0, 0, 0},  // 2
		//  0  1  2  3  4  5  6
		// Recorriendo el arreglo nos damosa cuenta que solamente tiene un enlace con otro nodo que es el nodo 2, checamos en nuestro arreglo seen para ver si ya lo hemos visitado
		// [f,t,t,f,t,f,f]
		//  0 1 2 3 4 5 6
		// Como ya lo hemos visitado nos saltamos este, y como no tenemos mas enlaces no tenemos nada mas que agregar a la cola, siguiente iteracion

		// La siguiente iteracion en nuestra cola es la del nodo 4
		// {0, 1, 0, 5, 0, 2, 0},  // 4
		//  0  1  2  3  4  5  6
		// Recorremos el nodo y vemos que tiene 3 enlaces, con el nodo 1, el nodo 3 y el nodo 5, como estamos haciendo el recorrido en un ciclo checaremos si ya hemos visitado esos nodos
		// i == 1
		// [f,t,t,f,t,f,f]
		//  0 1 2 3 4 5 6
		// Ya hemos visitado al nodo 1 por lo que saltamos al siguiente
		// i == 3
		// [f,t,t,f,t,f,f]
		//  0 1 2 3 4 5 6
		// Aun no hemos visitado el nodo 3 por lo que  podremos agregarlo a nuestra cola y marcar a 4 como su padre
		//
		// [-1, 0, 0, 4, 1,-1,-1]
		//  0   1  2  3  4  5  6
		// [f,t,t,t,t,f,f]
		//  0 1 2 3 4 5 6
		// Y agregamos al nodo 3 a nuestra cola
		// Ahora vamos a las siguiente iteraciones
		// i == 5
		// Checamos si hemos visitado el nodo 5
		// [f,t,t,t,t,f,f]
		//  0 1 2 3 4 5 6
		// Como aun no lo hemos visitado pues podemos meterlo a la cola y marcar al nodo 4 como su padre
		// [-1, 0, 0, 4, 1, 4,-1]
		//  0   1  2  3  4  5  6
		// [f,t,t,t,t,t,f]
		//  0 1 2 3 4 5 6
		// Ahora agregamos al nodo 5 a la cola
		// El siguiente elemento para sacar de nuestra cola es el nodo 3, pero pues como este no tiene enlace con otro nodo pos lo mandamos a la cola
		// {0, 0, 0, 0, 0, 0, 0},  // 3

		// Ahora el siguiente nodo en la cola es el nodo  5
		// {0, 0, 18, 0, 0, 0, 1}, // 5
		// Vemos que tiene un enlace con el nodo 3 pero si checamos nuestro bookeeper el nodo 3 ya lo hemos visitado
		// [f,t,t,t,t,t,f]
		//  0 1 2 3 4 5 6
		// Asi que seguimos iterando hasta que llegamos al ultimo enlace del nodo 5 con el nodo 6, como no hemos visitado el nodo 6 podemos hacer el procedimiento
		// Marcamos al nodo 5 como el padre del nodo 6
		// [-1, 0, 0, 4, 1, 4, 5]
		//  0   1  2  3  4  5  6
		// Lo marcamos como visitado en nuestro arreglo de seen
		// [f,t,t,t,t,t,t]
		//  0 1 2 3 4 5 6
		// Y lo agregamos a la cola
		// Ahora,
		// Cada vez que hemos sacado un elemento de la cola hemos verificado si este elemento es iugal al que estamos buscando, en este ejemplo el que estabamos buscando era el nodo 6, por lo que rompemos el ciclo y podemos retornar el camino que se puede seguir para llegar al nodo 6 obtenido por breadth first search
		// if curr == needle {
		// 	break
		// }

		// Lo que queremos hacer ahora mismo es el camino desde nuestro needle a nuestro source, por lo que la manera de hacerlo es ir a nuestro arreglo en el que teniamos los padres de cada nodo, nuestro prev, en el indice 6 (nuestro needle en este caso) ahi se encontrara el padre de nuestro nodo.

		// curr = needle
		// out := arraylist.NewArrayList(50)

		// En este momento nuestro needle es el nodo 6 y su padre es el nodo 5, por lo que lo que pasara en la primera iteracion de este ciclo que se detendra cuando no haya un padre (es decir que ningun otro nodo haya hecho referencia a el) sera lo siguiente

		// Pushearemos nuestro curr que lo cambiamos para que sea nuestro needle, luego haremos que el curr sea el nodo padre de nuestro needle, que recordemos que para hacer eso tenemos que ir a nuestro arreglo de prev que contiene los padres de los nodos en los indices que les corresponden a sus hijos, en el indice 6 se encuentra el 5. Luego en la siguiente iteracion se pondra a 5 en out, y buscaremos el padre del nodo 5 para que este ahora sea curr, que el padre del nodo 5 es el nodo 4, lo metemos al out, y vamos al nodo padre del 4 que es el 1, y boom lo metemos, y ya se acaba el ciclo, pero por que si falta el nodo padre del 1 que es el 0, si pero el ciclo sigue solamente si el padre del current no es -1, y cuando ya metemos el 1 a oout cambiamos a curr para que sea el padre del 1, el cual es 0, y al revisar el padre del nodo 0 como su padre es -1 pues no va a ejecutarse el codigo, por lo que tenemos que agregarlo a mano

		// for prev.ArrayList[curr] != -1 {
		// 	out.Push(curr)
		// 	curr = prev.ArrayList[curr].(int)
		// }
		// out.Push(source)
		// return out.Reverse()
		{0, 3, 1, 0, 0, 0, 0},  // 0
		{0, 0, 0, 0, 1, 0, 0},  // 1
		{0, 0, 7, 0, 0, 0, 0},  // 2
		{0, 0, 0, 0, 0, 0, 0},  // 3
		{0, 1, 0, 5, 0, 2, 0},  // 4
		{0, 0, 18, 0, 0, 0, 1}, // 5
		{0, 0, 0, 1, 0, 0, 1},  // 6
	}

	source := 0
	needle := 6

	r := BFSMatrix(graph, source, needle)

	expected := arraylist.NewArrayList(10)
	expected.Push(0)
	expected.Push(1)
	expected.Push(4)
	expected.Push(5)
	expected.Push(6)
	fmt.Println("Expected:", expected.String())
	fmt.Println("Got:", r.String())
	fmt.Println(r.String() == expected.String())
}
