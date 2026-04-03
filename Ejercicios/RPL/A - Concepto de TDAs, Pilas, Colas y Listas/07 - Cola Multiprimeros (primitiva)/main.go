// Implementar la primitiva func (cola *colaEnlazada[T]) Multiprimeros(k int) []T
// que dada una cola y un número k, devuelva los primeros k elementos de la cola,
// en el mismo orden en el que habrían salido de la cola.
//
// En caso que la cola tenga menos de k elementos.
// Si hay menos elementos que k en la cola, devolver un slice del tamaño de la cola.
// Indicar y justificar el orden de ejecución del algoritmo.

package main

// O(k) si n > k y O(n) si k > n
func (c *colaEnlazada[T]) Multiprimeros(k int) []T {
	elementos := make([]T, 0, k)
	nodoActual := c.primero
	for i := 0; i < k; i++ {
		if nodoActual == nil {
			break
		}
		elementos = append(elementos, nodoActual.dato)
		nodoActual = nodoActual.siguiente
	}
	return elementos
}
