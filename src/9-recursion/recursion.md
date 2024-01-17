# Recursividad
 
La manera mas basica de pensar sobre la recursividad es una funcion que se invoca a si misma hasta que un problema este resuelto. Esto normalmente involucra algo que se llama base case. El base case es el punto en el que el problema se resuelve 

El ejemplo mas simple de recursividad es el siguiente

function foo(n int) int {
    // Base case (Cuando nuestro problema se soluciona)
    if n == 1 {
        return n
    }
    // Si no cumplimos con nuestro base case tenemos que recursar
    return n + foo(n - 1)
}

Esto lo que hace es sumar todos los antecesores de un numero e incluyendolo a el mismo

Por ejemplo

n = 5

Lo que hace la funcion es 

1 + 2 + 3 + 4 + 5 = 15

Esto es posible gracias a la recursividad, veamos la funcion mas de cerca

Supongamos que llamamos al a funcion con el parametro 5

foo(5)

function foo(5) int {
    if 5 == 1 { X----------------- Vemos si nuestro parametro cumple 
                                    con nuestro base case, NO LO HACE
        return n
    }
    return 5 + foo(5 - 1) <------- Como no cumplimos con nuestro base case 
                                llamamos a la funcion de nuevo (Recursamos)
}

Esto lo que hace es poner en nuestro stack de funciones ahora la nueva funcion y dejamos la anterior en "pausa" porque esta aun queda por retornar algo al lugar donde la llamamos (nuestro codigo) pero una parte del valor de retorno es el retorno del llamado de otra funcion

function stack

funcion  retorno

foo(5)   5 + ?
foo(4)   4 + ?

Ahora foo(5) llamo a foo(4) por lo que hacemos lo mismo

function foo(4) int {
    if 4 == 1 { X----------------- Vemos si nuestro parametro cumple 
                                    con nuestro base case, NO LO HACE
        return n
    }
    return 4 + foo(4 - 1) <------- Como no cumplimos con nuestro base case 
                                llamamos a la funcion de nuevo (Recursamos)
}

Y ahora foo(4) llamo a foo(3)


funcion  retorno

foo(5)   5 + ?
foo(4)   4 + ?
foo(3)   3 + ?

Ahora foo(4) llamo a foo(3) por lo que hacemos lo mismo

function foo(3) int {
    if 3 == 1 { X----------------- Vemos si nuestro parametro cumple 
                                    con nuestro base case, NO LO HACE
        return n
    }
    return 3 + foo(3 - 1) <------- Como no cumplimos con nuestro base case 
                                llamamos a la funcion de nuevo (Recursamos)
}


funcion  retorno

foo(5)   5 + ?
foo(4)   4 + ?
foo(3)   3 + ?
foo(2)   2 + ?

Ahora foo(3) llamo a foo(2) por lo que hacemos lo mismo

function foo(2) int {
    if 2 == 1 { X----------------- Vemos si nuestro parametro cumple 
                                    con nuestro base case, NO LO HACE
        return n
    }
    return 2 + foo(2 - 1) <------- Como no cumplimos con nuestro base case 
                                llamamos a la funcion de nuevo (Recursamos)
}

funcion  retorno

foo(5)   5 + ?
foo(4)   4 + ?
foo(3)   3 + ?
foo(2)   2 + ?
foo(1)   1

Ahora foo(2) llamo a foo(1) por lo que hacemos lo mismo

function foo(1) int {
    if 1 == 1 { <----------------- Vemos si nuestro parametro cumple 
                                    con nuestro base case, en este caso si lo hace!!!
        return n
    }
    return 1 + foo(1 - 1) X------- Como si cumplimos con nuestro base case 
                                    esto no recursa ya
}

Entonces como ya hay una funcion que tiene un valor real para retornar se iran pasando el valor una a la otra (Porque la direccion de retorno son las propisa funciones recursando puesto que son las propias funciones quienes hicieron el llamado de ellas por lo que el valor de retorno les pertenece a ellas en ese espacio en memoria)

Por lo que se ve masomenos asi

Ejemplo: (Las x es solamente para que no se ponga de color verde esto)


func main() {
    foo(5) <----------------------------+
}                                      ?|
                foo(5)   5 + ?  <--+----+
x                                 ?|
            foo(4)   4 + ? <-+-----+
x                           ?|
        foo(3)   3 + ? <-+---+
x                       ?|
    foo(2)   2 + ? <-+---+
x                   1| 
foo(1)   1 <---------+

func main() {
    foo(5) 
}                                     
                foo(5)   5 + ? 
x                             
            foo(4)   4 + ? 
x                         
        foo(3)   3 + ?
x                    
    foo(2)   2 + 1 <-+
x                   1| 
foo(1)   1 <---------+

x               foo(5)   5 + ? 
x                             
            foo(4)   4 + ? 
x                         
        foo(3)   3 + 3 <-+
x                       3|
    foo(2)       3 <-+---+
x                   1| 
foo(1)   1 <---------+

x               foo(5)   5 + ? 
x                             
            foo(4)   4 + 6 <-+
x                           6|
        foo(3)       6 <-+---+
x                       3|
    foo(2)       3 <-+---+
x                   1| 
foo(1)   1 <---------+

x               foo(5)   5 + 10 <--+
x                                10|
            foo(4)      10 <-+-----+
x                           6|
        foo(3)       6 <-+---+
x                       3|
    foo(2)       3 <-+---+
x                   1| 
foo(1)   1 <---------+

func main() {
    foo(5) <----------------------------+
}                                     15|
x               foo(5)   15 <-----------+
x                                10|
            foo(4)      10 <-+-----+
x                           6|
        foo(3)       6 <-+---+
x                       3|
    foo(2)       3 <-+---+
x                   1| 
foo(1)   1 <---------+

foo(5) -> foo(4) -> foo(3) -> foo(2) -> foo(1) (cumple con base case)

main() <- 15 - foo(5) <- 10 - foo(4) <- 6 - foo(3) <- 3 - foo(2) <- 1 - foo(1)

La recursividad tiene 3 estados

El pre

    En este caso nuestro pre es el sumar n a lo que nos de foo(n-1)

El recurse
    
    Basicamente el recursar que en nuestro caso es el foo(n-1)

El post
    
    Podemos hacer algo con el resultado de la recursion, por ejemplo

var list ArrayList

function foo(n int) int {
    if n == 1 {
        return n
    }
    result := n + foo(n - 1)
    fmt.Println(result)
    list.add(result)
    return result
}

Antes lo que haciamos era retornarlo directamente pero pues podemos guardarlo en una variable o loggearlo o guardar cada elemento en una lista

Esto es algo muy importante para cosas como el pathing

Implementacion!!!!

MazeSolver baby

Tenemos una matriz con un laberinto

E = Exit
S = Start

"######################E##"
"#                     x #"
"# xxxxxxxxxxxxxxxxxxxxx #"
"# x                     #"
"##S######################"
Para programar esto con recursividad tenemos que tener en cuentra nuestros base cases, 

    1.-  Nuestra posicon ya es E
    2.-  Hay un muro en la direccion donde nos queremos mover
    3.-  Ya hemos estado en ese lugar
    4.-  Nos salimos del mapa

Y nuestro recursive case simplemente lo que hace es checar las 4 direcciones a las que se puede mover
        
         p
         o
         T
         |
    tfeL-#-Right
         |
         B
         o
         t
         t
         o
         m
    
Cuando juntamos ambas cosas boom programa comploeto gamer
