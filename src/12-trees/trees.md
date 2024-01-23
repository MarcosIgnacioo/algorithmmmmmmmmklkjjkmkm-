# Trees

Los arboles son listas ramificadas

Por ejemplo 
         A
        /|\
       / | \
       B C D

El nodo A tiene de hijos a B, C y D

Un ejemplo de un arbol es el propio sistema de archivos de la computadora, cuando ponemos ls nos dicen varias cosas
~ ls

Desktop                    Public                     intelephense
Documents                  adf                        laraveltest
Downloads                  algorithmmmmmmmmklkjjkmkm- notas
Library                    asdf.go                    popo
Movies                     culo.go                    the-end-of-ramgelion
Music                      go                         your-name
Pictures                   hola.go

En este caso ~ es el padre de todos estos nodos (archivos)

En el codigo cada nodo se representa de la siguiente manera

type Node struct {
    Value interface {}
    Children [] *Node
}

Por ejemplo el DOM tambien es un arbol

             HTML
            /    \
           /      \
          Head     Body
         /        / |  \
        links   div div div

root - Top most node (Por ejemplo en los casos de arriba serian HTML en el caso del DOM y ~ en el del sistema de archivos)

height - La distancia entre el nodo root y el nodo hijo mas lejano

binary tree - Arbol con maximo 2 hijos y la menor cantidad de hijos permitidos es 0

general tree - Arbol de 0 a N cantidad de hijos

binary search tree - Un arbol con un orden de los nodos especifico

leaves - Un nodo sin hijos

balanced - Se dice que un nodo esta perfectamente balanceado cuando esta asi

        5
       / \
      3   8
     / \ / \
    1  4 7  9

Es decir cuando tenemos la misma distancia entre todos

branching factor - cantidad de hijos que tiene el arbol

Traversals de trees

Pre order

        5
       / \
      3   8
     / \ / \
    1  4 7  9

Traversar es simple

    Lo que vamos a hacer es algo llamado visitar a un nodo
    
    Lo que significa que haremos algo con el valor del nodo
    
    Y luego vamos a recursar

Vamos a tener un pre y un post recurse step

Si quisieramos imprimir el valor del nodo que estemos visitando actualmente lo que hariamos seria lo siguiente

visitNode()
print(node.value)
recurse()

        5
       / \
      3   8
     / \ / \
    1  4 7  9

        +----- Empezamos en root, imprimimos el valor de root, vemos si     |      tiene un nodo a la izquierda. Si existe y aun no lo hemos    v      visitado avanzamos a ese
        5
       / \
      3   8
     / \ / \
    1  4 7  9

output : 5

 Estamos en el nodo que tiene de valor 3, lo imprimimos, vemos si tiene un nodo a la izquierda.vemos si tiene un nodo a la izquierda. Si existe y aun no lo hemos visitado avanzamos a ese
 ^
 |      5
 |     / \
 +--> 3   8
     / \ / \
    1  4 7  9

output : 5,3

 Estamos en el nodo que tiene de valor 1, lo imprimimos, vemos si tiene un nodo a la izquierda.vemos si tiene un nodo a la izquierda. Si existe y aun no lo hemos visitado avanzamos a ese, pero este ya no tiene nodos a la izquierda, ni a la derecha, es una hoja por lo que tenemos que retroceder al padre de este nodo, el nodo anterior el que tenia el valor de 3

^       5
|      / \
|     3   8
|    / \ / \
+-> 1  4 7  9
output : 5,3,1

 Estamos en el nodo que tiene de valor 3, pero como ya habiamos estado aqui significa que ahora tenemos que ir al nodo a la derecha
 ^
 |      5
 |     / \
 +--> 3   8
     / \ / \
    1  4 7  9

        5
       / \
      3   8
     / \ / \
    1  4 7  9
       ^
       Ahora que estamos en el nodo con el valor 4 lo que hacemos es imprimirlo, y vemos si tiene un nodo a la izquierda, no tiene, vemos si tiene un nodo a la derecha, tampoco tiene, por lo que regresamos al nodo padre del 4 que es el anterior (el nodo con valor de  3)

output : 5,3,1,4

 como ya visitamos los nodos hijos de este nodo retrocedemos otra vez al nodo padre del nodo actual
 ^
 |      5
 |     / \
 +--> 3   8
     / \ / \
    1  4 7  9

        v------------- Como ya visitamos el nodo a la izquierda de nuestra  5              raiz ahora visitaremos el de la derecha
       / \
      3   8
     / \ / \
    1  4 7  9

        5
       / \
      3   8 <------ Imprimimos el valor del nodo actual vemos si tiene un nodo a la izquierda como si lo tiene avanzamos a ese
     / \ / \
    1  4 7  9

output : 5,3,1,4,8

        5
       / \
      3   8
     / \ / \
    1  4 7  9
         ^
         Imprimimos el valor actual como no tiene ningun hijo retrocedemos

output : 5,3,1,4,8,7

        5
       / \
      3   8 <------ Como ya habiamos visitado la izquierda ahora visistamos a la derecha
     / \ / \
    1  4 7  9
        5
       / \
      3   8 
     / \ / \
    1  4 7  9 <----- Imprimimos el valor actual y como es una hoja nos tenemos que regresar

