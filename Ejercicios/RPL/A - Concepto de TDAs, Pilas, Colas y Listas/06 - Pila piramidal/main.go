// Dada una pila de enteros, escribir una función que determine si es piramidal.
// Una pila de enteros es piramidal si cada elemento es menor a su elemento inferior
// (en el sentido que va desde el tope de la pila hacia el otro extremo).
// La pila no debe ser modificada.

package main

func EsPiramidal(pila Pila[int]) bool {
	piramidal := true
	pilaAuxiliar := CrearPilaDinamica[int]()
	for !pila.EstaVacia() {
		desapilado := pila.Desapilar()
		pilaAuxiliar.Apilar(desapilado)
		if !pila.Estavacia() && (pila.VerTope() < desapilado) {
			piramidal = false
			break
		}
	}
	for !pilaAuxiliar.EstaVacia() {
		pila.Apilar(pilaAuxiliar.Desapilar())
	}
	return piramidal
}
