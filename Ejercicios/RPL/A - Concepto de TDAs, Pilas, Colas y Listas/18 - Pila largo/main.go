// Implementar una función recursiva que reciba una pila y devuelva,
// sin utilizar estructuras auxiliares, la cantidad de elementos de la misma.
// Al terminar la ejecución de la función la pila debe quedar en el mismo estado al original.

package main

func Largo[T any](pila Pila[T]) int {
	if pila.EstaVacia() {
		return 0
	}
	desapilado := pila.Desapilar()
	largo := 1 + Largo(pila)
	pila.Apilar(desapilado)
	return largo
}
