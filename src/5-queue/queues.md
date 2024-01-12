# Queues

Una queue es una fifo structure, es decir que el primer elemento que entra es el primero que sale, es literalmente una cola de una tienda, si eres el primero en formarte eres al primero que van a atender y por ende el primero que va a poder salir

(A) -> (B) -> (C) -> (D) -> (E)
 ^                           ^
 |                           |
 head                        tail

Si vamos a hacer un pop a nuestra cola el elemento que saldria seria al que nuestra head apunta (el primero)

Y seria tal que asi

outFirst = head.next
newFirst = outFirst.next
head.next = newFirst
outFirst.next = null

Y para insertar seria de la siguiente manera

oldLast = tail.next
oldLast.next = newNode
tail.next = newNode



