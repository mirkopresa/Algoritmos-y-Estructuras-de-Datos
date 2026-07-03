// Suponer que queremos hacer un ejercicio típico de árboles, de los sencillos, como calcular la cantidad de hojas de
// un árbol. En general, lo implementaríamos de forma recursiva, por ejemplo, con un preorder. ¿Cómo lo harías,
// siguiendo la misma lógica de recorrido, para realizarlo de forma iterativa? ¿Utilizarías alguna otra estructura de
// datos vista en la materia? ¿cómo quedaría la complejidad?

package main

type ab struct {
	izq  *ab
	der  *ab
	dato int
}

// Basicamente usariamos una pila para simular recursividad, en la que inicialmente estara la raiz del arbol
// Mientras la pila no este vacia, desapilamos el nodo actual y chequeamos que sea un nodo valido, y que no tenga hijos
// Si es asi, es hoja y sumamos 1, y si no, apilamos sus hijos

// Complejidad: O(n), siendo n la cantidad de nodos del arbol (entran todos los nodos a la pila)
func Altura(arbol ab) int {
	pila := CrearPilaDinamica[*ab]()
	pila.Apilar(arbol)
	contador := 0
	for !pila.EstaVacia() {
		nodo := pila.Desapilar()
		if nodo == nil {
			continue
		}
		if nodo.izq == nil && nodo.der == nil {
			contador++
		} else {
			// Esto va alreves a comparacion de la recursividad, porque al apilar der y luego izq, primero se procesara izq, y luego der
			pila.Apilar(nodo.der)
			pila.Apilar(nodo.izq)
		}
	}
	return contador
}
