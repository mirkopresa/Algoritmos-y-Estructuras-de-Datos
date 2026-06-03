// Implementar una primitiva Diferencia para el heap, que reciba otro Heap y cree un nuevo Heap con los elementos del primero que
// no se encuentren en el segundo. La función de comparación del nuevo heap debe ser la del primer heap. Los heaps trabajan con
// mismos tipos de datos. Indicar y justificar la complejidad de la función implementada.

package main

func (h1 *colaConPrioridad[T]) Diferencia(h2 *colaConPrioridad[T]) ColaPrioridad[T] {
	dict := CrearHash[T, bool]()
	for _, elem := range h2.datos {
		if !dict.Pertenece(elem) {
			dict.Guardar(elem, true)
		}
	}
	nuevoDatos := make([]T, 0)
	for _, elem := range h1.datos {
		if !dict.Pertenece(elem) {
			nuevoDatos = append(nuevoDatos, elem)
		}
	}
	nuevoHeap := CrearHeapArr[T](nuevoDatos, h1.cmp)
	return nuevoHeap
}
