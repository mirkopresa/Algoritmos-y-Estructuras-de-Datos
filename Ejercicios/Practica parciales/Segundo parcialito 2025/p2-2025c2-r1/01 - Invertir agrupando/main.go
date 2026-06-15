// Implementar la función InvertirAgrupando con la siguiente firma: func InvertirAgrupando(dicc Diccionario[K, V],
// cmpValores func(V, V) bool) Diccionario[V, Lista[K]]
// La función debe devolver un nuevo diccionario, cumpliendo las siguientes condiciones:
// (i) Las claves del nuevo diccionario serán los valores V del diccionario original.
// (ii) El valor para cada clave será una lista de las claves K originales que estaban asociadas a ese valor V (no importa el orden
// en el que queden).
// Ejemplo: Si el diccionario original es: {'Ana': 10, 'Beto': 20, 'Carla': 10, 'Dani': 20, 'Facu': 40}
// El nuevo diccionario tendría las siguientes claves asociadas a las correspondientes listas: {10: ['Ana', 'Carla'], 20: ['Beto',
// 'Dani'], 40: ['Facu']}
// Indicar y justificar la complejidad temporal del algoritmo implementado.

package main

func InvertirAgrupando[K comparable, V any](dicc Diccionario[K, V]) Diccionario[V, Lista[K]] {
	res := CrearHash[V, Lista[K]]()
	for iter := dicc.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
		clave, dato := iter.VerActual()
		if !res.Pertenece(dato) {
			lista := CrearListaEnlazada[K]()
			lista.InsertarUltimo(clave)
			res.Guardar(dato, lista)
		} else {
			lista := res.Obtener(dato)
			lista.InsertarUltimo(clave)
		}
	}
	return res
}
