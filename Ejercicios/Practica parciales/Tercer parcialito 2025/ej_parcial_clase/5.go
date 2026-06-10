// Ejercicio 5: Heap.Patova(...)
// Implementar una primitiva Filter(f func(T) bool) para el Heap, que deje al heap únicamente con los elementos para los
// que la función devuelva true. La primitiva debe funcionar en O(n), siendo n la cantidad de elementos inicialmente en el heap.
// Por supuesto, luego de aplicar la operación, el heap debe quedar en un estado válido para poder seguir operando.
// Justificar la complejidad de la primitiva implementada.

package ejparcialclase

func (heap *colaConPrioridad[T]) Filter(f func(T) bool) {
	nuevoArreglo := make([]T, 0)
	for _, elem := range heap.datos { // O(n) siendo n los elementos del heap, f una funcion O(1)
		if f(elem) {
			nuevoArreglo = append(nuevoArreglo, elem)
		}
	}
	heap.cantidad = len(nuevoArreglo)
	heap.datos = nuevoArreglo
	for i := heap.cantidad - 1; i >= 0; i-- { // O(n), hacer downheap a todos los elementos es O(log n) solo en la raiz, luego la sumatoria de todos los costos nos da O(n)
		heap.downheap(i)
	}

}
