// Implementar una primitiva para el Heap merge(otroHeap heap[T]) ColaPrioridad[T] que dado otro heap, nos devuelva un
// nuevo heap (con la función de comparación del heap al que se le invoca la primitiva) que contengan los elementos de ambos heaps,
// pero con un único elemento de misma prioridad (no importa cuál elemento se elige).
// Indicar y justificar la complejidad de la primitiva implementada.

package main

// basicamente merge heaps sin repetidos
func (h1 *colaPrioridad[T]) Merge(h2 colaPrioridad[T]) ColaPrioridad[T] {
	nuevoHeap := CrearHeap[T](h1.cmp)
	arreglo := make([]T, 0)
	arreglo = append(arreglo, h1.datos...)
	for !h2.EstaVacia() {
		arreglo = append(arreglo, h2.Desencolar())
	}
	arreglo = MergeSort(arreglo, h1.cmp)
	for i, elem := range arreglo {
		if i > 0 && h1.cmp(elem, arreglo[i-1]) == 0 {
			continue
		}
		nuevoHeap.Encolar(elem)
	}
	return nuevoHeap
}

func MergeSort[T any](arr []T, cmp func(T, T) int) []T {
	if len(arr) <= 1 {
		return arr
	}
	mitad := len(arr) / 2
	mitadIzq := MergeSort(arr[:mitad], cmp)
	mitadDer := MergeSort(arr[mitad:], cmp)
	return merge(mitadIzq, mitadDer, cmp)
}

func merge[T any](mitadIzq, mitadDer []T, cmp func(T, T) int) []T {
	resultado := make([]T, 0, len(mitadIzq)+len(mitadDer))
	i, j := 0, 0
	for i < len(mitadIzq) && j < len(mitadDer) {
		if cmp(mitadIzq[i], mitadDer[j]) >= 0 {
			resultado = append(resultado, mitadDer[j])
			j++
		} else {
			resultado = append(resultado, mitadIzq[i])
			i++
		}
	}
	resultado = append(resultado, mitadIzq[i:]...)
	resultado = append(resultado, mitadDer[j:]...)
	return resultado
}
