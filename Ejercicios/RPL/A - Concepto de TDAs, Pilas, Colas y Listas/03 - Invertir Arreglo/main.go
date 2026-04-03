// Implementar una función que reciba un arreglo genérico e invierta su orden,
// utilizando los TDAs vistos. Indicar y justificar el orden de ejecución.

package main

func InvertirArreglo[T any](arr []T) {
	pila := CrearPilaDinamica[T]()
	for _, valor := range arr {
		pila.Apilar(valor)
	}
	for i := range arr {
		arr[i] = pila.Desapilar()
	}
}
