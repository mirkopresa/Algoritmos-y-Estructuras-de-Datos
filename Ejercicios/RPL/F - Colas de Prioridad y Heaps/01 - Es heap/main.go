// Implementar en lenguaje Go una función recursiva con la firma func esHeap(arr []int).
// Esta función debe devolver true o false de acuerdo a si el arreglo que recibe
// como parámetro cumple la propiedad de heap (de mínimos).

package main

func esHeap(arr []int) bool {
	return _esHeap(arr, 0)
}

func _esHeap(arr []int, i int) bool {
	if i >= len(arr) {
		return true
	}
	hijoIzq := 2*i + 1
	hijoDer := 2*i + 2
	if hijoIzq < len(arr) {
		if arr[i] > arr[hijoIzq] {
			return false
		}
		if hijoDer < len(arr) {
			if arr[i] > arr[hijoDer] {
				return false
			}
		}
	}
	esHeapIzq := _esHeap(arr, hijoIzq)
	esHeapDer := _esHeap(arr, hijoDer)
	return esHeapIzq && esHeapDer
}
