// Implementar una función eliminarRepetidos(arreglo []int) []int que dado un arreglo de números, nos devuelva
// otro en el que estén los elementos del original sin repetidos. La primera aparición debe mantenerse, y las demás no ser
// consideradas. Indicar y justificar la complejidad del algoritmo implementado.

package main

// O(n) siendo N la cantidad de elementos del arreglo, y el resto de operaciones O(1)
func eliminarRepetidos(arreglo []int) []int {
	dictRepetidos := CrearHash[int, bool]()
	resultado := make([]int, 0)
	for _, elem := range arreglo {
		if !dictRepetidos.Pertenece(elem) {
			dictRepetidos.Guardar(elem, false)
		} else {
			dictRepetidos.Guardar(elem, true)
		}
		if dictRepetidos.Obtener(elem) == false {
			resultado = append(resultado, elem)
		}
	}
	return resultado
}

// simplificado
// O(n) siendo N la cantidad de elementos del arreglo, y el resto de operaciones O(1)
func eliminarRepetidos2(arreglo []int) []int {
	dictRepetidos := CrearHash[int, bool]()
	resultado := make([]int, 0)
	for _, elem := range arreglo {
		if !dictRepetidos.Pertenece(elem) {
			dictRepetidos.Guardar(elem, true)
			resultado = append(resultado, elem)
		}
	}
	return resultado
}
