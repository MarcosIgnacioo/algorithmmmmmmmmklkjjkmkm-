# Grafos

Todo lo que hemos estado haciendo es realmente un grafo

Por ejepmlo los árboles son grafos, sin embargo no todos los grafos son arboles!

Un grafo es simplemente nodos que pueden o no tener conexiones entre si

Por ejemplo

        (5)--+
         ^   v
         |  (6)
         v
        (7)
              (9)

Cada numero entre parentesis es un nodo, (o también conocidos mas formalmente como vértices). No hay una regla de que si un nodo debe de tener solamente x cantidad de conexiones, simplemente las puede tener o no tener, en la direccion que quiera, pudiendo ser bidireccional o unidireccional, teniendo peso en los caminos o no, etc

Pero como cada nodo puede tener diferentes cosas, existen maneras de poder clasificar a un nodo o un grafo

Ciclos

Si es posible realizar la visita de 3 (incluyendo el inicial) o mas nodos y volver al nodo inicial esto se considera un ciclo

Ejemplo de un ciclo

+-+    
| | +-+
v + v |
A | C |
^ | ^ |
| | | |
| | | |
v +-v |
D   B |
^-----+

A,B,C,D,A

Aciclica 

Un grafo que no contiene ciclos

Conectada

Cada nodo puede alcanzar a cualquier otro nodo, es decir, que a partir de cualquier nodo puedes ir a otro

Esto es un ejemplo de un grafo conectado, porque desde cualquier nodo puedes ir a cualquier otro, por ejemplo si estas en D y quieres llegar a A simplemente tienes que ir a C y luego ir a A

A<->B
^   ^   
|   |       
v   v     
C<->D


Esto es un ejemplo de un grafo que no esta conectado, porque a pesar de que los nodos A,B,C,D pueden ir de unos a otros, desde el nodo F no podemos ir al nodo A porque no tiene una direccion de regreso
A<->B
^   ^   
|   |       
v   v     
C<->D
|
v
F

Direccionado

Cuando las conexiones entre los nodos tienen direcciones
(Twitter)

Todas las conexiones de los nodos tienen una o mas direcciones

A<->B
^   ^   
|   |       
v   v     
C<->D
|
v
F


No-direccionado

Cuando no existen conecxiones entre los nodos
(facebook)

Cuando no existe una direccion entre las conexiones de los nodos
A---B
|   |   
|   |       
|   |     
C---D
|
|
F

Con peso

Cuando los caminos cuestan n cantidad de x "energia" de recorrer

Por ejemplo cuando vas de un lugar a otro en la ciudad te toma mas o menos tiempo teniendo en cuenta que puede haber mucho o poco trafico

Aqui ir de A a B cuesta 50, tanto de ida como de regreso, sin embargo, para ir de B a C cuesta 60 pero para regresar cuesta 30

*1      *2
  50     60 >
A<--->B ---- C
       < 30

*1 Es una relacion simetrica porque los costos de ambas direcciones son iguales
*2 Es una relacion asimetrica porque los costos de ambas direcciones son diferentes

DAG

Directed
Acyclic
Graph

(Grafo direccionado aciclico)

Es decir un grafo que tiene caminos con direcciones pero no se puede regresar al nodo inicial cuando se recorren 3 nodos (contando al inicial)

Esto es un ejemplo de un grafo DAG porque no podemos realizar ni un solo ciclo, de A podemos ir a D y de D a C pero ahi nos quedamos, de D podemos ir a C y ahi nos quedamos, pero podemos ir a A y volver a D, sin embargo esto seria simplemente recorrer solamaente 2 nodos asi que no se considera un ciclo, y de B solo podemos ir a A y ya no podriamos regresar y tabmien ir a C pero ya no podriamos regresar 

A<--B
^   |   
|   |       
v   v     
D-->C

En este grafo lo mas lejos que podriamos viajar seria de B a C pero de ahi no podemos volver a B, ni si quiera podriamos volver a D

Mas terminologia

Nodes: Los nodos (vertices)

Edges: Las conexiones entre los nodos

Big O en graphs es generalmente representada con V y E (vertices y edges) por lo que la formula se ve masomenos asi

O(V*E)

En los grafos podemos representarlos de dos maneras diferentes (siendo la primera la mas usada)

Aunque la Matrix puede ser realmente util en dadas situaciones la mayoria de las veces se va a utilizar una lista de adyacencia porque es la que mejor rendimiento tiene, porque la matrix es O(V^2)

    Adjacency List

    Adjacency Matrix


##   Adjacency List

(0)-->(1)
 ^ \   ^
 |  \  |
 |   \ |
 |    v|
(2)<--(3)

0 tiene los siguientes Edges:
    1, 3
