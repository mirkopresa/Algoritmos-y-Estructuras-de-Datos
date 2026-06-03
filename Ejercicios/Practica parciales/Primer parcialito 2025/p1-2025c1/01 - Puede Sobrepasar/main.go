// Sincero Piladibujo, un famoso corredor de autos, se está yendo a la costa a disfrutar sus vacaciones de Semana Santa.
// Resulta que viene manejando por la ruta y tiene una fila de autos por delante, bastante extensa, que quiere sobrepasar.
// Al ser una ruta doble mano, puede ser bastante peligroso intentar sobrepasar toda la fila y que venga un auto de
// frente, por lo cual nos pidió ayuda. Se sabe que la fila se representa como una Cola[Auto] y que se tiene una función
// ObtenerTiempoDeSobrepaso(auto Auto) int, que se ejecuta en O(1) y que brinda el tiempo que tardará Sincero en
// sobrepasar a un auto en particular.
// Implementar una función PuedeSobrepasar(filaAutos Cola[Auto], tiempoMaximoDeManiobra int) bool siendo
// tiempoMaximoDeManiobra el tiempo en el que, por el paso del auto de la mano contraria, genera que la maniobra sea
// imposible y demasiado peligrosa. Al finalizar la ejecución de la función, la cola debe quedar en el estado original que
// tenía antes de ser ejecutada. Indicar y justificar la complejidad del algoritmo.

package main

// O(n) + O(n) = O(n), importante saber que para resetear la cola la debo vaciar completamente y luego volver a llenar
func PuedeSobrepasar(filaAutos Cola[Auto], tiempoMaximoDeManiobra int) bool {
	suma := 0
	colaAux := CrearColaEnlazada[Auto]()
	for !filaAutos.EstaVacia() {
		autoDesencolado := filaAutos.Desencolar()
		tiempoManiobra := ObtenerTiempoDeSobrepaso(autoDesencolado)
		suma += tiempoManiobra
		colaAux.Encolar(autoDesencolado)
	}
	for !colaAux.EstaVacia() {
		filaAutos.Encolar(colaAux.Desencolar())
	}
	if suma > tiempoMaximoDeManiobra {
		return false
	}
	return true
}
