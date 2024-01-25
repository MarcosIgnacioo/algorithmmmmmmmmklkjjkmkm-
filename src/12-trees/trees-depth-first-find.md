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


