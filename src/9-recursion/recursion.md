# Recursion


func foo(n int) int {
    // Base case
    if (n == 1) {
        return n
    }
    // We shall recurse
    return n + foo(n-1)
}

Si conocemos nuestro base case bien la recursion se vuelve extremadamente simple asi qeu es la parte que tenemos que saber bien asi bien

Lo que creo que hace esta funcion

TOma un numero y suma todos sus numeros hasta el anterior

asi 

5 + 4 + 3 + 2 + 1


Todas las funciones tienen su 

Return Adress

    Es la direccion en memoria donde se llamo al a funcion, pues ahi es donde va a retornar el resultado

Return Value

    El valor que vamos a retornar a la return adress

Arguments

    Los argumentos que le pasamos a la funcion


func foo(n int) int {
    // Base case
    if (n == 1) {
        return n
    }
    // We shall recurse
    return n + foo(n-1)
}

*En donde llamemos a la funcion foo(5) en nuestro codigo

 rA      |   rV  | A
-------------------
 foo(5)* |       | 5      
         |  n + ?|       

-------------------
 foo(4)  |       | 4
 dentro  |  n + ?|
 de foo(5) 

-------------------
 foo(3)  |       | 4
 dentro  |  n + ?|
 de foo(4) 

-------------------
 foo(2)  |       | 4
 dentro  |  n + ?|
 de foo(3) 

-------------------
 foo(1)  |       | 4
 dentro  |  n + ?|
 de foo(2) 


QUe pasa cuando llega hasta abajo


Parte del codigo donde foo(5) fue llamado
^___________
 rA      |  | rV  | 
-------------------
 foo(5)* |       |       
         | 5 + 10|<-------+ ya nos devolvieron el valor de ?
                          | que es 10
-------------------       |
 foo(4)  |       |        |
 dentro  |  4 + 6|<-------+ ya nos devolvieron el valor de ?
 de foo(5)                | que es 6
                          |
-------------------       |
 foo(3)  |       |        |
 dentro  |  3 + 3|<-------+ aqui ya nos devolvieron el valor de ? 
 de foo(4)                | que es 3
                          | 
-------------------       |
 foo(2)  |       |        |
 dentro  |  2 + 1|<---+---+
 de foo(3)            |
                      | aqui cambiamos el ? que estaba antes por el resultad
-------------------   | de abajo (1)
 foo(1)  |       |    |
 dentro  | n == 1| ---+
 de foo(2) asi que hacemos el 
            return n (es 1)


func foo(n int) int {
    // Base case
    if (n == 1) {
        return n
    }
    // We shall recurse
    return n + foo(n-1)
}

Como explicar la recursividad es dificil al menos por escrito voy a intentar pasar lo que entiendo a palpabras o sea si

Nosotros tenemos nuestro stack de donde se llaman las funciones y en este stack kas funciones tienen una direccion en memoria de donde deben de retornar el resultado que den, tienen tambien el valor del resultado y los argumentos 

Algo asi

+--------------+
| rA | rV | Ar |
|----+----+----+
|    |    |    |
|----+----+----|
+--------------+

Cuando llamamos auna funcion esta nos devuelve un valor y ese valor se guarda en la direccion de memoria de donde fue llamada dicha funcion

La recursividad parte de aprovechar esto, en el ejemplo de 

func foo(n int) int {
    // Base case
    if (n == 1) {
        return n
    }
    // We shall recurse
    return n + foo(n-1)
}

Ingresaremos un numero y nos devolvera la suma de 1 hasta n, es decir si ponemos 5 nos regresara el resultado de 1+2+3+4+5 = 15

Esto pasa porque hace el llamado a la funcion restando en una unidad al argumento que le pasamos, y pasa esto hasta qeu esa resta de 1, entonces, ya se regresa y ese valor lo usamos en cada como regresada de la recursividad


