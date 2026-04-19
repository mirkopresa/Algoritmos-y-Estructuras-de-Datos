// Implementar una función func Balanceado(texto string) bool, que retorne si el texto
// esta balanceado o no. Un texto sólo puede contener los siguientes caracteres: [,],{,}(,).
// Indicar y justificar la complejidad de la función implementada.
// Un texto esta balanceado si cada agrupador abre y cierra en un orden correcto.

package main

const (
	PARENTESIS_APERTURA = "("
	PARENTESIS_CIERRE   = ")"
	LLAVE_APERTURA      = "{"
	LLAVE_CIERRE        = "}"
	CORCHETE_APERTURA   = "["
	CORCHETE_CIERRE     = "]"
)

func Balanceado(texto string) bool {
	pila := CrearPilaDinamica[string]()
	for _, c := range texto {
		caracter := string(c)
		if caracter == PARENTESIS_APERTURA || caracter == LLAVE_APERTURA || caracter == CORCHETE_APERTURA {
			pila.Apilar(caracter)
		} else {
			if pila.EstaVacia() {
				return false
			}
			tope := pila.VerTope()
			if caracter == PARENTESIS_CIERRE && tope != PARENTESIS_APERTURA {
				return false
			} else if caracter == LLAVE_CIERRE && tope != LLAVE_APERTURA {
				return false
			} else if caracter == CORCHETE_CIERRE && tope != CORCHETE_APERTURA {
				return false
			}
			pila.Desapilar()
		}
	}
	if pila.EstaVacia() {
		return true
	} else {
		return false
	}
}
