package main

import "fmt"

func promedios(arr []float64) []float64 {
	resultado := make([]float64, len(arr))

	suma := 0.0
	for i := 0; i < len(arr); i++ {
		suma += arr[i]
		resultado[i] = suma / (float64(i) + 1)
	}

	return resultado
}

func main() {
	arr := []float64{5, 3, 4, 8, 1}
	fmt.Println(promedios(arr))
}
