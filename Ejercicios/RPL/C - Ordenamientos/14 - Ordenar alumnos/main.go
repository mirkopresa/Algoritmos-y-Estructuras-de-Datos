// Implementar en Go un algoritmo de RadixSort para ordenar un arreglo de Alumnos (estructuras)
// en función de sus notas en parcialitos, de menor a mayor. Los alumnos tienen su nombre y
// las notas (numéricas, 0-10) de los 3 parcialitos (tenemos las notas finales).
// El arreglo debe quedar ordenado primero por el promedio de notas.
// No importan los decimales, nada más si tuvo “promedio 9”, “promedio 8”, etc.,
// es decir la parte entera del promedio. Luego, en caso de igualdad en este criterio,
// los alumnos deben quedar ordenados por la nota del parcialito 1, en caso de persistir la igualdad,
// la del parcialito 2, y finalmente por la del 3.
// En ningún caso se utiliza el nombre para desempatar.
// Suponer que cualquier algoritmo de ordenamiento auxiliar que se requiera ya se encuentra implementado.
// Sí justificar por qué se utiliza el o los algoritmos auxiliares utilizados, y no otros.
// Indicar y justificar la complejidad del algoritmo.
// El desarrollo de la complejidad debe tar completo, no se aceptan resultados parciales.

package main

import "fmt"

type Alumno struct {
	nombre     string
	p1, p2, p3 int
}

// El ordenamiento seria O(n), ya que la complejidad de RadixSort es de O(d * (Complejidad del ordenamiento auxiliar))
// y en este caso, d = 4 (4 veces ordenamos), y la complejidad del ordenamiento auxiliar es O(n+k)
// Siendo n la cantidad de alumnos a ordenar, y k el rango acotado, que en este caso es siempre 11
// Entonces O(4*(n+11)) = O(n)
// El ordenamiento auxiliar tiene que ser estable, porque si no lo es, estariamos perdiendo el orden original
// en el caso de que sean iguales las notas del tercer, segundo, primer parcialito y en el caso de promedios iguales

func CountingSort[T any](arr []T, rango int, f func(T) int) []T {
	frecuencias := make([]int, rango)
	for _, num := range arr {
		frecuencias[f(num)]++
	}

	sumasAcumuladas := make([]int, rango)
	for j := 1; j < len(frecuencias); j++ {
		sumasAcumuladas[j] = sumasAcumuladas[j-1] + frecuencias[j-1]
	}

	arrResultado := make([]T, len(arr))
	for _, num := range arr {
		posicion := sumasAcumuladas[f(num)]
		arrResultado[posicion] = num
		sumasAcumuladas[f(num)]++
	}
	return arrResultado
}

func main() {
	alumnosOrdenados := []Alumno{
		{nombre: "Juan", p1: 8, p2: 9, p3: 10},
		{nombre: "Maria", p1: 9, p2: 8, p3: 7},
		{nombre: "Pedro", p1: 7, p2: 6, p3: 5},
		{nombre: "Ana", p1: 10, p2: 9, p3: 8},
	}
	alumnosOrdenados = CountingSort(alumnosOrdenados, 11, func(a Alumno) int {
		return a.p3
	})
	alumnosOrdenados = CountingSort(alumnosOrdenados, 11, func(a Alumno) int {
		return a.p2
	})
	alumnosOrdenados = CountingSort(alumnosOrdenados, 11, func(a Alumno) int {
		return a.p1
	})
	alumnosOrdenados = CountingSort(alumnosOrdenados, 11, func(a Alumno) int {
		return (a.p3 + a.p2 + a.p1) / 3
	})
	fmt.Println(alumnosOrdenados)
}
