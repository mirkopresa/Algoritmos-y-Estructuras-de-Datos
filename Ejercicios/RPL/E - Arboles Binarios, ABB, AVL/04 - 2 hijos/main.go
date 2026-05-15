// Dado un árbol binario, escriba una primitiva recursiva que cuente la cantidad de nodos
// que tienen exactamente dos hijos directos.
// ¿Qué orden de complejidad tiene la función implementada?

package main

type ab struct {
	izq  *ab
	der  *ab
	dato int
}

func (arbol *ab) DosHijos() int {
	if arbol == nil {
		return 0
	}
	if arbol.izq != nil && arbol.der != nil {
		return 1
	}
	return arbol.izq.DosHijos() + arbol.der.DosHijos()
}
