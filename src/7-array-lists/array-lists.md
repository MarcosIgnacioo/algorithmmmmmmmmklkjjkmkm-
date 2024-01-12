Que es lo que pasa cuando quieres lo mejor de ambos mundos?

Nacen las arraylists

Las arraylist son basicamente arrays puros y cuando es necesario agregar elementos y la cantidad a agregar excede el tamaño de nuestro arreglo, creamos otro arreglo con mayor espacio para nuevos elementos

Las arraylists parten del concepto de dos propiedades fundamentales para poder operar

    length 
    capacity

Length almacenaría la cantidad de elementos que tiene nuestra arraylist actualmente, mientras que capacity la cantidad de elementos que soporta actualmente nuestro arreglo.

Si por ejemplo tenemos un arreglo asi

[1,_,_]

Length = 1
Capacity = 3

Actualmente nuestro arreglo cuenta con 1 elemento y tiene la capacidad de soportar hasta 3 elementos en total

Si por ejemplo quisieramos hacerle un push de un nuevo elemento (un nuevo numero)

lo que tendriamos que hacer en nuestra funcion de push seria
    if Length + 1 <= Capacity {
        ////////////////////////
        //  [1 , _ , _]       //
        //   ^       ^        //
        //   length  capacity //
        //   = 1     = 3      //
        ////////////////////////
        array[Length] = nuevoElemento
        Length++
        ////////////////////////
        //  [1 , 2 , _]       //
        //       ^   ^        //
        //   length  capacity //
        //   = 2     = 3      //
        ////////////////////////
        return array
    } else {
        /////////////////////////
        //  [1 , 2 , 3]        //
        //           ^         //
        //           capacity  //
        //           && length //
        //           = 3       //
        /////////////////////////
        var newArray [Capacity * 2] interface {}
        ////////////////////////////////////
        //  [_ , _ , _ , _ , _ , _]       //
        //   ^                   ^        //
        //   length              capacity //
        //   = 0                 = 6      //
        ////////////////////////////////////
        newArray = array
        ////////////////////////////////////
        //  [1 , 2 , 3 , _ , _ , _]       //
        //           ^           ^        //
        //           length      capacity //
        //           = 3         = 6      //
        ////////////////////////////////////
        newArray[Length] = nuevoElemento
        Length++
        ////////////////////////////////////
        //  [1 , 2 , 3 , 4 , _ , _]       //
        //               ^       ^        //
        //               length  capacity //
        //               = 4     = 6      //
        ////////////////////////////////////
        //  [1 , 2 , 3 , 4 , _ , _]       //
        ////////////////////////////////////
        return newArray
    }

Tenemos que checar que la cantidad de elementos que tenemos más uno (porque es el que estariamos agregando) en nuestro arraylist no sea superior a la de nuestra capacidad, en caso de que no sea superior, podemos simplente asignar el nuevo elemento en la posicion Length del arreglo (Recordemos que como funcionan los indices en los arreglos, en este caso la Length esta en 1 porque nuestro arreglo tiene un elemento, y este ocupa el indice 0, asi que siempre que intentemos insertar al final del arreglo lo haremos en la posicion de Length)
    
    Si el arreglo tuviese 2 elementos y quisieramos insertar otro, tambien funcionaria lo de ponerlo en la posicion de Length porque, de nuevo, el indice 2 es el del tercer elemento, asi que es una regla que siempre funcionaria pero pues para los casos que queramos hacer insert al final de nuestro arreglo

En caso de que nuestra capacidad sea menor a la cantidad de elementos que llegariamos a almacenar si es que agregamos el elemento que queremos lo que hacemos es crear un nuevo arreglo cn el doble de capacidad de tipo generico, guardamos el arreglo original, y ponemos en la poscion de Length el nuevo elemento de uestro arreglo y retornamos el nuevo arreglo

Eso en el caso de que queramos hacer un push, si queremos hacer un pop

En ese caso simplente hacemos 

if Length - 1 > 0 {
   Length -- 
}

No es necesario poner el ultimo valor en el valor 0 del tipo que sea nuestra arraylist pero podriamos hacerlo, sin embargo con solamente restar nuestra length es mas que suficiente porque recordemos que es con la length que podemos hacer la insercion de nuevos elementos al arreglo, entonces si por ejemplo quisieramos quitar un elementos del siguiente arreglo

[4,3,9]
     ^
     Length = 3
Si le hacemos pop entonces lo que hariamos seria sacar el 3 reduciendo la length en 1

old := array[Length]
Length--
[4,3,9]
   ^
   Length = 2

Asi que el siguiente elemento que queramos insertar a nuestra arraylist va a sobreescribir lo que teniamos antes en el tercer lugar de nuestro arraylist
y retornamos el elemento que hemos quitado

return old


arreglo.push(8)

Length en este caso vale 2 porque le quitamos el tercer elemento
arreglo[Length] = 8

Ahora que pasa si queremos hacer las funciones especiales de las queues

Enqueue no es tan dificil simplemente tenemos que agregar a nuestro elemento al final de nuestro arreglo, es con dequeue que  todo se complica

Recordemos que lo que hace dequeue es sacar al primer elemento de nuestro arreglo
    for i := 1; i < Length; i++ {
        array[i-1] = array[i]
    }
    Length--
    [1,2,3,4]
             i
    [1,2,3,4]1
     ^_|
    [2,2,3,4]
//////////////////
    [2,2,3,4]2
       ^_|
    [2,3,3,4]2
//////////////////
    [2,3,3,4]3
         ^_|
    [2,3,4,4]3
         ^
         Length = 3
Esta es una manera de hacerlo

Otra manera de hacerlo es con slices de array, 

array[1:]
Pero esto no se que haga realmente (Es decir su O N)

Para insertar en cierto indice es parecido al problema de borrar, tenemos que hacer un ciclo

array := [1,2,3,4,5,6,7,8,9,_]
array.PushAt(11,2)
// Pushearemos el elemento 11 en el indice 4
    for i := Length-1; i <= idx; i-- {
        array[i+1] = array[i]
    }
    array[idx] = nuevoElemento
Primera iteracion
                 v
[1,2,3,4,5,6,7,8,9,_]
                 |_^

                 v
[1,2,3,4,5,6,7,8,9,9]
                 |_^

Segunda iteracion
               v 
[1,2,3,4,5,6,7,8,9,_]
               |_^

               v
[1,2,3,4,5,6,7,8,8,9]
               |_^

Tercera iteracion
             v 
[1,2,3,4,5,6,7,8,9,_]
             |_^

             v
[1,2,3,4,5,6,7,7,8,9]
             |_^

Cuarta iteracion
           v 
[1,2,3,4,5,6,7,8,9,_]
           |_^

           v
[1,2,3,4,5,6,6,7,8,9]
           |_^

Quinta iteracion
         v
[1,2,3,4,5,6,6,7,8,9]
         |_^

         v
[1,2,3,4,5,5,6,7,8,9]
         |_^

array[4] = 11
Length++


Delete at
    for i := Length-2; i <= idx; i-- {
        array[i-1] = array[i]
    }

array := [1,2,3,4,5,6,7,8,9,_]
array.DeleteAt(6)

Primera iteracion
                 v
[1,2,3,4,5,6,{7},8,9,_]
                 |_^

                 v
[1,2,3,4,5,6,7,8,9,9]
                 |_^
