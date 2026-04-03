// Implementar en Go una primitiva func (lista *listaEnlazada[T]) Invertir()
// que invierta la lista, sin utilizar estructuras auxiliares.
// Indicar y justificar el orden de la primitiva.

package main

// O(n)
func (lista *ListaEnlazada[T]) Invertir() {
	actual := lista.primero
	var anterior *nodoLista[T] = nil
	for actual != nil {
		auxiliarSiguiente := actual.siguiente
		actual.siguiente = anterior
		anterior = actual
		actual = auxiliarSiguiente
	}
	lista.primero, lista.ultimo = lista.ultimo, lista.primero
}
