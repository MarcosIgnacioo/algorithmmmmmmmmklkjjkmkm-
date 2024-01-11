# Busqueda lineal y kata setup

La busqueda lineal se implementa leyendo todo un arreglo (en el peor de los casos). Esto es O(N) porque crece de manera lineal, por ejemplo, si tuvieramos un arreglo de 4 caracteres y estamos buscando: "E" pero este no se encuentra en el arreglo (significando asi que tenemos que recorrer todo el arreglo para darnos cuenta de eso) por lo que si al mismo algoritmo le ponemos un arreglo del doble de largo (8 caracteres) y tampoco tiene el que buscamos, el coste de realizar este algoritmo crece de manera lineal pues se nos dio el doble de elementos por los que iterar por lo que nos cuesta el doble de tiempo recorrer el arreglo

