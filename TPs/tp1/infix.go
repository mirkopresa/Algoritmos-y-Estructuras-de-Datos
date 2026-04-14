package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	TDAPila "tdas/pila"
	"unicode"
)

// Funcion que sirve para devolver la procedencia de un operador
func Procedencia(operador string) int {
	switch operador {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	case "^":
		return 3
	case "(", ")":
		return 4
	}
	return 0 // Es un numero
}

// Funcion que devuelve la asociatividad de un operador, true para izquierda y false para derecha
func Asociatividad(operador string) bool {
	if operador == "^" {
		return false
	}
	return true
}

func numeroNoVacio(numero *strings.Builder, vectorInfix []string) []string {
	if numero.Len() > 0 {
		vectorInfix = append(vectorInfix, numero.String())
		numero.Reset()
		return vectorInfix
	}
	return vectorInfix
}

// Funcion que convierte una linea de texto en un array de strings,
// donde cada string es un numero o un operador
func TextoAVector(linea string) []string {
	vectorInfix := make([]string, 0)
	numero := strings.Builder{}
	for _, caracter := range linea {
		switch {
		case unicode.IsDigit(caracter): // Es un numero
			numero.WriteRune(caracter)
		case Procedencia(string(caracter)) >= 1 && Procedencia(string(caracter)) <= 4: // Es un operador
			vectorInfix = numeroNoVacio(&numero, vectorInfix)
			vectorInfix = append(vectorInfix, string(caracter))
		}
	}
	vectorInfix = numeroNoVacio(&numero, vectorInfix)
	return vectorInfix
}

func ConvertirInfixToPostfix(vectorInfix []string) []string {
	pila := TDAPila.CrearPilaDinamica[string]()
	vectorPostfix := make([]string, 0)
	for _, caracter := range vectorInfix {
		switch {
		case unicode.IsDigit(rune(caracter[0])): // Es un numero
			vectorPostfix = append(vectorPostfix, caracter)
		case caracter == "(":
			pila.Apilar(caracter)
		case caracter == ")":
			for !pila.EstaVacia() && pila.VerTope() != "(" {
				vectorPostfix = append(vectorPostfix, pila.Desapilar())
			}
			pila.Desapilar()
		default: // Es un operador
			for !pila.EstaVacia() && pila.VerTope() != "(" {
				operador := Procedencia(caracter)
				operadorTope := Procedencia(pila.VerTope())
				if operadorTope > operador || (operadorTope == operador && Asociatividad(caracter)) {
					vectorPostfix = append(vectorPostfix, pila.Desapilar())
				} else {
					break
				}
			}
			pila.Apilar(caracter)
		}
	}
	for !pila.EstaVacia() {
		vectorPostfix = append(vectorPostfix, pila.Desapilar())
	}
	return vectorPostfix
}

// Funcion que imprime el vectorPostfix por Stdout
func ImprimirResultado(vectorPostfix []string) {
	lineaPostfix := strings.Join(vectorPostfix, " ")
	fmt.Fprintf(os.Stdout, "%s\n", lineaPostfix)
}

func main() {
	lector := bufio.NewScanner(os.Stdin)
	for lector.Scan() {
		linea := lector.Text()
		vectorInfix := TextoAVector(linea)
		vectorPostfix := ConvertirInfixToPostfix(vectorInfix)
		ImprimirResultado(vectorPostfix)
	}
}
