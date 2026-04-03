// Implementar una función que ordene de manera ascendente una pila de enteros
// sin conocer su estructura interna y utilizando como estructura auxiliar sólo otra pila auxiliar.
// Por ejemplo, la pila [ 4, 1, 5, 2, 3 ] debe quedar como [ 1, 2, 3, 4, 5 ]
// (siendo el último elemento el tope de la pila, en ambos casos).
// Indicar y justificar el orden de la función.

package main

// O(n²)

func Ordenar(pila Pila[int]) {
	pilaAuxiliar := CrearPilaDinamica[int]()
	for !pila.EstaVacia() {
		desapilado := pila.Desapilar()
		for !pilaAuxiliar.EstaVacia() && pilaAuxiliar.VerTope() < desapilado {
			pila.Apilar(pilaAuxiliar.Desapilar())
		}
		pilaAuxiliar.Apilar(desapilado)
	}
	for !pilaAuxiliar.EstaVacia() {
		pila.Apilar(pilaAuxiliar.Desapilar())
	}
}
