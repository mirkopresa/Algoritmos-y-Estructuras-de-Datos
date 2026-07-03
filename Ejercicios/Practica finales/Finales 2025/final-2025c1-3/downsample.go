// Implementar para la lista enlazada la primitiva Downsample(k int). Esta debe eliminar todos los elementos que se encuentren en
// posiciones múltiplos de k (k > 1). La primera posición es la posición 0. Indicar y justificar la complejidad del algoritmo implementado.

package main

func (l *listaEnlazada[T]) Downsample(k int) {
	i := 0
	actual := l.primero
	var anterior *nodoLista[T]
	for actual != nil {
		if i%k == 0 {
			if actual == l.primero {
				l.primero = l.primero.siguiente
			} else if actual == l.ultimo {
				anterior.siguiente = nil
				l.ultimo = anterior
			} else {
				anterior.siguiente = actual.siguiente
			}
		} else {
			anterior = actual
		}
		actual = actual.siguiente
		i++
	}
}
