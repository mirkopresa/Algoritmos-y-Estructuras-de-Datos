// Se tiene una larga lista de números de tres cifras abc que representan
// números en notación científica de la forma: a,b * 10^c.
// Por ejemplo 123 representaría el número 1,2 * 10^3.
// Implemente un algoritmo para ordenar los números según su valor en notación científica. ¿De qué orden es?
// Se puede asumir que a nunca será 0 salvo que el número sea efectivamente 0. Es decir, la notación es correcta.

// Orden de importancia (y en el que vamos a ordenar) b -> a -> c

package main

import (
	"fmt"
	"strconv"
)

func ObtenerPrimerDigito(numero string) int {
	num, _ := strconv.Atoi(numero)
	return num % 10
}

func ObtenerSegundoDigito(numero string) int {
	num, _ := strconv.Atoi(numero)
	return (num / 10) % 10
}

func ObtenerTercerDigito(numero string) int {
	num, _ := strconv.Atoi(numero)
	return num / 100
}

// Complejidad O(d * (n + k)) siendo d la cantidad de ordenamientos,
// n la cantidad de elementos a ordenar y k el rango de los dígitos (en este caso 10)
// d = 3, k = 10, entonces la complejidad es O(n)
func CountingSort(arr []string, rango int, f func(string) int) []string {
	frecuencias := make([]int, rango)
	for _, num := range arr {
		frecuencias[f(num)]++
	}

	sumasAcumuladas := make([]int, rango)
	for j := 1; j < len(frecuencias); j++ {
		sumasAcumuladas[j] = sumasAcumuladas[j-1] + frecuencias[j-1]
	}

	arrResultado := make([]string, len(arr))
	for _, num := range arr {
		posicion := sumasAcumuladas[f(num)]
		arrResultado[posicion] = num
		sumasAcumuladas[f(num)]++
	}
	return arrResultado
}

func main() {
	arrAOrdenar := []string{"123", "991", "214"}
	arrResultado1 := CountingSort(arrAOrdenar, 10, ObtenerSegundoDigito)
	fmt.Println(arrResultado1)
	arrResultado2 := CountingSort(arrResultado1, 10, ObtenerTercerDigito)
	fmt.Println(arrResultado2)
	arrResultado3 := CountingSort(arrResultado2, 10, ObtenerPrimerDigito)
	fmt.Println(arrResultado3)
}
