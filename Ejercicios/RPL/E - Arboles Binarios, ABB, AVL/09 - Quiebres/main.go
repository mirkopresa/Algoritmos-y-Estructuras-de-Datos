// Definimos como quiebre en un árbol binario cuando ocurre que:

//     un hijo derecho tiene un solo hijo, y es el izquierdo
//     un hijo izquierdo tiene un solo hijo, y es el derecho

// Implementar una primitiva para el árbol binario func (arbol Arbol) Quiebres() int que,
// dado un árbol binario, nos devuelva la cantidad de quiebres que tiene.
// La primitiva no debe modificar el árbol.

package main

type ab struct {
	izq  *ab
	der  *ab
	dato int
}

func (arbol *ab) Quiebres() int {
	if arbol == nil {
		return 0
	}
	quiebres := 0
	if arbol.der != nil {
		if arbol.der.der == nil && arbol.der.izq != nil {
			quiebres++
		}
	}
	if arbol.izq != nil {
		if arbol.izq.izq == nil && arbol.izq.der != nil {
			quiebres++
		}
	}
	return arbol.izq.Quiebres() + arbol.der.Quiebres() + quiebres
}
