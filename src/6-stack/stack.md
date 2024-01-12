# Stack

El stack es lo opuesto a una cola, es decir, el primero en entrar sera el ultimo en salir porque es una estructura de datos que se basa en el concepto de apilar cosas, tomemos de ejemplo apilar platos, si vas a sacar el primer plato que se apilo tendras que quitar todos los que pusiste encima de el antes que nada

- Ultimo elemento agregado
- 
- 
- 
- 
- 
- Primer elemento agregado 

(A) <- (B) <- (C) <- (D) <- (E)   
                            ^
                            head

Agregar un nodo al stack (apilar)

- >>>>
  >>>> -
- >>>> - 
- >>>> -

En el caso del stack nuestra head es el ultimo elemento agregado, por eso se dice que es inverso a una Queue

Por lo que si quisieramos agregar un nuevo elemento a nuestro stack lo que tendriamos que hacer seria lo siguiente

## Agregar un nodo al stack

Hacer que nuestro nuevo nodo apunte a la head actual
----------------------------------------------------

Hacemos que la propiedad prev de nuestro nuevo nodo apunte al head de nuestro stack

node := Node { Value:"F" }
currentHead := stack.head
node.prev = currentHead

(A) <- (B) <- (C) <- (D) <- (E)    (F)
                            ^
                            head

(A) <- (B) <- (C) <- (D) <- (E) <- (F)

----------------------------------------------------


Hacer que la head del stack apunte al nuevo nodo
----------------------------------------------------

Luego lo que hacemos es hacer que la head de nuestro stack apunte al nuevo nodo insertado

stack.head = node

(A) <- (B) <- (C) <- (D) <- (E) <- (F)
                            ^
                            head

(A) <- (B) <- (C) <- (D) <- (E) <- (F)
                                    ^
                                    head

----------------------------------------------------

## Sacar el ultimo elemento agregado del stack

Hacemos que nuestra head apunte al elemento previo de la head actual
--------------------------------------------------------------------

oldHead := stack.head
stack.head = oldHead.prev

Hacemos que nuestro head apunte al elemenot previo

(A) <- (B) <- (C) <- (D) <- (E) <- (F)
                                    ^
                                    head

(A) <- (B) <- (C) <- (D) <- (E) <- (F)
                            ^
                            head

--------------------------------------------------------------------


Hacemos que el elemento que sacamos ya no apunte al nodo previo de nuestro stack
--------------------------------------------------------------------

Hacemos que la head vieja apunte a nil
oldHead.prev = nil

(A) <- (B) <- (C) <- (D) <- (E)    (F)
                            ^
                            head

--------------------------------------------------------------------
