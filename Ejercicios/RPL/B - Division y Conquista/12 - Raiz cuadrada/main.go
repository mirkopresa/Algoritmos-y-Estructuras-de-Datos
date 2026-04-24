// Implementar un algoritmo que, por división y conquista, permita obtener la parte entera de la
// raíz cuadrada de un número n, en tiempo O(log n). Por ejemplo, para n = 10 debe devolver 3,
// y para n = 25 debe devolver 5.

package main

func ParteEnteraRaiz(n int) int {
	return _parteEntera(n, 0, n)
}

func _parteEntera(n, inicio, fin int) int {
	if inicio > fin {
		return fin
	}
	mitad := (inicio + fin) / 2
	cuadrado := mitad * mitad
	if cuadrado == n {
		return mitad
	} else if cuadrado > n {
		return _parteEntera(n, inicio, mitad-1)
	} else {
		return _parteEntera(n, mitad+1, fin)
	}
}
