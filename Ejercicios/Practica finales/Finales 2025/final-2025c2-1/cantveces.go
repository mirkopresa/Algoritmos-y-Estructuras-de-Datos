// Implementar un algoritmo de determine la cantidad de veces que aparece un elemento en un arreglo ordenado, en O(log n). Justificar
// adecuadamente la complejidad del algoritmo implementado. Si te resulta muy difícil de resolver, podés considerar que el elemento
// aparece a lo sumo O(log n) veces (en caso de usar esta presunción, el ejercicio estará para a lo sumo B=). Explicar dónde/cómo se
// usaría esto, en caso de haberlo hecho.

package main

func CantVeces(arr []int, k int) int {
	primeraPos := _cantVecesPrimero(arr, 0, len(arr)-1, k)
	if primeraPos == -1 {
		return -1
	}
	ultimaPos := _cantVecesUltimo(arr, 0, len(arr)-1, k)
	return ultimaPos - primeraPos + 1
}

func _cantVecesPrimero(arr []int, inicio, fin, k int) int {
	if inicio > fin {
		return -1
	}
	if inicio == fin {
		return inicio
	}
	mitad := (inicio + fin) / 2
	if arr[mitad] == k {
		if mitad == 0 || arr[mitad-1] != k {
			return mitad
		} else {
			return _cantVecesPrimero(arr, inicio, mitad-1, k)
		}
	} else if arr[mitad] > k {
		return _cantVecesPrimero(arr, inicio, mitad-1, k)
	} else {
		return _cantVecesPrimero(arr, mitad+1, fin, k)
	}
}

func _cantVecesUltimo(arr []int, inicio, fin, k int) int {
	if inicio == fin {
		return inicio
	}
	mitad := (inicio + fin) / 2
	if arr[mitad] == k && arr[mitad+1] != k {
		return mitad
	} else if arr[mitad] > k {
		return _cantVecesUltimo(arr, inicio, mitad-1, k)
	} else {
		return _cantVecesUltimo(arr, mitad+1, fin, k)
	}
}
