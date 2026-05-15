// Escribir una primitiva con la firma func (arbol *Arbol) Invertir() que invierta el árbol binario
// pasado por parámetro, de manera tal que los hijos izquierdos de cada nodo se conviertan en hijos derechos.

package main

type ab struct {
	izq  *ab
	der  *ab
	dato int
}

func (arbol *ab) Invertir() {
	if arbol == nil {
		return
	}
	arbol.izq, arbol.der = arbol.der, arbol.izq
	arbol.izq.Invertir()
	arbol.der.Invertir()
}
