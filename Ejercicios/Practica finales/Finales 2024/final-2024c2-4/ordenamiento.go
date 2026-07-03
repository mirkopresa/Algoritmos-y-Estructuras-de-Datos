// Tenemos un arreglo de n elementos en el que cada elemento se encuentra a lo sumo k posiciones de la que le correspondería
// si el arreglo estuviera ordenado (2 ≤ k ≤ n). Implementar un algoritmo de ordenamiento que funcione en O(n log k).

package main

// k = 3, n = 6 -> [3, 4, 5, 0, 1, 2]
// k = 4, n = 6 -> [2, 3, 4, 5, 0, 1]

func cmp_min(a, b int) int {
	return b - a
}

// Complejidad: O(nlogk + klogk), y como n >= k, O(nlogk)
func Ordenar(arr []int, k int) []int {
	res := make([]int, 0)
	heap := CrearHeap[int](cmp_min)
	i := 0
	// O(klogk)
	for i <= k {
		heap.Encolar(arr[i])
		i++
	}
	// k + 1 elementos encolados, minimo asegurado
	// O((n-k)log k) ?
	for k+1 < len(arr) {
		res = append(res, heap.Desencolar())
		heap.Encolar(arr[k+1])
		k++
	}
	// O(k * log(k+1))
	for !heap.EstaVacia() {
		res = append(res, heap.Desencolar())
	}
	return res
}
