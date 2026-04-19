// Es el año 1700, y la pirata Barba-ra Verde atacó un barco de la Royal British Shipping & Something,
// que transportaba una importante piedra preciosa de la corona británica.
// Al parecer, la escondieron en un cofre con muchas piedras preciosas falsas, en caso de un ataque.
// Barba-ra Verde sabe que los refuerzos británicos no tardarán en llegar, y deben huir lo más rápido posible.
// El problema es que no pueden llevarse el cofre completo por pesar demasiado.
// Necesita encontrar rápidamente la joya verdadera. La única forma de descubrir la joya verdadera es pesando.
// Se sabe que la joya verdadera va a pesar más que las imitaciones, y que las imitaciones pesan todas lo mismo.
// Cuenta con una balanza de platillos para poder pesarlas (es el 1700, no esperen una balanza digital).

// En el ejemplo de código inicial de la actividad mostramos un llamado de ejemplo a la función balanza,
// a la que se le deben pasar los dos conjuntos de joyas a verificar. La cantidad de joyas en cada conjunto debe ser la misma,
// para que el resultado de la balanza de platillos nos dé información.

// Si los dos platillos pesan lo mismo, balanza devuelve 0.
// Si el primer platillo es más pesado, balanza devuelve 1.
// Si el segundo platillo es más pesado, balanza devuelve -1.

package main

[ - - - _ - -]

func encontrarJoyaRecursivo(arr []int, inicio, fin int) int {
	if fin - inicio == 1 {
		return 1
	}
	mitad := (inicio + fin) / 2
	var resultado, mitadDer int
	// Caso par, no excluimos a la mitad, caso impar, excluimos a la mitad a la hora de pesar
	if (fin - inicio) % 2 == 0 {
		resultado = balanza(arr[inicio:mitad], arr[mitad:fin])
		mitadDer = mitad
	} else {
		resultado = balanza(arr[inicio:mitad], arr[mitad+1:fin])
		mitadDer = mitad + 1
	}
	if resultado == 0 {
		return mitad
	} else if resultado == 1 {
		return encontrarJoyaRecursivo(arr, inicio, mitad)
	} else {
		return encontrarJoyaRecursivo(arr, mitadDer, fin)
	}
}

func encontrarJoya(joyas []int) int {
	return encontrarJoyaRecursivo(joyas, 0, len(joyas))
}
