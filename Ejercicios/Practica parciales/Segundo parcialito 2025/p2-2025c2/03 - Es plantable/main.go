// Implementar en Go una primitiva para Árbol Binario func (ab *Arbol[int]) ArbolEsPlantable() bool que determine si un
// árbol es plantable, o no. Para que lo sea, todo nodo debe cumplir: el dato del nodo debe ser mayor al dato de sus hijos (si los tiene),
// y además, el dato del nodo no puede superar la altura de dicho nodo. Implementar la primitiva en O(n)

package main

type Arbol struct {
	dato int
	izq  *Arbol
	der  *Arbol
}

func (ab *Arbol) ArbolEsPlantable() bool {
	_, plantable := ab._arbolEsPlantable()
	return plantable
}

func (ab *Arbol) _arbolEsPlantable() (int, bool) {
	if ab == nil {
		return 0, true
	}
	alturaIzq, plantableIzq := ab.izq._arbolEsPlantable()
	alturaDer, plantableDer := ab.der._arbolEsPlantable()
	alturaActual := max(alturaIzq, alturaDer) + 1
	if !plantableIzq || !plantableDer || ab.dato > alturaActual {
		return alturaActual, false
	}
	esMayor := true
	if ab.izq != nil {
		if ab.izq.dato >= ab.dato {
			esMayor = false
		}
	}
	if ab.der != nil {
		if ab.der.dato >= ab.dato {
			esMayor = false
		}
	}
	return alturaActual, plantableIzq && plantableDer && esMayor
}
