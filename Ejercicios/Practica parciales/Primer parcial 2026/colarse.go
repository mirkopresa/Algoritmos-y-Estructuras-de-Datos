// Debido a la alta demanda de figuritas por el inminente inicio del Mundial, el supermercado Lo de Diego está más
// concurrido que nunca. Para evitar el caos, Diego designó al gerente Hollander para supervisar las filas de las cajas. Él
// puede decidir mover un cliente al inicio de la fila si tiene pocos items.
// Implementar la primitiva Colarse para el TDA ColaEnlazada que, dado un dato, lo mueva al desde su lugar (asumir
// que hay solo una aparición de cada dato) hasta el inicio de la Cola, para que pueda ser usada por Hollander para mover
// clientes.
// Indicar y justificar la complejidad de la primitiva.

package main

import "reflect"

func (cola *colaEnlazada[T]) Colarse(dato T) {
	var anterior *nodo[T]
	anterior = nil
	actual := cola.primero
	for actual != nil { // O(n) siendo n los elementos de la cola, si el elemento no esta, la va a recorrer toda
		if reflect.DeepEqual(actual.dato, dato) {
			if anterior == nil {
				return
			}
			anterior.siguiente = actual.siguiente
			actual.siguiente = cola.primero
			cola.primero = actual
			if anterior.siguiente == nil {
				cola.ultimo = anterior
			}
			return
		}
		anterior = actual
		actual = actual.siguiente
	}
}
