# Linked Lists

En js tenemos esto

const a = []

Que aunque parezca un arreglo en realidad no lo es, porque un arreglo real solamente tiene el metodo de length  sin embargo en js con estos "arreglos" tenemos disponibles los metodos de push pop indexof etc

Entonces que es realmente? Es una linked list

Pero que es una linked list

Es una lista de nodos que estan conectados entre si

Node <T>
    value:  T
    next?:  Node<T>
    prev?:  Node<T>

(El prev solo esta si es una double linked list pero pues esta mejor que lo pongamos ya de una para que la lista sea mas usable)


(A)<->(B)<->(C)<->(D)<->(E)
 ^                       ^
 |                       |
 |head                   |tail

Una lista enlazada tiene dos propiedades fijas, head, que apunta al primer nodo de nuestra lista y tail que apunta a nuestro ultimo nodo

Las linked se alojan en el heap

Por que?
investigar

Lo chido de las linked lists es que las insertions y deletions pueden ser muy rapidas puesto que podemos hacer lo siguiente

Tenemos nuestra lista y queremos ingresar un nuevo elemento (F) que este entre A y B

(A)<->(B)<->(C)<->(D)<->(E)

Para hacer esto ocupamos hacer que:

    (A) apunte a (F) en su next: (A) -> (F)

    (F) apunte a (A) en su prev: (A) <- (F)
    (F) apunte a (B) en su next: (F) -> (B)

    (B) apunte a (F) en su prev: (F) <- (B)

Y hacer esto es en constant time porque simplemente estamos accediendo a 3 objetos no estamos recorriendo nada


Si quisieramos borrar C lo que tendriamos que hacer es lo siguiente

(A)<->(F)<->(B)<->(C)<->(D)<->(E)

Accedemos a C

B = C.prev
D = C.next
B.next = D
D.prev = B
// Hacemos que B ahora apunte a D y que D (su prev) apunte a B

C.prev = null
C.next = null

(A)<->(F)<->(B)<->(D)<->(E)

Esto tambien es constant time porque siempre vamos a modificar 4 enlaces de nodos y no importa el tamaño de nuestra linked list esta operacion va a costar lo mismo

O(1)

Para visitar un nodo especifico en una lista tenemos que hacerlo de manera lineal

current := head
for 0..5 {
    current = current.next
}

Operaciones rapidas en linked lists

    *
    Obtener head/tail
    Borrar en el principio o el final
    Preprend/Append (Insertar al inicio o al final)
        *Recordemos que tenemos un puntero fijo a la head y tail por lo que siempre es algo fijo por lo que es constant time

Operaciones lentas en linked lists
    
    *
    Insertar en el medio de la linked list
    Borrar en el medio de la linked list
        *Como tenemos que navegar en la lista la cual es una operacion muy costosa en cuanto a rendimiento



