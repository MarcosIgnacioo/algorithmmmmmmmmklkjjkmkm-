# Depth-First Find


           <=5<
          /    \            
         4      6         
        / \    / \       
       3  2    7  8     

Todo en el lado izquierdo de un arbol es menor o igual al nodo padre

Todo en el lado derecho del nodo padre es mayor que él

Esta es la base para poader organizar nuestros arboles

Y en general se applica a cada rama

El elemento que se ponga en la izquierda del nodo debe de ser menor o igual a su nodo padre correspondiente y se pone en la derecha si es mayor a su nodo padre correspondiente


Pongamos de ejemplo este arbol, en el caso de que nosotros queramos poner un elemento igual a nuestra root en nuestro arbol es algo que podriamos hacer, sin embargo este nodo no podria tener un valor a la derecha, puesto que para que tuviera que tener un nodo a su derecha tendria que ser mayor a 15 pero como 15 es nuestra raiz, dicho valor se iria directamente a la rama derecha de nuestra raiz

//        15
//       / \
//      15  17
//     /  \
//    8    [?]

Para evitar complicaciones trabajaremos con este
//        15
//       / \
//      5  17
//     / \
//    3   8

El uso funcional de estos arboles ordenados de esta manera es que podemos encontrar las cosas de manera mas rapida y podemos facilitar el agregar las cosas


//        15
//       / \
//      5  17
//     / \
//    3   8
    /   \
   5<n<=3  8>n>3

Para poder insertar en las ramas del nodo con valor 3 
    En la rama izquierda del nodo con valor 3 debe de ser menor o igual a 3
    En la rama derecha del nodo con valor 3 debe se der menor a 8 y mayor que 3


//        15
//       / \
//      5  17
//     / \
//    3   8
        /   \
     n<=8>3  15>=n>3

En el peor de los casos esto tardaria O(N), pero en realidad es 
O(H) porque en todo caso depende realmente de la altura de nuestro arbol mas que del input

Pero generalmente sera logN, porque no visitaremos cada nodo

Si tenemos un arbol completo la altura de nuestro arbol es logN

Si tenemos un arbol binario no coompleto, asi por ejemplo

             5
            /
           5
          /
         5
Nuestra altura si que será N

Entonces nuestra O seria entre LogN y N


Hay 3 casos  importantes que tenemos que considerar a la hora de borrar un elemento de nuestro arboll

El primero y el mas facil es que nuestro nodo no tenga hijos. por lo que simplemente podemos borrarlo asi bien facil

El segundo es que tenga un hijo, por lo que tenemos que enlazar el padre de nuestro nodo con el hijo del nodo

        +--->65
        |   /
        |  55
     +--+ /
     |   45
     +---^
Y si tiene dos hijos lo que haremos sera trasversar hacia la derecha todo el rato o la izquierda hasta que lleguemos al valor mas grande o pequeño, y reemplazaremos a nuestro nodo con ese valor. Como sabemos cuando ir por la derecha y la izquierda

Simplemente lo que tenemos que hacer es saber si el nodo que queremos borrar es de derecha o izquierda, eso se puede saber facilmente comparandolo con su nodo padre, siendo que si el padre es mayor que nuestro nodo signfica que estamos en la derecha por lo que ocuparemos reemplazar nuestro nodo por uno mas pequeño porque asi podemos mantener los nodos a la derecha, mientras que si estamos a la izquierda ocupamos a nuestro valor mas grande porque asi podemos mantener an uestros valores a la izquierda aqui un ejemplo

	           15
	         /    \
	        /     \
	        5     27<----+
	      /  \   /       |
	     /   \  22       31
	    /    \ 20 25   29 36
	   3     8
	  /    / |
	 2

Entonces como estamos a la derecha empezamos a buscar el valor mas pequeño de nuestro lado que se encontrara la izquierda del todo, lo que al final seria 20, por lo que reemplazamos las cosas y boom

Pero si estuvieramos a la izquierda por ejemplo si quisieramos borrar el 22 lo que pasaria seria que iriamos por el valor mas grande a partird de ese nodo, que en ese caso es 25 y boom
