# Arrays Data Structure

La idea mas general de un arreglo es que es un espacio continuo en memoria
(conitnuo signficiando irrompible) que contiene cierta cantidad de bytes

Todo en las computadoras son solamente numeros, 0 y 1. Y la manera en la que la memoria es interpretada se basa en lo que noosotros le decimos alcompilador como queremos que lo interprete.

Por ejemplo, cuando ve un pedazo de memoria que tiene 4 bytes el compilador dice ohh voy a tratar esto como un numero solamente. Y asi es como nacen los integers de 32 bits (1 byte = 8 bits 4 bytes = 32 bits) pero pues esto es porque nosotros en los lenguajes podemos especificar si queremos que un numero sea i32 por ejemplo en go o rust, asi que no tiene ese significado hasta que nosotros le demos ese significado

Un arreglo es la version extendida de esto, pues tenemos un espacio continuo en memoria que va a almacenar un tipo especifico de datos



base_address = 1000  
element_size = 8 

index = 3
memory_position = base_address + (index * element_size)
>> 24

Okay como funciona el calculo de posiciones

Nuestro primer elemento de nuestro arreglo va a empezar en nuestra direccion de memoria donde alojamos nuestro arreglo, recordemos que las direcciones de memoria pues no indican exactamente tooodo lo que ocupa, nos indican donde empieza lo que este en ese espacio en memoria

Entonces por ejemplop tenemos un arreglo de unsigned integers de 8 bits
que se encuentra alojado en la direccion 1000

El elemento en el indice 0 esta ubicado justo en nuestra direccion de memoria, es decir, 1000, y como es de 8 bits por lo que abarcara otros 7 bits aparte (8 en total) asi:

     1000 1001 1002 1003 1004 1005 1006 1007
n    1    2    3    4    5    6    7    8

y asi tenemos nuestro primer elemento, pero que pasa con los siguientes como puede saber la computadora que tiene que ponerlos al lado? muy facil, matematicas pues multiplicamos el tamaño de nuestro tipo de dato que estamos guardando en el arreglo por el indice y le sumamos la direccion donde inicio, por lo que para el siguiente elemenot que se encontrá en el indice 1 su posicion en memoria se calcula de la siguiente manera

1000 + (i * 8)
1000 + (1 * 8)
1000 + 8
1008

y esta al lado de nuestro primer elemento en el indice 0

    1008 1009 1010 1011 1012 1013 1014 1015
n   8    9    10   11   12   13   14   15
    1    2    3    4    5    6    7    8

y asi hasta el final, ou yeaaaa entonces por que paso lo que paso aqui

ThePrimeagen lo que hizo fue crear un buffer de un array que es basicamente creo que como un espacio en memoria que puedes usar para crear un arreglo de un tipo en especifico pasandoselo como parametro al constructor del arreglo que quieras crear por ejemplo aqui

> const a = new ArrayBuffer(12)
> const a8 = new Uint8Array(a)

Le pasa el buffer y ahora (a8) alojara los elementos de su arreglo en el arraybuffer (a)

entonces si hacemos la operacion de asignar un valor en nuestro (a8) en el indice 0 lo que hara es poner un valor efectivamente en el indice 0 y nuestro arreglo con los bits se veria tal que asi

  > a8[0] = 45
  > [Uint8Contents]: <2d 00 00 00 00 00 00 00 00 00 00 00>,

Y aqui hizo lo interesante creo un arreglo de unsigned ints de 16 bits, es decir del doble del tamaño que nuestro primer arreglo, y almacenara las cosas en el mismo buffer (direccion de memoria) entonces como reacciona nuestra memoria cuando alojamos un valor en cierto indice en nuestro arreglo de 16 bits

  > a16[3]=65
  > [Uint8Contents]: <2d 00 00 00 00 00 41 00 00 00 00 00>,

whoaaa algo raro verdad, bueno como ahora es el doble del tamaño que nuestro arreglo original se encuentra al doble de distancia de la que deberia

  > [Uint8Contents]: <2d 00 00 00 00 00 41 00 00 00 00 00>,
                      0  1  2  3  4  5  6

en memoria se veria masomenos asi (quitale zoom al editor para mas claridad)

supongamos que la direccion de nuestro arreglo es 100

Ubicacion de nuestro primer elemento que pusimos en el arreglo de 8 bits

100 101 102 103 104 105 106 107

Ahora como estamos poniendo desde nuestro arreglo de 16 bits va a multiplicar nuestro indice por el tamaño (16) y a eso le sumara la direccion inicial para obtener el lugar en el que va a posicionar nuestro entero de 16 bits

100 + (16 * 3)
100 + 48
148

100 101 102 103 104 105 106 107 ...........................................................................................................................................................................................
148 149 150 151 152 153 154 155

Y que tiene que ver Big O con todo esto?

Pues bueno, ya que sabemos como funciona la locacion de los elementos de un arreglo en memoria ver un elemento en el index 10000 por ejemplo algo que vaya a aumentar nuestra O? No, porque la manera en la que se lee un elemento especifico del arreglo o se borra un elemento de un arreglo no lo hace checando uno por uno, simplemente hace el calculo con la formula que ya vimos y ya

Esto se conoce como constant time, que es hacer un numero constante de cosas, en este caso hacemos las mismas operaciones si pasamos como indice 0 o 10020301200023 es la misma cantidad de pasos

User
> const a = new ArrayBuffer(12)
undefined
> a
ArrayBuffer {
  [Uint8Contents]: <00 00 00 00 00 00 00 00 00 00 00 00>,
  byteLength: 12
}
> const a8 = new Uint8Array(a)
undefined
> a8
Uint8Array(12) [
  0, 0, 0, 0, 0,
  0, 0, 0, 0, 0,
  0, 0
]
> a8[0] = 45
45
> a8
Uint8Array(12) [
  45, 0, 0, 0, 0,
   0, 0, 0, 0, 0,
   0, 0
]
> a
ArrayBuffer {
  [Uint8Contents]: <2d 00 00 00 00 00 00 00 00 00 00 00>,
  byteLength: 12
}
> const a16 = new Uint16Array(a)
undefined
> a
ArrayBuffer {
  [Uint8Contents]: <2d 00 00 00 00 00 00 00 00 00 00 00>,
  byteLength: 12
}
> a16[3]=65
65
> a
ArrayBuffer {
  [Uint8Contents]: <2d 00 00 00 00 00 41 00 00 00 00 00>,
  byteLength: 12
}
