# Linked Lists

Cuando hagamos inserts en listas lo qeu hacemos primero siempre es conectar el nuevo nodo a l indice donde estemos conectandolo y luego desconectamos las conexiones viejas con las nuevas

Por ejemplo pongamos que queremos insertar un elemento F en el lugar de C, entonces lo que queremos hacer es que el next de B apunte a F y el prev de F apunte a F, luego que el prev de C apunte a F y luego que el next de F apunte a C

B <-> C <-> D 
      ^
      F

Orden para hacer las cosas

1.- Hacer que el Prev de F apunte a B y que el Next de F apunte a C
      -+ t
      || x
      || e
      v+ n
B <-> C| <-> D 
 ^    ^|
 |____F_
  prev

2.- Hacer que el Next de B apunte a F y que el prev de C apunte a F

B <-> F <-> C <-> D 
