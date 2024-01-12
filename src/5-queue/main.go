package main

import "fmt"

func main()  {
    var arreglo [3]int  
    fmt.Println(len(arreglo))
}

type Node struct {
    Value interface{}
    Next *Node
}

type Queue struct {
    Length uint
    Head *Node
    Tail *Node
}

// Head y Tail son directorios que apuntan a nuestro primer (head) elemento y nuestro ultimo (Tail) elemento

func NewQueue() *Queue {
   return &Queue { Length: 0, Head: nil, Tail: nil }
}

// Enqueue : Agregar un elemento a nuestra queue (como es una queue vamos a agregarlo hasta el final de ella)

// Tomamos el largo de nuestra cola y lo aumentamos en 1

// Creamos un nuevo nodo con el valor del item que nos pasaron por el constructor

// Si la tail de nuestra queue es nil, signfica que nuestra Queue esta vacia por lo que haremos que la head y la tail sean el mismo nodo porque pues ahora es de 1 elemento

// Si la tail no es nil, es decir ya hay mas Nodos dentro de nuestra queue, lo que haremos sera hacer que el ultimo nodo actual (nuestra cola antes de la insercion) su next apunte al nuevo nodo, y luego hacemos que la tail de la queue sea el nuevo nodo

    // (E) Es nuestro nuevo nodo que queremos insertar

    // (A) -> (B) -> (C) Es nuestra cola a la que queremos insertar nuestro nuevo nodo

// Primero checamos que nuestra queue no tenga una cola en nil (Es decir que nuestra queue este vacia)

    // nil
    // ^tail

    // Hacemos que nuestro nodo nuevo sea tanto la cola como la cabeza porque pues ahora es una queue de 1 elemento
    // (E)
    // ^head
    // ^tail

// En caso de que nuestra queue tenga mas elementos hacemos lo siguiente

    // (A) -> (B) -> (C) 
    // ^head         ^tail

// Hacemos que nuestra tail actual (C) su next apunte a nuestro nuevo nodo
    // (A) -> (B) -> (C) -> (E)
    // ^head         ^tail


// Hacemos que nuestra tail apunte a nuestro nuevo nodo
    // (A) -> (B) -> (C) -> (E)
    // ^head                ^tail

func (q *Queue) Enqueue(item interface{}) {
    q.Length++
    n := Node { Value: item }

    if q.Tail == nil {
        q.Head = &n
        q.Tail = &n
    }

    oldTail := q.Tail
    // Guardamos la tail actual de nuestra cola (C)

    oldTail.Next = &n
    // Hacemos que nuestra tail actual (C) su next apunte a nuestro nuevo nodo
    // (A) -> (B) -> (C) -> (E)
    // ^head         ^tail

    q.Tail = &n
    // Hacemos que nuestra tail apunte a nuestro nuevo nodo
    // (A) -> (B) -> (C) -> (E)
    // ^head                ^tail

}

// Dequeue : Sacar a nuestro primer elemento de la queue (First In First Out)

// Si nuestra head es nil signfica que nuestra queue esta vacia asi que non podemos sacar nada

// De caso contrario lo que hacemos es disminuir la Length de nuestra queue y hacemos que la head de nuetra queue ahora sea el segundo elemento de la queue, y nuestra head vieja su next sea nil para desconectarla por completo de la queue

// Cosas que hicimos

    // Hacer que la head apunte al segundo elemento de la lista
    // Hacer que nuestra head antigua (la que vamos a sacar) ahora ya no apunte a nada en su .Next para completar la desvinculacion

    // (A) -> (B) -> (C)
    // ^head          ^tail
// Por ejemplo si le hicieramos un Dequeue() a esta Queue lo que pasa es lo siguiente

// Primero hacemos que nuestra head apunte al elemento seguido (el segundo) del que vamos a sacar (el primero)
    // (A) -> (B) -> (C)
    //        ^head  ^tail

// Ahora lo que hacemos es desvincular al elemento de la Queue haciendo que su next pase de apuntar al nodo B a que apunte a un nil

    // (A)    (B) -> (C)
    //        ^head  ^tail

func (q *Queue) Dequeue() interface{}  {
    if q.Head == nil {
        return nil
    }
    q.Length--

    oldHead := q.Head
    // Guardamos al primer elemento (Al que vamos a sacar)
    // (A) -> (B) -> (C)
    // ^head         ^tail

    q.Head = q.Head.Next

    // Hacemos que la head apunte al segundo elemento de nuestra queue
    // (A)    (B) -> (C)
    //        ^head  ^tail

    oldHead.Next = nil
    // Desvinculamos a la cabeza antigua de la queue
    // (A)    (B) -> (C)
    //        ^head  ^tail

    return oldHead.Value
}

    // Basicamente lo que estamos haciendo es cambiar a donde apuntan los next y a donde apunta nuestras head o tail

// Devulve el valor del primer elemento de nuestra Queue (Nuestra head)
func (q *Queue) Peek() interface{}  {
    return q.Head.Value
}
