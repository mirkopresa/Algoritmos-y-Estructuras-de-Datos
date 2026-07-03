// Implementar una función map[T any, V any](Lista[K], func(K) V) Lista[V] que dada una lista original, cree una nueva lista
// con el resultado de aplicarle a cada elemento la función pasada por parámetro. Para que el ejercicio esté completamente bien, se
// espera que se implemente utilizando el iterador interno de la lista. Indicar y justificar la complejidad de la función.

package main

func Map[K any, V any](lista Lista[K], f func(K) V) Lista[V] {
	resultado := CrearListaEnlazada[V]()
	lista.Iterar(func(dato K) bool {
		resultado.InsertarUltimo(f(dato))
		return true
	})
	return resultado
}
