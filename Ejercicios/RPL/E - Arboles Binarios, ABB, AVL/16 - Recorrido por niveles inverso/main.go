// Implementar en Go una primitiva de ABB (DiccionarioOrdenado) que funcione como un iterador interno que haga un recorrido por niveles inverso.
// Es decir, que visite los elementos del nivel más inferior hasta la raiz.
// Para el ABB cuyo preorder es 5, 2, 1, 3, 4, 7, 9 (comparación numérica habitual), el recorrido debe ser: 4, 9, 3, 1, 7, 2, 5.
// En el 4to nivel está sólo el 4. En el 3er nivel están el 1, 3 y 9 pero hay que leerlos de derecha a izquierda.
// Luego en el 2do el 2 y 7 con la misma lógica, y finalmente la raíz 5 al final.
// Indicar y justificar la complejidad de la primitiva implementada.

package main

type abb[K comparable, V any] struct {
	izq   *abb[K, V]
	der   *abb[K, V]
	clave K
	dato  V
}

func (arbol *abb[K, V]) NivelesInverso(visitar func(clave K, dato V) bool) {
	// no logre entenderlo
}
