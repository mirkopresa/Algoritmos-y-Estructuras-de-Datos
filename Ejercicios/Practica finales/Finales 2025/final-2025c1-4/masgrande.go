// Implementar un algoritmo que dado un arreglo de dígitos (0-9) determine cuál es el número más grande que se puede
// formar con dichos dígitos. Indicar y justificar la complejidad del algoritmo implementado.

package main

import "fmt"

func CountingSort(arr []int, k int) []int {
	frecuencias := make([]int, k)
	for _, elem := range arr {
		frecuencias[elem]++
	}
	sumasAcumuladas := make([]int, k)
	for i := 1; i < k; i++ {
		sumasAcumuladas[i] = sumasAcumuladas[i-1] + frecuencias[i-1]
	}
	resultado := make([]int, len(arr))
	for _, elem := range arr {
		resultado[sumasAcumuladas[elem]] = elem
		sumasAcumuladas[elem]++
	}
	return resultado
}

func NumeroMasGrande(arr []int) int {
	ordenado := CountingSort(arr, 10)
	numero := 0
	for i := len(ordenado) - 1; i >= 0; i-- {
		numero = (numero * 10) + ordenado[i]
	}
	return numero
}

func main() {
	fmt.Println(NumeroMasGrande([]int{8, 3, 6, 1, 9, 9, 4}))
}
