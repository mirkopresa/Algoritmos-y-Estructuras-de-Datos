// Implementar una función que reciba una pila de enteros y devuelva la suma de todos los elementos.
// Al finalizar la ejecución de la función, la pila debe quedar en el mismo estado que tenía antes de ejecutar la misma.
// La función no puede utilizar estructuras auxiliares (incluyendo arreglos). La firma de la
// función debe ser pilaSumar(pila Pila[int]) int.
// Indicar y justificar la complejidad de la función implementada.

package main

// O(n)

func pilaSumar(pila Pila[int]) int {
	if pila.EstaVacia() {
		return 0
	}
	desapilado := pila.Desapilar()
	suma := desapilado + pilaSumar(pila)
	pila.Apilar(desapilado)
	return suma
}
