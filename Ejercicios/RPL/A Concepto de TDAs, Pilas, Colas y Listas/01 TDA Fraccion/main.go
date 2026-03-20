package main

import "fmt"

type Fraccion struct {
	numerador   int
	denominador int
}

// CrearFraccion crea una fraccion con el numerador y denominador indicados.
// Si el denominador es 0, entra en panico.
func CrearFraccion(numerador, denominador int) Fraccion {
	if denominador == 0 {
		panic("Division por cero!")
	} else if denominador < 0 {
		return Fraccion{numerador: -numerador, denominador: valorAbsoluto(denominador)}
	}
	return Fraccion{numerador: numerador, denominador: denominador}
}

// Sumar crea una nueva fraccion, con el resultante de hacer la suma de las fracciones originales
func (f Fraccion) Sumar(otra Fraccion) Fraccion {
	numerador := (f.numerador * otra.denominador) + (otra.numerador * f.denominador)
	denominador := f.denominador * otra.denominador
	return CrearFraccion(numerador, denominador)
}

// Multiplicar crea una nueva fraccion con el resultante de multiplicar ambas fracciones originales
func (f Fraccion) Multiplicar(otra Fraccion) Fraccion {
	return CrearFraccion(f.numerador*otra.numerador, f.denominador*otra.denominador)
}

// ParteEntera devuelve la parte entera del numero representado por la fracción.
// Por ejemplo, para "7/2" = 3.5 debe devolver 3.
func (f Fraccion) ParteEntera() int {
	return f.numerador / f.denominador
}

// Representacion devuelve una representación en cadena de la fraccion simplificada (por ejemplo, no puede devolverse
// "10/8" sino que debe ser "5/4"). Considerar que si se trata de un número entero, debe mostrarse como tal.
// Considerar tambien el caso que se trate de un número negativo.
func maximoComunDivisor(a, b int) int {
	for a%b != 0 {
		a, b = b, (a % b)
	}
	return b
}

func valorAbsoluto(numero int) int {
	if numero < 0 {
		return -numero
	}
	return numero
}

func (f Fraccion) Representacion() string {
	if f.numerador%f.denominador == 0 {
		return fmt.Sprintf("%v", f.ParteEntera())
	}
	divisor := maximoComunDivisor(valorAbsoluto(f.numerador), valorAbsoluto(f.denominador))
	if divisor != 1 {
		return fmt.Sprintf("%v/%v", f.numerador/divisor, f.denominador/divisor)
	}
	return fmt.Sprintf("%v/%v", f.numerador, f.denominador)
}
