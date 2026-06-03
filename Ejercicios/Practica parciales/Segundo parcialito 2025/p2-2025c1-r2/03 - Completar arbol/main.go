// Se tiene un árbol binario que representa la fase eliminatoria del mundial. En cada nodo guarda el nombre del país, así
// como la cantidad de goles que convirtió en dicha fase (incluyendo la tanda de penales, si fuera necesario). El padre
// del nodo debe si o si tener al hijo que ganó (tuvo mayor cantidad de goles). Implementar una primitiva para el árbol
// donde solamente están los nombres de los equipos en las hojas (no en los internos), y deje el árbol completado con los
// ganadores en cada fase. Se puede asumir que el árbol es o bien completo, o que al menos todos los nodos internos tienen
// exactamente 2 hijos. La cantidad de goles en la raíz no es relevante.

package main

type Arbol struct {
	pais  string
	goles int
	izq   *Arbol
	der   *Arbol
}

func (arbol *Arbol) CompletarArbol() {
	if arbol == nil {
		return
	}
	arbol.izq.CompletarArbol()
	arbol.der.CompletarArbol()
	if arbol.izq != nil && arbol.der != nil {
		if arbol.izq.goles > arbol.der.goles {
			arbol.pais = arbol.izq.pais
		} else {
			arbol.pais = arbol.der.pais
		}
	}
}
