// Implementar la función func Multiprimeros[T any](cola Cola[T], k int) []T
// con el mismo comportamiento de la primitiva anterior:

// Dada una cola y un número k, devuelva los primeros k elementos de la cola,
// en el mismo orden en el que habrían salido de la cola. En caso que la cola tenga menos de k elementos.
// Si hay menos elementos que k en la cola, devolver un slice del tamaño de la cola.
// Indicar y justificar el orden de ejecución del algoritmo.

package main

// O(n)
func Multiprimeros[T any](cola Cola[T], k int) []T {
	elementos := make([]T, 0, k)
	colaAuxiliar := CrearColaEnlazada[T]()
	i := 0
	for !cola.EstaVacia() {
		desencolado := cola.Desencolar()
		if i < k {
			elementos = append(elementos, desencolado)
		}
		colaAuxiliar.Encolar(desencolado)
		i++
	}
	for !colaAuxiliar.EstaVacia() {
		cola.Encolar(colaAuxiliar.Desencolar())
	}
	return elementos
}
