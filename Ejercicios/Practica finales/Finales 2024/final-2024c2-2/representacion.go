// Implementar en Go una primitiva que reciba un árbol binario que representa un heap (árbol binario izquierdista, que
// cumple la propiedad de heap), y devuelva la representación en arreglo del heap. La firma de la primitiva debe ser
// RepresentacionArreglo() []T. Indicar y justificar la complejidad de la primitiva.

package main

type ab[T any] struct {
	izquierda *ab[T]
	derecha   *ab[T]
	dato      T
}

// O(n), siendo n la cantidad de elementos del heap (todos entran y salen de la cola)
func (ab *ab[T]) RepresentacionArreglo() []T {
	cola := CrearColaEnlazada[*ab[T]]()
	res := make([]T, 0)
	cola.Encolar(ab)
	for !cola.EstaVacia() {
		nodo := cola.Desencolar()
		if nodo == nil {
			continue
		}
		res = append(res, nodo.dato)
		cola.Encolar(nodo.izquierda)
		cola.Encolar(nodo.derecha)
	}
}
