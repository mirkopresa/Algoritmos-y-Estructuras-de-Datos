// Implementar una función masGrandePosible(digitos []int) int que dado un arreglo de dígitos (0-9) determine
// cuál es el número más grande que se puede formar con dichos dígitos. Indicar y justificar la complejidad del algoritmo
// implementado.

package main

// MergeSort tiene complejidad O(n log n)
// A = 2, B = 2, C = 1; log en base 2 de 2 = 1 == C -> O(n^c log n) -> O(n log n)
// Complejidad total O(n log n) + O(n) -> O(n log n)
func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	mitad := len(array) / 2
	mitadIzq := MergeSort(array[:mitad])
	mitadDer := MergeSort(array[mitad:])
	return Merge(mitadIzq, mitadDer)
}

func Merge(mitadIzq, mitadDer []int) []int {
	i, j := 0, 0
	mergeado := make([]int, 0, len(mitadIzq)+len(mitadDer))

	for i < len(mitadIzq) && j < len(mitadDer) {
		if mitadIzq[i] <= mitadDer[j] {
			mergeado = append(mergeado, mitadDer[j])
			j++
		} else {
			mergeado = append(mergeado, mitadIzq[i])
			i++
		}
	}
	mergeado = append(mergeado, mitadIzq[i:]...)
	mergeado = append(mergeado, mitadDer[j:]...)
	return mergeado
}

func masGrandePosible(digitos []int) int {
	digitos = MergeSort(digitos)
	resultado := 0
	// O(n) siendo n la cantidad de digitos
	for _, digito := range digitos {
		resultado = resultado*10 + digito
	}
	return resultado
}

func main() {
	digitos := []int{3, 1, 4, 1, 5, 9}
	println(masGrandePosible(digitos))
}
