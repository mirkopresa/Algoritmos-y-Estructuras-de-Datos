package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

const _VOLUMEN = 10000

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	require.True(t, cola.EstaVacia(), "La cola deberia estar vacia")
	require.Panics(t, func() { cola.Desencolar() }, "La cola esta vacia por ende deberia retornar Panic")
	require.Panics(t, func() { cola.VerPrimero() }, "La cola esta vacia por ende deberia retornar Panic")
}

func TestFIFO(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 3; i++ {
		cola.Encolar(i * 2)
	}
	require.Equal(t, 0, cola.VerPrimero(), "Deberia retornar 0")
	cola.Desencolar()
	require.Equal(t, 2, cola.VerPrimero(), "Deberia retornar 2")
	cola.Desencolar()
	require.Equal(t, 4, cola.VerPrimero(), "Deberia retornar 4")
}

func TestDesencolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[float64]()
	for i := 0; i < 4; i++ {
		cola.Encolar(float64(i * 2))
	}
	require.Equal(t, 0.0, cola.Desencolar(), "Deberia retornar 0.0")
	require.Equal(t, 2.0, cola.Desencolar(), "Deberia retornar 2.0")
	cola.Encolar(8.0)
	require.Equal(t, 4.0, cola.Desencolar(), "Deberia retornar 4.0")
	require.Equal(t, 6.0, cola.Desencolar(), "Deberia retornar 6.0")
	require.Equal(t, 8.0, cola.Desencolar(), "Deberia retornar 8.0")
}

// Test que comprueba que la cola no se autodestruya al vaciarla y querer encolar
func TestColaNoSeDestruye(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[float32]()
	require.True(t, cola.EstaVacia(), "La cola deberia estar vacia")
	cola.Encolar(66.6)
	require.False(t, cola.EstaVacia(), "La cola no deberia estar vacia")
	cola.Desencolar()
	require.True(t, cola.EstaVacia(), "La cola deberia estar vacia")
	cola.Encolar(666.666)
	require.False(t, cola.EstaVacia(), "La cola no deberia estar vacia")
}

// Test que comprueba que c.primero y c.ultimo sean nil al encolar elementos y vaciarla
func TestComportamientoAlEncolarDesencolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[float64]()
	require.True(t, cola.EstaVacia(), "La cola deberia estar vacia")
	cola.Encolar(100)
	cola.Encolar(200)
	require.False(t, cola.EstaVacia(), "La cola no deberia estar vacia")
	cola.Desencolar()
	require.False(t, cola.EstaVacia(), "La cola no deberia estar vacia")
	cola.Desencolar()
	require.True(t, cola.EstaVacia(), "La cola deberia estar vacia")
}

func TestVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia(), "La cola deberia estar vacia")
	for i := 1; i <= _VOLUMEN; i++ {
		cola.Encolar(i * 10)
	}
	require.Equal(t, 10, cola.VerPrimero(), "Deberia retornar %d", 10)
	for !cola.EstaVacia() {
		desencolado := cola.VerPrimero()
		require.Equal(t, desencolado, cola.Desencolar(), "Deberia retornar %d", desencolado)
		if !cola.EstaVacia() {
			require.Equal(t, desencolado+10, cola.VerPrimero(), "Deberia retornar %d", desencolado+10)
		}
	}
	require.True(t, cola.EstaVacia(), "La cola deberia estar vacia")
}
