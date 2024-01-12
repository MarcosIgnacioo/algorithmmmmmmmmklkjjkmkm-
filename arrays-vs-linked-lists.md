# Arrays vs linked lists

Arrays

Los arrays tienen la ventaja de poder acceder a los valores mediante su index, lo que es muy rapido y fácil

Recordemos que los arreglos se encuentran en el stack pegaditos y para ir leyendo cada uno simplemente se hacia con la fórmula de 
    
    dir_memoria_array + (tamaño_de_tipo_de_dato * i)


    1000 + (8 * i)

Sin embargo los arreglos no  permiten insertar valores nuevos o borrar elementos a menos que hagas un despapaye en un ciclo (creando un nuevo arreglo y en el indice que quieres insertar el elemento insertarlo y en los demas poner los valores anteriores del arreglo pero respetando ahora el nuevo orden)

Obtener un elemento de un array es O 1
Cambiar un elemento de un array es O 1

Las linked lists por el otro lado nos permiten insertar y borrar elementos de manera sencilla, pero, el acceder a un elemento en especifico de nuestra lista es mas dificil que en un array, pues en un array tenemos los indices que con una simple operacion matematica permite tener el elemento en constant time sin importar el tamaño del  indice, sin embargo en una linked list tenemos que recorrer toda la lista hasta llegar a cierto punto

Situaciones en las que queremos usar listas enlazadas

Cuando queremos tener la posibilidad de insertar o eliminar elementos que esten en los bordes siendo fifo con las colas y no-fifo con los stacks

Cuando no vamos a querer obtener cierto elemento en especifico muchas veces 

Un ejemplo de cuando usar una cola (linked list) en el mundo real es cuando por ejemplo quieres tener solamente 5 conexiones a tu network, entonces lo que quieres en este caso es tener la capacidad de quitar a los primeros que llegaron y que siga el siguiente, o cuando tienes a muchos usuarios queriendo acceder a tu pagina web, seguramente quieras limitar eso a X cantidad y una lista es perfecta para eso porque en ningun momento vamos a requerir trasversar sobre ella simplemente queremos tener la capacidad de estar removiendo los usuarios de manera eficiente 

Una manera de trasversar en una linked list es la siguiente

let current := list.head

for i := 0; i < idx && current != nil ; i++ {
   current = current.next 
}    

return current.value
