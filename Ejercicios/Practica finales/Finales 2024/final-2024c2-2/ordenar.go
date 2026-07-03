// Implementar un algoritmo que permita ordenar cronológicamente un arreglo de cadenas que representan horarios en
// formato HH:MM:SS. Indicar y justificar la complejidad del algoritmo implementado.

package main

import (
	"strconv"
)

func extraerSegundos(cadena string) int {
	segundos, _ := strconv.Atoi(cadena[6:])
	return segundos
}

func extraerMinutos(cadena string) int {
	minutos, _ := strconv.Atoi(cadena[3:5])
	return minutos
}

func extraerHoras(cadena string) int {
	horas, _ := strconv.Atoi(cadena[:2])
	return horas
}

func CountingSort(arr []string, k int, f func(string) int) []string {
	frecuencias := make([]int, k)
	for _, elem := range arr {
		frecuencias[f(elem)]++
	}
	sumasAcumuladas := make([]int, k)
	for i := 1; i < k; i++ {
		sumasAcumuladas[i] = sumasAcumuladas[i-1] + frecuencias[i-1]
	}
	res := make([]string, len(arr))
	for _, elem := range arr {
		res[sumasAcumuladas[f(elem)]] = elem
		sumasAcumuladas[f(elem)]++
	}
	return res
}

func OrdenarTiempos(arr []string) []string {
	arr = CountingSort(arr, 60, extraerSegundos)
	arr = CountingSort(arr, 60, extraerMinutos)
	arr = CountingSort(arr, 24, extraerHoras)
	return arr
}
