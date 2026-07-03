// Edgar necesitaba hacer unos cambios en un TDA que ya había implementado Alan, con todo y pruebas (todas
// funcionando correctamente, como es de esperar de Alan). Esto le llevó a Edgar n días. Cada día realizó un commit en el
// sistema de control de versiones que utilizan. El problema es que no corrió las pruebas hasta el día n, y recién ahí notó
// que fallaban. Edgar implementó una función func todoOkElDía(n int) bool que recibe un número de día y devuelve
// true si estaba todo ok hasta el día n o false si ese día ya fallaban.
// Implementar una función func buscarDiaFalla(diasTotales int) int que devuelva el número de día en el que
// empezaron a fallar las pruebas. Indicar y justificar la complejidad del algoritmo implementado (la complejidad de
// todoOkElDia es O(n)).

package main

func buscarDiaFalla(diasTotales int) int {
	return diaFalla(1, diasTotales)
}

// Estamos aplicando una funcion O(n) log(n) veces, entonces - > O(n log n)
func diaFalla(min, max int) int {
	if min > max {
		return -1
	}
	mitad := (min + max) / 2
	if !todoOkElDia(mitad) {
		if mitad == 1 || todoOkElDia(mitad-1) {
			return mitad
		} else {
			return diaFalla(min, mitad-1)
		}
	} else {
		return diaFalla(mitad+1, max)
	}
}
