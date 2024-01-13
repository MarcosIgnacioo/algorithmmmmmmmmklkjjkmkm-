# Array buffers o Ring buffers

* Se le dice shifting u unshifting a la accion de enqueue y dequeue en js porque estan moviendo el arreglo a la derecha (dequeue) para poder quitar el primer elemento (fifo) y unshifting (enqueue) para poder agregar un elemento al final 

Probablemente la estructura de datos mas sexy de implementar porque ohhh booy

Son como una arraylist pero en vez de usar el indice 0 como la head y el indice length como la tail si usa punteros llamados head y tail, 
Lo chido de esto es que podemos tener algo asi

[1,2,3,4,5,6,7,8,9]
   ^         ^
   head      tail
 
# # # SI QUIERES USAR EL ARRAYBUFFER COMO UNA QUEUE

# Si quieres hacer un dequeue

[1,2,3,4,5,6,7,8,9]
     ^     ^
     head  tail

Mueve el head 1 indice a la derecha

[1,2,3,4,5,6,7,8,9]
       ^   ^
      head tail

# Si quieres hacer un enqueue

1.- Mueves la cola por 1 indice a la derecha
[1,2,3,4,5,6,7,8,9]
     ^       ^
     head    tail

2.- Insertas el elemento que quieres que se una a la cola
[1,2,3,4,5,6,9,8,9]
     ^       ^
     head    tail

# # # SI QUEREMOS USAR NUESTRO ARRAY BUFFER COMO STACK (REALMENTE AUN TENDRIA UNA TAIL PORQUE PUES ES UN ARRAY BUFFER PERO SE LA QUITE PARA QUE TENGAMOS MAS CLARIDAD)

# Si quieres hacer un push (Insertar a la cabeza )

[1,2,3,4,5,6,9,8,9]
             ^
             head

1.- Mueve la head 1 elemento a la derecha

[1,2,3,4,5,6,9,8,9]
               ^
               head

2.- Inserta el valor que quieras

[1,2,3,4,5,6,9,89,9]
               ^
               head


# Si quieres hacer un pop (Sacar el ultimo elemento insertado (En un stack el ultimo elemento insertado es la head))


[1,2,3,4,5,6,9,89,9]
               ^
               head

1.- Mueve la head 1 elemento a la izquierda

[1,2,3,4,5,6,9,89,9]
             ^
             head


NO uses un arraybuffer como cola y stack al mismo tiempo porque eso seria como querer bañarte mientras te lavas los dientes, probablemente lo puedas hacer pero solo un psicopata haria eso

Pero que pasa si excedemos nuestro limite?

[1,2,3,4,5,6,9,8,9]
       ^         ^
       head      tail
Tenemos un arreglo de 9 elementos, supongamos que queremos hacer que nuestra tail sea 10, podemos hacerlo sin resizear el arreglo, con una simple operacion

tail % len 

10 % 9

1

Y le tenemos que restar 1 porque estamos usando len para validar esto y pues indices empiezan de 0

1-1

0

Entonces tenemos que mover nuestra cola al indice 0

[1,2,3,4,5,6,9,8,9]
 ^     ^         
 tail  head      


Por que usariamos un arraybuffer en vez de una array list?

Bueno hay situaciones en las que en nuestra arraylist estamos desalojando espacio en memoria que podriamos usar en vez de aumentar la capacidad de nuestra lista

por ejemplo

# Quitar primer elemento

Tenemos la siguiente lista, pero queremos quitar el primer elemento
[1,2,3,4,5,nil,nil,nil,nil]
 ^       ^         
 head    tail 

Pues simplemente movemos nuestra head a la derecha y ya

[1,2,3,4,5,nil,nil,nil,nil]
   ^     ^         
   head  tail

# Agregar un elmento al head
Pero que pasa si agregamos otro elemento a nuestra head

Pues movemos la head a la izquierda y le ponemos el elemento que queriamos insertar

[nil,nil,11,2,3,4,5]
         ^        ^         
         head     tail

# Agregar otro elemento a la tail

Normalmente en otra estructura de datos, como el stack, lo que tendriamos que hacer en esta situacion seria mover nuestro arreglo a la izquierda todos un lugar, sin embargo hacer esto nos cuesta O N operacines lo cual no podemos permitir!

Sobretodo teniendo en cuenta que realmaente nuestra lista todavia tiene la capacidad para almacenar mas elementos, ssolamente que estos espacios disponibles se encuentran al final de la lista, como podemos solucionar esto???

Array buffer babyyy

Primero calculamos la posicion gamer buffer (i = 6)

tail % len 

6 % 5

1

1 - 1

0

ya tenemos el indice en el que se puede agregar la tail, pero ahora debemos de checar que el indice ese sea menor al de nuestra head

Nuestra head se encuentra en el indice 2, 0 es menor que 2 asi que si podemos hacer el recorrido circular magico

[nil,nil,11,2,3,4,5]
 ^       ^                 
 tail    head

Cosas importantes: No se si pueda hacerse lo mismo con el head es dsecir que con el head estirarlo como hicimos con la cola pero pues mañana vemos eso 
