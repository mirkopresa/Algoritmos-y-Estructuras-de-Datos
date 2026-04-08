// Realizar un seguimiento de ordenar el siguiente arreglo utilizando
// Inserción, Selección, MergeSort, QuickSort y HeapSort.
// Contar la cantidad de operaciones (aproximadamente) para validar empíricamente
// la diferencia en el orden de cada uno, y poder comparar aquellos que sean iguales:
// [1, 7, 5, 8, 2, 4, 9, 6, 5].

package main

import "fmt"

// O(n²)
func Inserción(array []int) ([]int, int) {
	contador := 0
	for i, valor := range array[1:] {
		j := i
		for j >= 0 && valor < array[j] {
			array[j+1] = array[j]
			contador++
			j--
		}
		array[j+1] = valor
		contador++
	}
	return array, contador
}

func Maximo(array []int) (int, int) {
	maximo := 0
	contador := 0
	for i := 1; i < len(array); i++ {
		if array[i] > array[maximo] {
			maximo = i
		}
		contador++
	}
	return maximo, contador
}

// O(n²)

func Selección(array []int) ([]int, int) {
	contador := 0
	for i := len(array) - 1; i > 0; i-- {
		maximo, operaciones := Maximo(array[:i+1])
		array[i], array[maximo] = array[maximo], array[i]
		contador += operaciones
		contador++
	}
	return array, contador
}

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
		if mitadIzq[i] >= mitadDer[j] {
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

// En el peor de los casos O(n²) y en el mejor O(n log n)
func QuickSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	pivote := len(array) / 2
	arrayMenores, arrayPivote, arrayMayores := Partition(array, array[pivote])
	arrayMenoresOrdenado := QuickSort(arrayMenores)
	arrayMayoresOrdenado := QuickSort(arrayMayores)
	arrayOrdenado := make([]int, 0)
	arrayOrdenado = append(arrayOrdenado, arrayMenoresOrdenado...)
	arrayOrdenado = append(arrayOrdenado, arrayPivote...)
	arrayOrdenado = append(arrayOrdenado, arrayMayoresOrdenado...)
	return arrayOrdenado
}

func Partition(array []int, pivote int) ([]int, []int, []int) {
	arrayMenores := make([]int, 0)
	arrayPivote := make([]int, 0)
	arrayMayores := make([]int, 0)
	for _, elem := range array {
		if elem > pivote {
			arrayMayores = append(arrayMayores, elem)
		} else if elem < pivote {
			arrayMenores = append(arrayMenores, elem)
		} else {
			arrayPivote = append(arrayPivote, elem)
		}
	}
	return arrayMenores, arrayPivote, arrayMayores
}

func HeapSort() {
	// en algun momento lo hare
}

func main() {
	arrAOrdenar := []int{1, 7, 5, 8, 2, 4, 9, 6, 5}
	fmt.Println(Inserción(arrAOrdenar))
	fmt.Println(Selección(arrAOrdenar))
	fmt.Println(MergeSort(arrAOrdenar))
	fmt.Println(QuickSort(arrAOrdenar))
}
