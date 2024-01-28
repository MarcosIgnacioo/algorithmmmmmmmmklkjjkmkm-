# Tries

Tambien conocidos como prefix trees, digital tree, 

La manera mas facil de visualizar un trie es el autocompletado

Por ejemplo cuando escribimos la letra a que palabras nos va a dar el autocompletado de nuestra pc

Lo bueno es que puede hacerlo a O(1) si las letras presionadas ya son varias

abecedario

Estos son todos los hijos que puede tener nuestro trie, solamente si 

a b c d e f g h i j k l m n o p q r s t x y z


Por ejemplo si quisieramos agregar la palabra "cat" esto seria lo que hariamos

Uno de nuestros 26 hijos del arreglo del abecedario, en este caso la c, simplemente sera un hijo (individual) 
    
Si estan () significa que 
    IsWord = true
c 
|
a->r->(d)
|
(t)->t->l->(e)
|
(s)


Si queremos que se muestren las palabras que estan
