// Implementar para la cola enlazada la primitiva Consumir(accion func (T)) que aplique la función accion a todos
// los elementos de la cola. Al terminar la ejecución, la cola debe quedar vacía. Se espera que se implemente sin utilizar
// otras primitivas, para demostrar el conocimiento sobre estructuras enlazadas. Indicar y justificar la complejidad de la
// primitiva.

package main

func (c *colaEnlazada[T]) Consumir(accion func(T)) {
	actual := c.primero
	for actual != nil {
		accion(actual.dato)
		actual = actual.siguiente
	}
	c.primero = nil
	c.ultimo = nil
}
