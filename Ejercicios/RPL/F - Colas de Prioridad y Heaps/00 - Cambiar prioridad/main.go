// Implementar en Go una primitiva para el heap (siendo este un max-heap) que reciba un heap y
// una función de comparación y lo reordene de manera tal que se se comporte como max-heap para
// la nueva función de comparación (se cambia la función de prioridad).
// El orden de dicha primitiva debe ser O(n).

package main

func (heap *heap[T]) CambiarPrioridad(nuevaPrioridad func(a, b T) int) {
	heap.cmp = nuevaPrioridad
	for i := heap.cantidad - 1; i >= 0; i-- {
		heap.downheap(i)
	}
}


func (heap *heap[T]) downheap(posicionElem int) {
	posicionHijoIzq := 2*posicionElem + 1
	if posicionHijoIzq >= heap.cantidad {
		return // no hay ni hijo izquierdo, ni derecho
	}
	hijoMax := posicionHijoIzq
	posicionHijoDer := 2*posicionElem + 2
	if posicionHijoDer < heap.cantidad {
		comparacionHijos := heap.cmp(heap.datos[posicionHijoIzq], heap.datos[posicionHijoDer])
		if comparacionHijos < 0 {
			hijoMax = posicionHijoDer
		}
	}
	comparacion := heap.cmp(heap.datos[hijoMax], heap.datos[posicionElem])
	if comparacion > 0 {
		swap(&heap.datos[hijoMax], &heap.datos[posicionElem])
		heap.downheap(hijoMax)
	}
}
