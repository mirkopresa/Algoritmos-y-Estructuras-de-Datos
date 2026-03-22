// Esta función implementada recibe un arreglo y devuelve uno nuevo,
// donde en cada posición tiene el promedio del original
// contando los valores desde la posición 0 hasta dicha posición (inclusive).

// Por ejemplo, en el arreglo [5, 3, 4, 8, 1] devuelve [5, 4, 4, 5, 4.2].

// El problema es que… demora mucho. Primero preguntate qué complejidad tiene el algoritmo.

// Implementala de tal forma que ejecute notoriamente más rápido (pensá qué complejidad tendría dicho algoritmo).

package main

import "fmt"

// Funcion lenta
func promedioslento(arr []float64) []float64 {
	resultado := make([]float64, len(arr))

	for i := 0; i < len(arr); i++ {
		suma := 0.0
		for j := 0; j <= i; j++ {
			suma += arr[j]
		}
		resultado[i] = suma / (float64(i) + 1)
	}

	return resultado
}

// Funcion rapida
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
