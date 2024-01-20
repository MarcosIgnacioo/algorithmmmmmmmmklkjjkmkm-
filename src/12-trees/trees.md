# Trees

Los arboles son listas ramificadas

Por ejemplo 
         A
        /|\
       / | \
       B C D

El nodo A tiene de hijos a B, C y D

Un ejemplo de un arbol es el propio sistema de archivos de la computadora, cuando ponemos ls nos dicen varias cosas
~ ls

Desktop                    Public                     intelephense
Documents                  adf                        laraveltest
Downloads                  algorithmmmmmmmmklkjjkmkm- notas
Library                    asdf.go                    popo
Movies                     culo.go                    the-end-of-ramgelion
Music                      go                         your-name
Pictures                   hola.go

En este caso ~ es el padre de todos estos nodos (archivos)

En el codigo cada nodo se representa de la siguiente manera

type Node struct {
    Value interface {}
    Children [] *Node
}

Por ejemplo el DOM tambien es un arbol

             HTML
            /    \
           /      \
          Head     Body
         /        / |  \
        links   div div div

root - Top most node (Por ejemplo en los casos de arriba serian HTML en el caso del DOM y ~ en el del sistema de archivos)

height - La distancia entre el nodo root y el nodo hijo mas lejano

binary tree - Arbol con maximo 2 hijos y la menor cantidad de hijos permitidos es 0

general tree - Arbol de 0 a N cantidad de hijos

binary search tree - Un arbol con un orden de los nodos especifico

leaves - Un nodo sin hijos

balanced - Se dice que un nodo esta perfectamente balanceado cuando esta asi

        5
       / \
      3   8
     / \ / \
    1  4 7  9

Es decir cuando tenemos la misma distancia entre todos

branching factor - cantidad de hijos que tiene el arbol
