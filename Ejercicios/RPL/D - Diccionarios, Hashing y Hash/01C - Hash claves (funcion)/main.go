// Implementar una función func Claves[K comparable, V any](Diccionario[K, V]) Lista[K]
// que reciba un diccionario y devuelva una lista con sus claves.

package main

func Claves[K comparable, V any](dict Diccionario[K, V]) Lista[K] {
	listaResultado := CrearListaEnlazada[K]()
	for iter := dict.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, _ := iter.VerActual()
		listaResultado.InsertarUltimo(clave)
	}
	return listaResultado
}
