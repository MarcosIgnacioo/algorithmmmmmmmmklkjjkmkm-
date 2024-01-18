# Quick sort

Quick sort

Funciona de una manera parecida al binary search

Tenemos un arreglo 

[2,7,8,3,5,1,9,4]

Escogemos nuestro pivot, en este caso elegiremos 5

Entonces ahora tenemos que leer todo nuestro array de la siguiente manera

   [2,7,8,3,5,1,9,4]
 n: 0,1,2,3,4,5,6,7


////////////////////////////////////////////////////////////////////////////

n = 0

[2,7,8,3,5,1,9,4]
 ^
2 <= 5

En otro array lo que hacemos es poner el 2 a la izquierda del 5 y pues llevamos un contador de las posiciones que van agarrando 

low = 0
high = len(array)-1
array[low] = array[n]
n++
low++

[2,_,_,_,5,_,_,_]

////////////////////////////////////////////////////////////////////////////

n = 1

[2,7,8,3,5,1,9,4]
   ^
7 <= 5

Como el 7 es mayor que el 5 lo que tenemos que hacer es ponerlo hasta la derecha en el high y actualizar nuestro high


array[]
array[high] = array[n]
n++
high--

[2,_,_,_,5,_,_,7]


////////////////////////////////////////////////////////////////////////////

n = 2

[2,7,8,3,5,1,9,4]
     ^
8 <= 5

Como el 8 es mayor que el 5 lo que tenemos que hacer es ponerlo hasta la derecha en el high y actualizar nuestro high


array[]
array[high] = array[n]
n++
high--

[2,_,_,_,5,_,8,7]

////////////////////////////////////////////////////////////////////////////

n = 3

[2,7,8,3,5,1,9,4]
       ^
3 <= 5

Como el 3 es meno que el 5 lo que tenemos que hacer es ponerlo hasta la izquierda en la posicion de low 


array[]
array[low] = array[n]
n++
high++

[2,3,_,_,5,_,8,7]

////////////////////////////////////////////////////////////////////////////

n = 4

[2,7,8,3,5,1,9,4]
         ^
5 <= 5

skipeamos el pivot

////////////////////////////////////////////////////////////////////////////

n = 5

[2,7,8,3,5,1,9,4]
           ^
1 <= 5

Como el 8 es mayor que el 5 lo que tenemos que hacer es ponerlo hasta la derecha en el high y actualizar nuestro high


array[]
array[low] = array[n]
n++
low++

[2,3,1,_,5,_,8,7]

////////////////////////////////////////////////////////////////////////////

n = 6

[2,7,8,3,5,1,9,4]
             ^
9 <= 5

Como el 9 es mayor que el 5 lo que tenemos que hacer es ponerlo hasta la derecha en el high y actualizar nuestro high


array[]
array[high] = array[n]
n++
high--

[2,3,1,_,5,9,8,7]

////////////////////////////////////////////////////////////////////////////

n = 7

[2,7,8,3,5,1,9,4]
               ^
4 <= 5

Como el 4 es menor que el 5 lo que tenemos que hacer es ponerlo en el low 

array[]
array[low] = array[n]
n++
low++

[2,3,1,4,5,9,8,7]

Aqui tenemos un weak sort

Luego hacemos mas divisiones

[2,3,1,4,5,9,8,7]
     ^   ^   ^
     p2  p1  p3

Empecemos con el pivot 2

este se va acomparar a 2 y como 2 es mayor al pivot ira a la derecha

1,2

Luego el pivot se va a comparar con el 3 y como el 3 es mayor que el 1 se va a poner al a derecha 
1,2,3


Ahora con el pivot 3

7 es menor a 8 asi que se pone al a izquierda 

7,8

skipeamos el pivot

y luego el 9 es mayor que el 8 asi que se pone al a derecha

7,8,9

Juntamos todo y queda ya 

1,2,3,4,5,7,8,9

En el peor de los escenarios que pueden haber es que el pivot que escojamos sea 1 o el elemenmto mas grande del array apor lo que siempre escogemos el elemento del medio puesto que en el caso de que nos pasen un array invertido podemos tener biseccion perfecta y si nos pasan un array al azar pues simplemente tenemos la propbabilidad de que sea un elemento middlelish

en el caso de que ele emento sea uno o el elemento mas pequeño posible pues va a ser o(n^2)


