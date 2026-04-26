// La diferencia simétrica entre dos conjuntos A y B es un conjunto que
// contiene todos los elementos que se encuentran en A y no en B, y viceversa.

// Implementar una función DiferenciaSimetricaDict[K comparable, V any](d1, d2 Diccionario[K, V]) Diccionario[K, V]
// que devuelva un nuevo Diccionario (usar CrearHash) con la diferencia simétrica entre los dos recibidos por parámetro.
// La diferencia tiene que ser calculada teniendo en cuenta las claves, y los datos asociados a las
// claves deben ser los mismos que estaban en cada uno de los hashes originales.

package main

// O(n+m)
func DiferenciaSimetricaDict[K comparable, V any](d1, d2 Diccionario[K, V]) Diccionario[K, V] {
	resultado := CrearHash[K, V]()
	for iter := d1.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, valor := iter.VerActual()
		if !d2.Pertenece(clave) {
			resultado.Guardar(clave, valor)
		}
	}
	for iter := d2.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, valor := iter.VerActual()
		if !d1.Pertenece(clave) {
			resultado.Guardar(clave, valor)
		}
	}
	return resultado
}