1 NO tiene edges
    nil
2 tiene el siguiente edge
    0
3 tiene los siguiente edges
    2,1

Una lista de adyacencia es un arreglo, en el que los indices van a matchear al valor del nodo que van a representar
Por lo que el nodo con el valor de 0 va a ser almacenado en el indice 0 del arreglo

Por ejemplo, asi es como lo guardariamos adjacency list

[
    i=0:[
        Edges {
            To:1,
            Weight:10
        }, 
        Edges {
            To:3,
            Weight:12
        }, 
    ],
    i=1:[
        nil
    ],
    i=2:[
        Edges {
            To:0,
            Weight:15
        }
    ],
    i=3:[
        Edges {
            To:1,
            Weight:5
        }, 
        Edges {
            To:2,
            Weight:6
        }, 
    ],
]

Un nodo el cual ya no se puede mover a nada se conoce como terminal

##   Adjacency Matrix
Una matriz de adyacencia es casi lo mismo pero no lo mismo. 

[
         0  1   2  3
    i=0:[0, 10, 12,0],
    i=1:[0, 0,  0, 0],
    i=2:[15,0,  0, 0],
    i=3:[0, 5,  6, 0],
]

Aqui lo que hacemos es tener un arreglo en cada indice correspondiente al valor del nodo y dentro de ese arreglo rellenar con los pesos que haya en los edges correspondientes a los nodos con los que se conectan los cuales siguen la misma regla dae lo de los indices representan el valor del nodo


### BFS In Matrix

    1
(0)<->(1)
 | \
 |  \
4|   \ 5
 |    \ 
 v     \
(2)--->(3)----->(4)
    2       5


[
         0  1  2  3  4
    i=0:[0, 1, 4, 5, 0],
    i=1:[1, 0, 0, 0, 0],
    i=2:[0, 0, 0, 2, 0],
    i=3:[0, 0, 0, 0, 5],
    i=4:[0, 0, 0, 0, 0],
]


Como estamos recorriendo un grafo queremos generar un path de nuestro recorrido por lo que haremos el BFS de manera un poco diferente

Recordemos primeramente lo que haciamos en el BFS

Teniamos nuestro arbol binario, entonces lo que haciamos era crear una cola, meter a la cola la root de nuestro arbol binario (haciendo que la length de nuestra cola aumentara a 1). Despues iniciabamos un ciclo en el que desencolabamos a nuestro elemento que estaba primero en la cola, revisabamos si tenia el valor del que estabamos buscando, si lo tenia pues returneabamos true, si no, lo que haciamos era checar por los nodos hijos del que acababamos de sacar de la cola, si existian los metiamos a la cola y si no pues no los metiamos. Como seguian habiendo hijos pues se metian a la cola y la Length de ella aumentaba por lo que nuestro ciclo seguia existiendo, hacer esto permitia que fueramos yendo nivel por nivel, y de izquierda a derecha porque 

        if curr.Value.(*BinaryNode).Left != nil {
			q.Enqueue(curr.Value.(*BinaryNode).Left)
		}
		if curr.Value.(*BinaryNode).Right != nil {
			q.Enqueue(curr.Value.(*BinaryNode).Right)
		}
Tenemos el de la izquierda primero

Okay, como hacemos esto pero para una matriz 

[
         0  1  2  3  4
    i=0:[0, 1, 4, 5, 0],
    i=1:[1, 0, 0, 0, 0],
    i=2:[0, 0, 0, 2, 0],
    i=3:[0, 0, 0, 0, 5],
    i=4:[0, 0, 0, 0, 0],
]

Lo que hacemos sera

Tener un arreglo de los nodos visitados

Lo que vamos a ocupar

    matrix := [
             0  1  2  3  4
        i=0:[0, 1, 4, 5, 0],
        i=1:[1, 0, 0, 0, 0],
        i=2:[0, 0, 0, 2, 0],
        i=3:[0, 0, 0, 0, 5],
        i=4:[0, 0, 0, 0, 0],
    ]
    prev = [-1, -1, -1, ...]
    seen = [ f,  f,  f, ...]
    queue = [0]
             ^
             source node in here

seen[0][0] = true
queue.Enqueue(0)
path.push(0)
for len(queue) > 0 {
    curr := queue.Dequeue
    if curr == needle {
        return true
    }
    for i:= 0; i < len(matrix[curr]); i++ {
        if seen[curr][i] == matrix[curr][i] {
            continue
        }
        if matrix[curr][i] != 0 {
            seen[curr][i] = true
            path.push(i)
            queue.Enqueue(i)
        }
    }
}


Si queremos hacer BFS lo que tenemos que hacer es meter nuestros
 
