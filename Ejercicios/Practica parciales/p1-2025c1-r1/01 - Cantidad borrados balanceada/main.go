// Se tiene una cadena que contiene () y ningún otro caracter (considerar que un único caracter es de tipo rune). Un
// ejercicio típico es dada una cadena averiguar si está balanceada (es decir, todos los símbolos de apertura se cierran, y
// además respetan el orden en el que se abrieron. Ejemplos balanceados: "()()()", o "(())()". No balanceados: "(()",
// o ")(".
// Teniendo en cuenta esto, se tiene una cadena que se asegura que en caso de borrar algunos paréntesis la cadena será
// balanceada, se pide implementar una función func cantBorradosBalanceada(cadena string) int que dada una
// cadena de este tipo, devuelva la cantidad mínima de paréntesis que se deben borrar para que la cadena esté balanceada.
// Indicar y justificar la complejidad del algoritmo.

package main

func cantBorradosBalanceada(cadena string) int {
	pila := CrearPilaDinamica[rune]()
	contador := 0
	for _, c := range cadena {
		if c == '(' {
			pila.Apilar(c)
		} else {
			if pila.EstaVacia() {
				contador++
			} else {
				pila.Desapilar()
			}
		}
	}
	for !pila.EstaVacia() {
		pila.Desapilar()
		contador++
	}
	return contador
}
