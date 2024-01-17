package arraybuffer

import (
	"fmt"
	"reflect"
)

type ArrayBuffer struct {
	Array    *[]interface{}
	Length   int
	Capacity int
	Type     reflect.Type
	Head     int
	Tail     int
	IsElden  bool
}

func NewArrayBuffer(c int) ArrayBuffer {
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
	}

	if ab.Type != reflect.TypeOf(item) {
		return
	}

	if ab.Length+1 > ab.Capacity && ab.Tail != ab.Capacity {
		ab.Capacity += 5
		newArray := make([]interface{}, ab.Capacity)
		newArray = appendGamer(newArray, *ab.Array)
		ab.Array = &newArray
		if ab.IsElden {
			fmt.Println(ab.Head)
			fmt.Println(ab.Tail)
		}
	}

	if ab.Tail == ab.Length && ab.Tail != 0 {
		ab.Tail = ab.Tail % ab.Length
		ab.IsElden = true
	} else {
		ab.Tail++
	}

	if ab.Length > 0 {
		(*ab.Array)[ab.Tail] = item
	} else {
		(*ab.Array)[ab.Head] = item
		ab.Tail = ab.Head
	}
	ab.Length++

}

func SettleHeadDown(ab *ArrayBuffer) {
	for i := 0; i < ab.Capacity; i++ {

	}
}

func (ab *ArrayBuffer) Dequeue() interface{} {

	var dequeued interface{}

	if ab.Length > 0 {
		ab.Length--
		dequeued = (*ab.Array)[ab.Head]
		(*ab.Array)[ab.Head] = nil
		if ab.Head+1 != ab.Capacity {
			ab.Head++
		} else {
			ab.Head = ab.Tail - 1
		}

	}

	return dequeued
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
		ab.Capacity += 5
		newArray := make([]interface{}, ab.Capacity)
		newArray = appendGamer(newArray, *ab.Array)
		ab.Array = &newArray
	}

	if ab.Length > 0 && ab.Head == 0 {

		for i := ab.Length; i > 0; i-- {
			(*ab.Array)[i] = (*ab.Array)[i-1]
			// Se tiene que poner entre parentesis el *ab.Array porque recordemos que estamos usando un pointer, por lo que, si le quitaramos los parentesis estariamos intentando acceder a la direccion en memoria de algo que todavia no sabemos, porque pues esta tomando todo en conjunto. POrque si no le estamos diciendo basicamente eso ve a 0x2309[i] lo cual no se puede
		}
	}
	if ab.Head != 0 {
		// Creo que head deberia de equivaler a la Length
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
