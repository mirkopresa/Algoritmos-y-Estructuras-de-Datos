// Implementar un algoritmo que reciba un arreglo desordenado de enteros, su largo (n) y
// un número K y determinar en O(n) si existe un par de elementos (dos elementos) en el arreglo que sumen exactamente K.

package main

func ParElementosSumanK(arr []int, k int) bool {
	dict := CrearHash[int, bool]()
	for _, num := range arr {
		posibleSuma := k - num
		if dict.Pertenece(posibleSuma) {
			return true
		}
		dict.Guardar(num, true)
	}
	return false
}
