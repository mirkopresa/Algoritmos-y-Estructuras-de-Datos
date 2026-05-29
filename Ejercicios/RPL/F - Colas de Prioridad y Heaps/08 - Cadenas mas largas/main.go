// Escribir una función en Go que, dado un arreglo de n cadenas y un entero
// positivo k, devuelva una lista con las k cadenas más largas. Se espera que el
// orden del algoritmo sea O(n + k log n). Justificar el orden.

package main

func cmpLongitud(c1, c2 string) int {
	l1 := len(c1)
	l2 := len(c2)
	if l1 < l2 {
		return -1
	} else if l1 > l2 {
		return 1
	} else {
		return 0
	}
}

func CadenasLargas(arr []string, k int) []string {
	heap := CrearHeapArr(arr, cmpLongitud)
	resultado := make([]string, 0, k)
	i := 0
	for i < k && !heap.EstaVacia() {
		resultado = append(resultado, heap.Desencolar())
		i++
	}
	return resultado
}
