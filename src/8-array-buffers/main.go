package main

import (
	"fmt"
	"reflect"
)

type ArrayBuffer struct {
	Array    *[]interface{}
	Length   uint
	Capacity uint
	Type     reflect.Type
	Head     uint
	Tail     uint
}

func NewArrayBuffer(c uint) ArrayBuffer {
	array := make([]interface{}, c)
	return ArrayBuffer{
		Array:    &array,
		Length:   0,
		Capacity: c,
		Type:     nil,
		Head:     0,
	}
}

func (ab *ArrayBuffer) Enqueue(item interface{}) {
	if ab.Type == nil {
		ab.Type = reflect.TypeOf(item)
		ab.Tail = ab.Capacity
	}

	if ab.Type != reflect.TypeOf(item) {
		return
	}

	if ab.Length+1 > ab.Capacity {
		newArray := make([]interface{}, ab.Capacity*2)
		newArray = appendGamer(newArray, *ab.Array)
		ab.Array = &newArray
	}

	if ab.Length > 0 {

		for i := ab.Head; i > 0; i-- {
			(*ab.Array)[i] = (*ab.Array)[i-1]
		}

		ab.Length++
		(*ab.Array)[ab.Tail] = item
	}
}

func (ab *ArrayBuffer) Push(item interface{}) {

	if ab.Type == nil {
		ab.Type = reflect.TypeOf(item)
		ab.Head = 0
	}

	if ab.Type != reflect.TypeOf(item) {
		return
	}

	if ab.Length+1 > ab.Capacity {
		newArray := make([]interface{}, ab.Capacity*2)
		newArray = appendGamer(newArray, *ab.Array)
		ab.Array = &newArray
	}

	if ab.Length > 0 {

		for i := ab.Length; i > 0; i-- {
			(*ab.Array)[i] = (*ab.Array)[i-1]
			// Se tiene que poner entre parentesis el *ab.Array porque recordemos que estamos usando un pointer, por lo que, si le quitaramos los parentesis estariamos intentando acceder a la direccion en memoria de algo que todavia no sabemos, porque pues esta tomando todo en conjunto. POrque si no le estamos diciendo basicamente eso ve a 0x2309[i] lo cual no se puede
		}
	}
	if ab.Head != 0 {
		ab.Head--
	}
	ab.Length++
	(*ab.Array)[ab.Head] = item
}

func (ab *ArrayBuffer) Pop() {
	if ab.Length > 0 {
		ab.Length--
		(*ab.Array)[ab.Head] = nil
		ab.Head++
		// Recordemos que en el pop pues estamos sacando el primer elemento por lo que nuestra head se tiene que mover a la right
		// Lo que pasaba antes del  18446744073709551615 es porque estamos usando unsigned ints por lo que el truco del math.max nose puede hacer porque 0 -1 en un unsigned int lo que hace es ponerle el numero mas grande porque es como si diera la vuelta qeu ironico verdad en un ring buffer
	}
}

func appendGamer(array []interface{}, secondArray []interface{}) []interface{} {
	for i := 0; i < len(secondArray); i++ {
		array[i] = secondArray[i]
	}
	return array
}

func main() {
	asdf := NewArrayBuffer(3)
	asdf.Push(453)
	asdf.Push(34)
	asdf.Push(34)
	asdf.Push(38)
	asdf.Push(41)
	asdf.Push(44)
	fmt.Println("arreglo asi enterito")
	fmt.Println(*asdf.Array...)
	asdf.Pop()
	fmt.Println("arreglo asi despues popo")
	fmt.Println(*asdf.Array...)
	asdf.Push(120)
	fmt.Println("arreglo asi despues de push")
	fmt.Println(*asdf.Array...)
}
