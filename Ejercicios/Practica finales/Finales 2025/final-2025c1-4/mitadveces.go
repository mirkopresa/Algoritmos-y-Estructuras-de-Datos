// Implementar un algoritmo que dado un arreglo de números, determine si hay un elemento dentro del mismo que
// aparece al menos la mitad veces. La complejidad del algoritmo debe ser lineal. Justificar la complejidad del algoritmo
// implementado.

package main

func MitadVeces(arr []int) bool {
	dicc := CrearHash[int, int]()
	for _, elem := range arr {
		if !dicc.Pertenece(elem) {
			dicc.Guardar(elem, 1)
		} else {
			cantidad := dicc.Obtener(elem)
			cantidad++
			dicc.Guardar(cantidad)
			if cantidad >= len(arr)/2 {
				return true
			}
		}
	}
	return false
}
