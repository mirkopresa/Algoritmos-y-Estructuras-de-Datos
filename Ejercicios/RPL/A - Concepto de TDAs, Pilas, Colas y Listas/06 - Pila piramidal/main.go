// Dada una pila de enteros, escribir una función que determine si es piramidal.
// Una pila de enteros es piramidal si cada elemento es menor a su elemento inferior
// (en el sentido que va desde el tope de la pila hacia el otro extremo).
// La pila no debe ser modificada.

package main

func EsPiramidal(pila Pila[int]) bool {
	piramidal := true
	arr := make([]int, 0)
	i := 0
	for !pila.EstaVacia() {
		arr[i] = pila.Desapilar()
		i++
	}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			piramidal = false
			break
		}
	}
	for _, valor := range arr {
		pila.Apilar(valor)
	}
	return piramidal
}