output : 5,3,1,4,8,9

        5 <---- Ya hemos visitado ambos hijos y es la raiz asi que el recorrido se da como terminado
       / \
      3   8 <---- Ya hemos visitado ambos hijos, retrocedemos
     / \ / \
    1  4 7  9

output : 5,3,1,4,8,7,9

Aqui lo que estamos haciendo es 

Imprimir el valor del nodo
Visitar el siguiente nodo con recursividad

In order

        5
       / \
      3   8
     / \ / \
    1  4 7  9

output : 1, 3, 4, 5, 7, 8, 9

Visitamos hasta el ultimo nodo, luego volvemos al padre lo visitamos tambien, vamos a la derecha, la visitamos tambien y luego regresamos otra vez al padre, regresamos al padre de ese padre y lo visitamos, luego vamos a la derecha ...

Post order

        5
       / \
      3   8
     / \ / \
    1  4 7  9

output : 1,4,3,8,7,9,5

Aqui visitamos a los hijos primero izquierdo y derecho y luego visitamos al padre de manera que siempre el padre es el ultimo en aparecer en el output


Pre order
output : [5],3,1,4,8,7,9
//////////////////////
In order
output : 1,3,4,[5],7,8,9
//////////////////////
Post order
output : 1,4,3,8,7,9,[5]
//////////////////////

En el preorden nuiestra root queda al principio,
En el inorder queda en el medio
En el post orden queda al final
Ver el final de implement tree traversal sobre como funcionan el stack de lsa funciones creo que puede servir bastante a como vemos la recursividad

Tree level traversal (Breadth tree traversal) Level order traversal el primer termino no es "correcto" pero es el mas entendible


            5
    ------------------ Primero
           / \
          3   8
    ------------------ Segundo
         / \ / \
        1  4 7  9
    ------------------ Tercero

En un tree level traversal lo que hacemos es visitar por niveles, es decir, a nuestro nivel de root pues solamente tendremos la root por lo que es el unico nodo que tenemos que visitar, luego estan el 3 y el 8, entonces visitamos el 3, luego el 8, luego en el siguiente nivel visitamos el 1, 4, 7 y 9

Para que al final nuestro traversal quede como

5,3,8,1,4,7,9

Pero lo que pasa debajo del todo es que estamos usando una cola, de la siguiente manera


            5
           / \
          3   8
         / \ / \
        1  4 7  9

Vamos a agregar 5, 5 tiene 2 hijos asi que los ponemos tambien
Queue:

ENQUEUE +
------------------
[]<---+
      |
^     |
head  tail

[5]
^      |
head   tail

-----------------

ENQUEUE +
------------------

Queue:

[5][3][8]<---+
^            |
head         tail
En este punto 5 ya no tiene mas hijos asi que lo imprimos y lo sacamos de la cola

------------------

DEQUEUE -
------------------

Queue:

[5][3][8]<---+
^            |
head         tail

[3][8]<------+
 ^           |
 head        tail
Antes de sacarlo lo imprimimos
output: 5

------------------


Ahora, 3 tiene hijos asi que los metemos a la cola

ENQUEUE +
------------------

[3][8][1][4]<+
 ^           |
 head        tail

------------------

Ahora 3 ya no tiene hijos que queramos agregar a la cola por lo que lo sacamos de ella preo antes lo imprimimos 

DEQUEUE -
-------------------

output: 5,3

[3][8][1][4]<+
 ^           |
 head        tail

[8][1][4]<---+
 ^           |
 head        tail

-------------------

ENQUEUE +
-------------------

8 tiene 2  hijos asi que los metemos a la cola

[8][1][4][7][9]
 ^           ^
 head        tail

-------------------

Como 8 ya no tiene mas hijos lo sacamos de la cola y lo imprimimos

output: 5,3,8

DEQUEUE -
-------------------

[8][1][4][7][9]
 ^           ^
 head        tail

[1][4][7][9]
 ^        ^
 head     tail

-------------------

Como 1 ya no tiene mas hijos lo sacamos de la cola y lo imprimimos

output: 5,3,8,1

DEQUEUE -
-------------------

[1][4][7][9]
 ^        ^
 head     tail

[4][7][9]
 ^     ^
 head  tail

-------------------

Como 4 ya no tiene mas hijos lo sacamos de la cola y lo imprimimos

output: 5,3,8,1,4

DEQUEUE -
-------------------

[4][7][9]
 ^     ^
 head  tail

[7][9]<-+
 ^      | 
 head  tail

-------------------

Como 7 ya no tiene mas hijos lo sacamos de la cola y lo imprimimos

output: 5,3,8,1,4,7

DEQUEUE -
-------------------

[7][9]<-+
 ^      | 
 head  tail

[9]<----+
 ^      | 
 head  tail

-------------------

Como 9 ya no tiene mas hijos lo sacamos de la cola y lo imprimimos

output: 5,3,8,1,4,7,9

DEQUEUE -
-------------------

[9]<----+
 ^      | 
 head  tail

[]<----+
 ^      | 
 head  tail

-------------------

Y asi es como jala nya

Si usaramos los arrays de js (arraylists) para hacer esto esta operacion nos costaria

ON^2 porque para hacer el reacomodo de las cosas es 
