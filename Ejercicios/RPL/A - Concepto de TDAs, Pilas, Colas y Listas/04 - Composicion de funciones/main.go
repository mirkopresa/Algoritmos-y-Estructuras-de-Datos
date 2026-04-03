// Implementar en Go el TDA ComposiciónFunciones que emula la composición de funciones (i.e. f(g(h(x))).

package main

type composicionFunc struct {
	funciones []func(float64) float64
}

func CrearComposicion() ComposicionFunciones {
	return &composicionFunc{funciones: make([]func(float64) float64, 0)}
}

// Primitivas

func (f *composicionFunc) AgregarFuncion(nuevaFuncion func(float64) float64) {
	f.funciones = append(f.funciones, nuevaFuncion)
}

func (f composicionFunc) Aplicar(num float64) float64 {
	for i := len(f.funciones) - 1; i >= 0; i-- {
		num = f.funciones[i](num)
	}
	return num
}
