# Heap

La manera mas facil de definir a la heap, tambien conocida como priority queue es un arbol binario en el cual el nodo padre es mayor a sus hijos (MaxHeap) y cuando es al reves, es decir, el nodo padre es menor a sus hijos se le llama MinHeap

Generalmente nunca recorremos la heap pero si lo hicieramos con Breadth First seria la mejor manera de hacerlo

Ejemplo de Min Heap

               50 <------- The HeapCondition*1
               /\
              /  \
             /    \
            /      \
           85      130
          /  \    /   \
         102 140 170  209


*1 The heap condition : Es la regla con la que en base a ella haremos nuestra heap

    MinHeap: Todos los nodos hijos debajo del padre son MENORES ó igual a él
                                                           <=
        The minimum item is the root

    MaxHeap: Todos los nodos hijos debajo del padre son MAYORES ó igual a él
                                                           >=
        The maximum item is the root

En este caso la heap condition es MinHeap, todos los nodos hijos son mayores al padre

Hay un truco para obtener la mediana de todo esto pero ocupariamos 2 heaps (es una pregunta que hacen en google a los gamers)

## Agregar nodos a nuestra heap

Una propiedad de las heaps es que siempre estan agregando de izquierda a derecha, por lo que NUNCA habria algo asi

               50 <------- The HeapCondition*1
               /\
              /  \
             /    \
            /      \
           85      130
          /  \    /   \
         220 nil 170  209

Si hay un espacio disponible vamos a poner el nodo ahi, no vamos a saltarnos lugares
Y el orden por el que iremos siempre sera de IZQUERDA a DERECHA

### Como agregamos nuestro nodo a un heap?

Vamos al ultimo nodo de nuestro heap

               50 
               /\
              /  \
             /    \
            /      \
           85      130
          /  \    /   \
     --> 220 171 170  209

Y agregamos el nodo, por ejemplo si queremos agregar un nodo que tenga el valor de 3

               50 
               /\
              /  \
             /    \
            /      \
           85      130
          /  \    /   \
        220 171  170  209
        /
        2

Pero como tenemos la regla de MinHeap, que todos los nodos hijos seran mayores a sus padres esto no es posible

Asi que en cada insercion vamos a pasar por el sgiuiente filtro

El nodo que insertamos es mayor o igual (caso MinHeap) que el nodo padre?

Si la respuesta es no, lo que tenemos que hacer es intercambiarlos

               50 
               /\
              /  \
             /    \
            /      \
           85      130
          /  \    /   \
         2  171  170  209
        /
        220

Ahora seguimos haciendo la misma pregunta, el nodo que acabamos de insertar es mayor o igual a su padre (ahora su padre es 85), la respuesta siendo no hacemos el intercambio de nuevo

               50 
               /\
              /  \
             /    \
            /      \
           2       130
          /  \    /   \
         85 171  170  209
        /
        220

De nuevo hacemos la pregunta, el nodo que acabamos de insertar es mayor o igual a su padre actual (50), la respuesta es no por lo que hacemos el intercambio de nuevo

                2
               / \
              /   \
             /     \
            /       \
           50      130
          /  \    /   \
         85 171  170  209
        /
        220

De nuevo hacemos la pregunta, el nodo que acabamos de insertar es mayor o igual a su padre actual (nil porque somos root), como ya no hay padre con el que comparar pues paramos de hacer intercambios, y nuestro heap ya se puede considerar "ordenado", no es un ordenamientro muy fuerte a comparacion de los arboles anteriores pero pues igualmente esta medio ordenado

### Borrar un nodo


                2 <--- Supongamos que queremos borrar el 2
               / \
              /   \
             /     \
            /       \
           50      130
          /  \    /   \
         85 171  170  209
        /
        220

               [?]
               / \
              /   \
             /     \
            /       \
           50      130
          /  \    /   \
         85 171  170  209
        /
        220
Entonces ahora nuestro nodo que borramos se ha quedado sin valor, por lo que el siguiente paso es agarrar al nodo mas abajo de nuestro heap y ponerlo en ese lugar


              [220]
               / \
              /   \
             /     \
            /       \
           50      130
          /  \    /   \
         85 171  170  209

Y ahora hacemos un bubble down, lo que haciamos en insercion pero para abajo y con un paso extra

Como lo que queremos hacer es que nuestros nodos hijos sean mayores que los padres

    MinHeap: Todos los nodos hijos debajo del padre son MENORES ó igual a él
                                                           <=
        The minimum item is the root
Porque es MinHeap, lo que tenemos que hacer es elegir a nuestro hijo con valor mas pequeño y compararlo con el nodo que andamos moviendo (220). En este caso el que es mas pequeño entre los hijos es el 50, bien ahora hacemos la siguiente comparacion

if 50 < 220 {
    bubbleDown(220)
} else {
    endBubbleDown()
}
Como 50 es menor que 220 hacemos bubble down 

                50
               /  \
              /    \
             /      \
            /        \
          [220]      130
          /  \      /   \
         85 171    170  209

Lo hacemos de nuevo, elegimos a nuestro nodo hijo mas pequeño y lo comparamos con nuestro nodo que andamos bubbleando

if 85 < 220 {
    bubbleDown(220)
} else {
    endBubbleDown()
}

                50
               /  \
              /    \
             /      \
            /        \
            85       130
          /    \    /   \
         [220] 171 170  209

Como ya no tiene hijos con los que hacer las comparaciones pues no tenemos necesidad damos por terminado nuestro proceso de bubble down

Ahora, como podemos obtener el ultimo elemento de nuestra queue, muy facil, con indices

Asignaremos a cada nodo un indice como si fuera un array de la siguiente manera

                i=0
                50
               /  \
              /    \
             /      \
            /        \
      i=1 ->85       130 <- i=2
          /    \    /   \
         220  171 170  209
         ^    ^   ^    ^
         i=3  i=4 i=5  i=6

   [50,85,130,220,171,170,209]
i=  0  1  2   3   4   5   6

Okay pero como sabemos cuales son los hijos

MATEMATICAS

Hay una formula para saber el indice que ocupan los hijos de un nodo

(2*i)+1 --- Hijo izquierdo

(2*i)+2 --- Hijo derecho

POngamoslo a prueba

El nodo 130 tiene el indice 2 y sus hijos son 170 (i = 5) y 209 (i = 6)

Hagamos la prueba

(2*i)+1 --- Hijo izquierdo 

(2*2)+1 --- (170)
4+1
5

SIii el indice de 170 es 5

(2*i)+2 --- Hijo derecho

(2*2)+2 --- 209
4+2
6

SIIII el indice de 209 es 6

Por lo que ya sabemos como podemos obtener los hijos de un nodo

PEro como obtenemos el padre de un nodo?

muy faicl, revertimos esta formula

(i-1)/2

Probemoslo

209 (i = 6)

(6-1)/2
5/2
2
(ambos son integers por lo que la division no dara un float, aunque en Js y python si que da un float pero GOD da un integer)

y booom se puede ver que si que funciona, realmente no tenemos que hacer lo de restar 2 cuando es el derecho y 1 cuando es el izquierdo porque son integers por lo que siempre el numero antes del punto siempre sera el index correcto por lo que nos ahorramos de poner unos ifs

Y la manera de trackear a nuestro ultimo nodo es simplemente con la length de nuestro heap - 1

Boom



