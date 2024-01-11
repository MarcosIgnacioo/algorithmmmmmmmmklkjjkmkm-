# Bubble Sort

Bubble sort funciona de la siguiente manera

Cada elemento del arreglo que no sea el ultimo se va a comparar con el elemento que tengan al lado, si el elemento i es mayor que el elemento i+1, intercambiarán lugares, si no, se pasa al siguiente elemento

arreglo = [1,3,7,2,4,6,5]

[1,3,7,2,4,6,5]
 ^
i     = 0
i + 1 = 1

1 > 3 NO, todo sigue igual

[1,3,7,2,4,6,5]
   ^

i     = 1
i + 1 = 2

3 > 7 NO, todo sigue igual

[1,3,7,2,4,6,5]
     ^

i     = 2
i + 1 = 3

7 > 2 SI, cambiamos de lugar al 7 y el 2


[1,3,2,7,4,6,5]
       ^

i     = 3
i + 1 = 4

7 > 4 SI, cambiamos de lugar al 7 y el 4

[1,3,2,4,7,6,5]
         ^

i     = 4
i + 1 = 5

7 > 6 SI, cambiamos de lugar al 7 y el 6

[1,3,2,4,6,7,5]
           ^

i     = 5
i + 1 = 6

7 > 5 SI, cambiamos de lugar al 7 y el 5


Asi es como queda despues de la primera iteracion completa del bubble sort, como se puede ver aun faltan cosas por acomodarse pero el elemento mas grande esta hasta el final, lo cual siempre pasa en la primera vez se recorre el arreglo para ordenarlo con bubble sort, asi que la siguiente vez que tengamos que recorrer el arreglo no tendremos que recorrer todo completo, lo iremos ordenando poco a poco y cada vez sera lo menos que se tendrá que recorrer

[1,3,2,4,6,5,7]

[1,3,2,4,6,5 |,7]

[1,2,3,4,5 |,6,7]


Como podemos saber la O del bubble sort?

De la misma manera que gauss descubrió la manera de sumar 1,2,3,4...100 con una multiplicacion facil

100 + 1 = 101
99  + 2 = 101
98  + 3 = 101
.
.
.
50 + 51 = 101

Entonces si sumamos los numeros extremos siempre da 101 

y para saber la suma de todos es tan simple como multiplicar 101 * 50

Y pasa de la misma manera con el bubble sort


[1,3,2,4,6,5,7]   N = 3

[1,3,2,4,6,5 |,7] N = 2

[1,2,3,4,5 |,6,7] N = 1

si esto fuera un arreglo mas largo la comparacion con gauss seria mas epica pero pues ni pedo

Pero pues basicamente es la suma de todas las N, que en realidad se puede calcular con la formula de gauss
N + 1 * (N/2)

 (N (N + 1))
-------------
     2

 (N^2 + N)
-----------
     2


Eliminamos el /2 porque no nos importan las constantes

 (N^2 + N)

Y en big o tambien quitamos valores sin signficado, porque por ejemplo en el caso de que N fuese 10000 el + N se convierte irrelevante

N^2


