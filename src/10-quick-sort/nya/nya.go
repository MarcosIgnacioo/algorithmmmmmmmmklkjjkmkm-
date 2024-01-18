package nya

import (
	"fmt"
)

func main() {
	array := []int{9, 7, 2, 3, 4, 1, 5}
	fmt.Println(array)
	QuickSort(array)
	fmt.Println(array)
}

func QuickSort(array []int) {
	QS(array, 0, len(array)-1)
}

// Funcion controladora de la recursividad

// Vamos a tener de parametros 3 cosas, el arreglo, high y low

//  El arreglo es el arreglo que vamos a acomodar

//  El low es nuestro inicio de la particion

//  El high es el final de nuestra particion

// En la primera vez que llamemos a esta funcion el low sera 0 (inicio del arreglo) y el high el indice del ultimo elemento de nuestro arreglo, es decir de extremos a extremos

func QS(array []int, low int, high int) {
	// Definiremos nuestro base case y nuestros casos en los que vamos a recursar
	if low >= high {
		// En el caso de que nuestro low ya haya alcanzado a nuestro high, es decir, que en las particiones que vayamos haciendo se reduzcan a ser 1,1 2,2 3,3 etc
		return
	}
	pivotIdx := Partition(array, low, high)
	// Utilizamos el index del pivot para desplazarnos a la izquierda en el arreglo para crear una nueva particion. Y tomamos eso como nuestro high
	QS(array, low, pivotIdx-1)
	// Utilizamos el index del pivot para desplazarnos a la derecha en el arreglo para crear una nueva particion, y tomamos eso como nuestro low para poder llenar el otro lado de nuestro arreglo y llamamos a QS para que verifique nuestro base case y ya cuando el low haya alcanzado al high significara que nuestro arreglo esta completamente ordenado porque ya no quedan mas elementos que ordenar
	QS(array, pivotIdx+1, high)
}

// Se encarga de pasar los elementos a la izquierda del pivot que seleccionemos en la particion utilizando idx como su bookeeper para llevar anotado lasposiciones en el arreglo que se deben de ir llenando
func Partition(array []int, low int, high int) int {
	// En esta funcion lo que vamos a hacer es leer de extremo a extremo de la particion, y vamos a ordenar en base al pivot que le corresponda a la particion

	// Agarramos nuestro pivot, que sera el ultimo elemento de la particion, si es la primera vez que hace la particion (low = 0 high = len(array)-1), el pivot sera el ultimo elemento del arreglo
	pivot := array[high]
	// Creamos nuestro "buffer" que nos servira para hacer los swaps entre los elementos del arreglo cuando se cumpla la condicion de que sean menores o iguales al pivot
	idx := low - 1

	for i := low; i < high; i++ {
		if array[i] <= pivot {
			idx++
			tmp := array[i]
			array[i] = array[idx]
			array[idx] = tmp
		}
	}
	// En nuestro ciclo lo que hacemos es simplemente recorrer la particion correspondiente a nuestro arreglo por ejemplo
	// i = 13  14  15  16  17  18  19  20
	// Supongamos que tenemos esta particion
	// EN este caso nuestro high seria el indice correspondiente en el arreglo al elemento 44, y nuestro low el indice correspondiente al elemento 36, (low = 13 high = 20). Entonces empezamos recorriendo nuestra particion, checamos que el elemento en i (como el ciclo inicia en low pues sera el primer elemento de nuestra particion) 36, como si es menor a 44 (nuestro pivot) hacemos el swap de nuestro elemento en i a la posicion de nuestro index buffer (cuya funcion es actuar de bookeeper e ira guardandando la siguiente posicion disponible para acomodar a nuestros elementos) y despues de hacer idx++ idx vale 13, no suena tan emocionante un swap entre 13 y 13 pero ahi va, luego el siguiente elemento es 46(i = 14) en el indice 14, este no cumple con nuestra condicion por lo que no tenemos que realizar el registro de que va a ocupar un lugar en nuestro bookeping, pero el siguiente elemento es 10(i = 15) y este si que cumple con nuestra funcion, entonces entramos a nuestro if, se hace el registro en el bookeeper  ocn el idx++ y vemos que idx ahora vale 14, entonces ahora si que hacer un cambio de elementos parece util, cambiamos al elemento que se encuentre en el indice 14 (46) por el que se encuentra en el indice actual (15), y como estamos en un ciclo pues no tenemos que preocuparnos poer volver a ver al 46 porque ya lo habremos pasado incluso despues del cambio porque vamos a iterar al siguiente numero y ya pues todo eso se hace con el resto de la particion para que quede weakly sorted que queda asi

	// ... 36, 10,  3 , 27, 8 , 19, 46, 44 , ...
	// Y para poder dividir y conquistar, tenemos que acomodar a nuestro pivot en la ultima habitacion disponible de nuestro bookeeping, pues esta esta al lado de nuestro ultimo elemento que es menor o igual a nuestro pivot, asi que realizamos el cambio entre ese elemento y nuestro pivot y vuala ahora nuestro pivot esta en medio, en este ejemplo no se ve tanto pero pues si hubieran mas elementos mayores a 44 habrian mas elementos a la derecha, porque recordemos que el bookeeper se encarga de registrar las posiciones consecutivas que ocuparan los elementos menores o iguales al pivot
	// ... 36, 10,  3 , 27, 8 , 19, 44, 46 , ...
	idx++
	// Y por eso tenemos que sumarle aqui uno porque queremos que este al lado del ultimo elemento que colocamos
	// Ademas que en caso de que no hubiese ningun elemento menor o igual a nuestro pivot, pues el pivot tendria que ponerse en la primera posicion de nuestra particion, por eso tenemos que poner - 1 arriba porque si pusieramos que es simplemente igual al primer indice de nuestra particion y llegaramos a esta parte sin haber puesto ningun elemento a la izquierda de nuestro pivot, pues le tendriamos que sumar 1 y pues eso lo pondria en la segunda posicion y como el ++ aqui es necesario en los otros casos pues hacer que el idx primariamente valga low - 1 hace que todo funcione bien
	array[high] = array[idx]
	array[idx] = pivot
	return idx
	// Retornamos el idx que ahora es el i del pivot
}
