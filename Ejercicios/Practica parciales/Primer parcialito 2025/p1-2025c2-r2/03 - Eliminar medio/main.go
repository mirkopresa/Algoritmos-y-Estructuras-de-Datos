// Se desea implementar una primitiva para la Cola Enlazada que permita “desencolar” el elemento
// del medio de la misma. Primero, implementar con la estructura de la cola tal cual se implementó
// durante la cursada. Indicar y justificar la complejidad de la primitiva. La firma debe ser func (cola
// *colaEnlazada[T]) EliminarMedio() T (en caso de no haber elementos, terminar con panic, símil
// a Desencolar). Una vez hecho esto, en caso que no hayas implementado la primitiva en tiempo
// constante, explicar cómo deberías hacer para poder implementarla en tiempo constante si se pudiera
// modificar la estructura la cola (y cómo deberían modificarse las demás primitivas para que el estado
// se mantenga correctamente).

package main

// O(n) + O(n/2) -> O(n)
// Se me ocurre que la cola de por si tenga un contador de elementos,
// y que aparte de tener un primero, y un ultimo, tengamos un puntero a la mitad
// Para esto, la cola tambien deberia calcular cada vez que desencolamos y encolamos la nueva mitad, y reasignar el puntero
// Tal vez (visto con IA) seria mejor tener un puntero al anterior a la mitad, para que reasignar los nodos sea mas facil
// y sea O(1)
func (cola *colaEnlazada[T]) EliminarMedio() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	cantidad := 0
	actual := cola.primero
	for actual != nil {
		cantidad++
		actual = actual.siguiente
	}
	mitad := cantidad / 2
	var anterior *nodo[T]
	actual = cola.primero
	for mitad > 0 {
		anterior = actual
		actual = actual.siguiente
		mitad--
	}
	if anterior == nil {
		cola.primero = actual.siguiente
	} else {
		anterior.siguiente = actual.siguiente
	}
	if actual.siguiente == nil {
		cola.ultimo == anterior
	}
	return actual.dato
}
