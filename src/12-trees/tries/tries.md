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


Cada palabra va a tener un abecedario para ella misma

type TriesNode struct {
	IsWord     bool
	Characters [26]*TriesNode
}
Porque asi se puede hacer todas las combinaciones posibles con esa letra

De tal manera que cada vez que queramos insertar una palabra lo que haremos simplemente sera hacer que el caracter en el arreglo de TriesNode en el espacio al que le correspopnda calculado con la siguiente formula

func GetCharCode(char rune) int {
	a := int([]rune("a")[0])
	return int(char) - a
}

entonces al final los nodos se ven algo asi 

a ----- a(nil), b(nil), c(nil), d(nil) , e(nil), f(nil), g(nil) 
b ----- a(nil), b(nil), c(nil), d(nil) , e(nil), f(nil), g(nil) 
c ----- a(nil), b(nil), c(nil), d(nil) , e(nil), f(nil), g(nil) 
d ----- a(nil), b(nil), c(nil), d(nil) , e(nil), f(nil), g(nil) 
e ----- a(nil), b(nil), c(nil), d(nil) , e(nil), f(nil), g(nil) 
f ----- a(nil), b(nil), c(nil), d(nil) , e(nil), f(nil), g(nil) 
g ----- a(nil), b(nil), c(nil), d(nil) , e(nil), f(nil), g(nil) 

Y pues por ejemplo para insertar la palabra "cat" lo que se hace es lo siguiente


Vamos nuestro nodo raiz y checamos 

		if curr.Characters[idx] == nil {
			curr.Characters[idx] = NewTrie()
		}
		curr = curr.Characters[idx]

Si nuestra primera letra no tiene ningun registro previo pues le creamos su abecedario. Luego entramos a ese abecedario

Esto se puede gracias a que lo hacemos d una manera inteligente

Porque al inicio de cada ciclo lo que hacemos es obtener el 

var idx int
curr := tn
wChars := []rune(word)
wLength := len(wChars)

for i := 0; i < wLength; i++ {
    ...
}

Creamos nuestra variable para controlar los indices correspondientes a nuestros caracteres

Luego creamos una variable auxiliar que nos va a servir para ir recorriendo nuestro trie, entonces

Lo que hacemos ya simplemente es obtener el indice correspondiente en el abecedario de la letra que estemos por insertar actualmente

    Si aun no esta creada la creamos 
    Si ya esta creada simplemente vamos a la siguiente
for i := 0; i < wLength; i++ {
    idx = GetCharCode(wChars[i])
    if curr.Characters[idx] == nil {
        curr.Characters[idx] = NewTrie()
    }
    curr = curr.Characters[idx]
}

Y al final nuestro current va a quedarse en el ultimo caracter de nuestra palabra que queremos insertar, por lo que tenemos que marcarlo ya como una palabra

curr.IsWord = true


Para buscar es casi lo mismo, solamente que pues si una letra aun no se ha registrado, es decir vale nil, pues no existe esa palabra
for i := 0; i < wLength; i++ {
    idx = GetCharCode(wChars[i])
    if curr.Characters[idx] == nil {
        return false
    }
    curr = curr.Characters[idx]
}
return curr.IsWord

Por que no usarmos return true?
porque puede ser que noostros insertemos la palabra "hola" pero si buscamos hol, como estos caracteres si se encunetran registrados porque son los ocmpopnentes de hola, pues si usaramos el return true diriamos que hol es una palabra cuando pues aun no lo es, por lo que checamos si el ultimo caracter que buscamos forma la palabra para asegurarnos siempre de que sea una busqueda exacta
