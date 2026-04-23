// Implementar una primitiva del TDA Cola que devuelva dos colas: una con los elementos de las posiciones pares, y otra
// con los elementos de las posiciones impares (el primero de la cola puede considerarse como elemento en la posición 0).
// La cola original debe quedar en el mismo estado inicial. Indicar y justificar la complejidad de la primitiva.
// La firma de la primitiva debe ser Dividir() (Cola[T], Cola[T]).

package main

// Complejidad O(n), siendo n la cantidad de elementos que recorremos en la cola
func (c *colaEnlazada[T]) Dividir() (Cola[T], Cola[T]) {
	indice := 0
	colaPares := CrearColaEnlazada[T]()
	colaImpares := CrearColaEnlazada[T]()
	actual := c.primero
	for actual != nil {
		dato := actual.dato
		if indice%2 == 0 {
			colaPares.Encolar(dato)
		} else {
			colaImpares.Encolar(dato)
		}
		indice++
		actual = actual.siguiente
	}
	return colaPares, colaImpares
}
