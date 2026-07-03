// Implementar la primitiva Invertir() para el Heap, que invierta su forma de comparar los elementos (es decir, si era de máximos,
// ahora sea de mínimos), sin modificar las funciones y primitivas previamente implementadas (simplemente contener todo el cambio
// dentro de la primitiva). Indicar y justificar la complejidad del algoritmo. Indicar qué consecuencias podría tener esta forma de
// implementación si se invoca a la primitiva Invertir una cantidad k de veces, y cómo podría resolverse si se permitiera modificar
// otras funciones y/o primitivas (y/o la estructura del heap en sí).

package main

// Esto generaria un problema de que se irian anidando funciones una y otra vez, haciendo que el resto de funciones cuesten
// O(k * (complejidad de la funcion)) y que si se acumulan muchas, puede dar un stack overflow

// Para solucionarlo, el struct podria tener una variable booleana invertido y el TDA podria tener una funcion "comparar"
// aparte, que si invertido es false, devuelva el resultado de heap.cmp como esta, y si no, lo devuelva negativo
// y que el resto de funciones en vez de usar heap.cmp, usen "comparar"
func (heap *colaPrioridad[T]) Invertir() {
	cmpOriginal := heap.cmp
	heap.cmp = func(a, b T) int {
		res := cmpOriginal(a, b)
		return -res
	}
	for i := len(heap.datos) - 1; i >= 0; i-- {
		heap.downheap(i)
	}
}
