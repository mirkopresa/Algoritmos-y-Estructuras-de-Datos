// Implementar una función que reciba un arreglo ordenado y devuelva una lista con los elementos en orden para ser insertados en un ABB,
// de tal forma que al insertarlos en dicho orden se asegure que el ABB quede balanceado.
// ¿Cómo cambiarías tu resolución si en vez de querer guardarlos en un ABB se fueran a insertar en un AVL?

package main

func balancear[K comparable](claves []K) Lista[K] {
	resultado := CrearListaEnlazada[K]()
	_balancear(claves, resultado)
	return resultado
}

func _balancear[K comparable](claves []K, res Lista[K]) {
	length := len(claves)
	if length == 0 {
		return
	}
	mitad := length / 2
	res.InsertarUltimo(claves[mitad])
	_balancear(claves[:mitad], res)
	_balancear(claves[mitad+1:], res)
}
