# Big O

Es la manera de categorizar el uso de memoria o tiempo requerido por nuestros algoritmos en base al input. No es una medida exacta, no dira la cantidad de ciclos de CPU necesita, mas bien generalizara el crecimiento de complejidad en tu algoritmo

Es decir es la manera de saber como reaccinara nuestro algoritmo ante x input

Asi que si alguien dice 

OH tu algoritmo es O de N, significa que nuestro algoritmo va a crecer linealmente 

Big O en pocas palabras

Que tanto cambia la computacion y el uso de memoria de tu algoritmo segun el input data crece

En el mundo real el crecimiento de uso de memoria no es algo gratis, puesto que cuesta tiempo asignar cierta memoria, pero cuando pensamos en algoritmos no pensamos necesariamente en eso.

En lenguajes como Go o JS esto es aun mas costoso, pues como la memoria se limpia con un garbage collector y muchas veces se queda pues esto hace que todo sea muchisimo mas lento de lo que deberia de serlo

Un ejemplo de crecimiento lineal es una funcion que imprima los caracteres que conforman un string dado, pues, esto se hace con un ciclo que se termina cuando se haya llegado al final del string, el cual tiene una longitud, y mientras mas largo el string, mas largo es su longitud, ergo, mas ciclos que se tienen que hacer para poder imprimir cada caracter de el.

Esto seria O N

func printChars(n string) {
    for _, v := range n {
        fmt.Println(v)
    }
}

Porque si tenemos un input que tiene 50 caraacteres, y luego hacemos un input que tenga otros 50 caracteres es decir 100 caracteres en total, nuestro algoritmo habra crecido un 100% por lo que sera el doble de lento que la primera vez que lo usamos, esto es un crecimiento lineal


Aqui seria O N^2
Porque estamos haciendo la misma operacion de nuevo por lo que es potenciado

func printChars(n string) {
    for _, v := range n {
        fmt.Println(v)
    }
    for _, v := range n {
        fmt.Println(v)
    }
}

Generalmente O N^2 es mas lento que O N, pero hay casos en los que esto no es asi como en el caso de que el input sea mas pequeño

Pero que pasa si tenemos un if statement? Que O N es?, bueno, para  saber esto tenemos que hacer una serie de pasos. 

    El primero es ponernos en la peor situacion posible, en este caso la peor situacion posible es que la i este hasta el final, porque significaria que tendriamos que recorrer todo el arreglo.
    
    Entonces, en el peor escenario lo que tenemos que hacer es recorrer todo el arreglo, por lo que, es realmente como si recorrieramos todo el arreglo como la primera vez, entonces aun con el if statement esto sigue siendo O N

    En entrevistas siempre te preguntan por el peor escenario posible 

func printChars(n string) {
    for _, v := range n {
        if (v == "i") {
            return
        }
        fmt.Println(v)
    }
}

Por que no usamos constantes?

Porque las cionstantes son no coool para esto

Por ejemplpo

| N  | Exponente | Constante|
|----------------| -------- | |
|5   | 2         | N/A      |
|----|-----------|-----------
|5   | N/A       | 5        |
|----|-----------|-----------

Esto se debe a que Big O se centra en el comportamiento asintótico, es decir, cómo se comporta el tiempo de ejecución del algoritmo a medida que la entrada tiende hacia el infinito.

Al ignorar constantes y términos de bajo orden, Big O proporciona una descripción general y simplificada del crecimiento del tiempo de ejecución en función del tamaño de la entrada. Esto hace que sea más fácil comparar y clasificar algoritmos en términos de su eficiencia relativa sin preocuparse por detalles específicos de implementación o condiciones exactas de entrada.


Tipos de Big O (Menos complejidad a Mayor complejidad)

O(1) Siempre hace la misma cantidad de operaciones sin importar el input

O(logN) Es como O(1)

O(N) Es crecimiento lineal

O(NlogN) 

O(N^2)

O(2^N)

O(N!) (Esto no se puede hacer en computadoras hoy en dia asi q ni pedo)

# Ejemplos de O N

O(N^2)


Los ciclos anidados son exponenciales porque estamos hablando de que por cada elemento del string se va a recorrer todo ese string otra vez, es como el area de un cuadrado

Por ejemplo si pusieramos el siguiente input : "Hola"
Se computaria asi

H o l a
H H H H 
o o o o
l l l l
a a a a

y recordemos area de un cuadrado es igual a l^2


func printChars(n string) {
    for _, v := range n {
        for _, v := range n {
            fmt.Println(v)
        }
    }
}

O(N^3)
Y pasa lo mismo para 3 ciclos anidados

Basicamente X cantidad de ciclos anidados sera el expontente de N en cuanto a complejidad

Multiplicar una matriz tambien es N^3 porque ocupas usar 3 ciclos

func printChars(n string) {
    for _, v := range n {
        for _, v := range n {
            for _, v := range n {
                fmt.Println(v)
            }
        }
    }
}
