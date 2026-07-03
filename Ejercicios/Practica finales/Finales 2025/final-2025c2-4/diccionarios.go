// Implementar una función que reciba dos diccionarios y devuelva true si ambos diccionarios son iguales, o false en caso contrario.
// Los diccionarios son iguales en caso que tanto las claves que contienen sean las mismas, como los valores asociadas a ellas. Se recibe
// una función de comparación para los valores. Indicar y justificar la complejidad de la función implementada.

package main

// O(n)
func SonIguales[K comparable, V any](dicc1, dicc2 Diccionario[K, V], cmp func(a, b V) int) bool {
	// ojo chequear que tengan misma cantidad de elementos primero
	if dicc1.Cantidad() != dicc2.Cantidad() {
		return false
	}
	for iter := dicc1.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
		clave, dato := iter.VerActual()
		if dicc2.Pertenece(clave) {
			if cmp(dato, dicc2.Obtener(clave)) == 0 {
				continue
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
