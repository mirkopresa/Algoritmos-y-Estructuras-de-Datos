// Determinar cómo es el Árbol cuyo pre order es EURMAONDVSZT, e in order es MRAUOZSVDNET.

package main

type ab struct {
	izq   *ab
	der   *ab
	clave string
}

func ReconstruirParticular() *ab {
	preorder := []string{"E", "U", "R", "M", "A", "O", "N", "D", "V", "S", "Z", "T"}
	inorder := []string{"M", "R", "A", "U", "O", "Z", "S", "V", "D", "N", "E", "T"}
	return Reconstruir(preorder, inorder)
}

func Reconstruir(preorder, inorder []string) *ab {
	if len(preorder) == 0 {
		return nil
	}
	arbol := &ab{izq: nil, der: nil, clave: preorder[0]}
	var posRaiz int
	for i, elem := range inorder {
		if elem == preorder[0] {
			posRaiz = i
			break
		}
	}
	arbol.izq = Reconstruir(preorder[1:posRaiz+1], inorder[:posRaiz])
	arbol.der = Reconstruir(preorder[posRaiz+1:], inorder[posRaiz+1:])
	return arbol
}
