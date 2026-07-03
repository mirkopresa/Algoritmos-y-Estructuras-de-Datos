// Implementar un algoritmo que ordene un arreglo con puntos (valores (x, y)) que se encuentran dentro del cuadrado centrado en
// el origen, de lado 2, (es decir, ∥x∥ ≤ 1 ∧ ∥y∥ ≤ 1), sabiendo que la distribución de los puntos es uniforme en dicho dominio.
// El criterio para ordenar es primero por coordenada x y, en caso de igualdad, desempatar por coordenada y. En ambos casos, de
// menor a mayor. Tener en cuenta que los números pueden tener “infinitos” decimales. El algoritmo debe ejecutar en tiempo lineal a la
// cantidad de elementos del arreglo a ordenar. Justificar la complejidad del algoritmo propuesto.

package main

type Punto struct {
	x float64
	y float64
}

func OrdenarPuntos(arr []Punto) []Punto {
	// falopa ni idea
}
