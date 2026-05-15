// Implementar una primitiva para el AB que reciba dos arreglos de claves.
// El primer arreglo corresponde al preorder de un árbol binario. El segundo al inorder del mismo árbol
// (ambos arreglos tienen los mismos elementos, sin repetidos).
// La función debe devolver un árbol binario que tenga dicho preorder e inorder.
// Indicar y justificar el orden de la primitiva (tener cuidado con este punto).

package main

type Arbol struct {
	izq   *Arbol
	der   *Arbol
	clave int
}

// preorder[1:i+1] - preorder[i+1:]
// inorder[:i] - inorder[i+1:]

// A = 2, B = 2, C = 1, log en base 2 de 2 -> 1 = C, O(n^c * log n) -> O(n log n)
// siendo n la cantidad de nodos del arbol (elementos del arreglo)
// pero en el peor de los casos, si un arbol tiene todos nodos izquierdos, o todos derechos
// no vamos a partir en mitades, y nos queda la sumatoria de gauss -> O(n²)
func Reconstruir(preorder, inorder []int) *Arbol {
	if len(preorder) == 0 {
		return nil
	}
	nuevoArbol := &Arbol{izq: nil, der: nil, clave: preorder[0]}
	var posicionRaiz int
	for i, elem := range inorder {
		if elem == preorder[0] {
			posicionRaiz = i
			break
		}
	}
	nuevoArbol.izq = Reconstruir(preorder[1:posicionRaiz+1], inorder[:posicionRaiz])
	nuevoArbol.der = Reconstruir(preorder[posicionRaiz+1:], inorder[posicionRaiz+1:])
	return nuevoArbol
}
