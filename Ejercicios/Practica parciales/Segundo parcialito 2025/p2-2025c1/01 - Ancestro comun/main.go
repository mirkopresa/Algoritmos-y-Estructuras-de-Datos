// Implementar una primitiva para el ABB func (arbol *abb[K, V]) AncestroComun(clave1, clave2 K) K que reciba
// 2 claves y devuelva el último ancestro en común entre ambas claves. Dicho ancestro en común podría ser incluso alguna
// de estas claves. Si alguna clave pasada no se encuentra en el árbol, finalizar con panic. Indicar y justificar la complejidad
// de la primitiva implementada.

package main

func (arbol *abb[K, V]) AncestroComun(clave1, clave2 K) K {
	if !arbol.Pertenece(clave1) || !arbol.Pertenece(clave2) {
		panic("Una o ambas claves no se encuentran en el arbol")
	}
	return arbol._ancestroComun(arbol.raiz, clave1, clave2)
}

func (ab *abb[K, V]) _ancestroComun(actual *nodoAbb[K, V], clave1, clave2 K) K {
	comparacion1 := ab.cmp(clave1, actual.clave)
	comparacion2 := ab.cmp(clave2, actual.clave)
	if comparacion1 < 0 && comparacion2 < 0 {
		return ab._ancestroComun(actual.izq, clave1, clave2)
	} else if comparacion1 > 0 && comparacion2 > 0 {
		return ab._ancestroComun(actual.der, clave1, clave2)
	} else {
		return actual.dato
	}
}
