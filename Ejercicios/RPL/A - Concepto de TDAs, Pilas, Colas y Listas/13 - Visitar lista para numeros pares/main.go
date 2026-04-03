// Sabiendo que la firma del iterador interno de la lista enlazada es: Iterar(visitar func(K) bool)

// Se tiene una lista en donde todos los elementos son punteros a números enteros.
// Implementar una función SumaPares que reciba una lista y, utilizando el iterador interno (no el externo),
// calcule la suma de todos los números pares.

package main

func SumaPares(lista Lista[*int]) int {
	suma := 0
	visitar := func(numero *int) bool {
		if numero != nil && *numero%2 == 0 {
			suma += *numero
		}
		return true
	}
	lista.Iterar(visitar)
	return suma
}
