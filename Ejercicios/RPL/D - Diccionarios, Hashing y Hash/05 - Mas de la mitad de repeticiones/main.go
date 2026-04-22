// Implementar una función de orden O(n) que dado un arreglo de n números enteros
// devuelva true o false según si existe algún elemento que aparezca más de la mitad de las veces.
// Justificar el orden de la solución.
// Ejemplos:
// [1, 2, 1, 2, 3] -> false
// [1, 1, 2, 3] -> false
// [1, 2, 3, 1, 1, 1] -> true
// [1] -> true

package main

func MasDeLaMitad(arr []int) bool {
	dict := CrearHash[int, int]()
	for _, elem := range arr {
		if dict.Pertenece(elem) {
			cantidad := dict.Obtener(elem)
			dict.Guardar(elem, cantidad+1)
		} else {
			dict.Guardar(elem, 1)
		}
		if dict.Obtener(elem) > len(arr)/2 {
			return true
		}
	}
	return false
}
