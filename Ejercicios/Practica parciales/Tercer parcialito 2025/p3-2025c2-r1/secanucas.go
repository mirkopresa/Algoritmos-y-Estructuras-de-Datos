package main

type vertice struct {
	dist   int
	nombre string
}

func topk(arr []vertice, k int) []string {
	heap := CrearHeapArr(arr, func(a, b vertice) int { return a.dist - b.dist })
	res := make([]string, 0)
	for k > 0 && !heap.EstaVacia() {
		res = append(res, heap.Desencolar().nombre)
		k--
	}
	return res
}
