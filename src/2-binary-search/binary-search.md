Binary search es O(N) porque como ignoramos las constantes, pero pues es medio logaritmica

Solo se hace si sabemos que nuestro arreglo está ordenado

m = floor(low + (high − low) / 2)

En la busqueda bianria lo que hacemos es tener como parametro un rango por el cual queramos buscar en nuestro arreglo dado

Por ejemplo

arr := [1,2,3,4,4,4,4,4,5,5,5,5,5,6,7,8,9]

Tenemos nuestro arreglo de 16 elementos

Queremos realizar la busqueda de 6

Asi que le damos de rango para que empiece a buscar desde el indice 5 (low)

Y el limite que sea el indice 15 (high)

Entonces vamos a empezar a buscar haciendo mitades y moviendonos entre ellas, partiendo desde un rango

[1,2,3,4,4,4,4,4,5,5,5, 5, 5, 6, 7, 8, 9]
 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16
           |          |             |
          low         mitad        high
Pero como vamos a recorrer el arreglo siempre tenemos que sumarle nuestro low porque
Por ejemplo en nuestra primera vuelta tenemos que restar 15 - 5 (10) y dividirlo entre 2 (5), es decir 5 indices a la derecha del low esta la mitad, por eso tenemos que sumarle el low porque es nuestro punto de partida (ordenada al origen)

m = floor(low + (high − low) / 2)


