package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

const _VOLUMEN = 10000

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	require.True(t, pila.EstaVacia(), "La pila deberia estar vacia")
	require.Panics(t, func() { pila.Desapilar() }, "La pila esta vacia por ende deberia retornar Panic")
	require.Panics(t, func() { pila.VerTope() }, "La pila esta vacia por ende deberia retornar Panic")
}

func TestApilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 5; i++ {
		pila.Apilar(i * 4)
		require.Equal(t, i*4, pila.VerTope(), "Deberia devolver %d", i*4)
	}
}

func TestDesapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float64]()
	for i := 0; i < 5; i++ {
		pila.Apilar(float64(i * 4))
	}
	require.Equal(t, 16.0, pila.Desapilar(), "Deberia desapilar y devolver 16.0")
	require.Equal(t, 12.0, pila.Desapilar(), "Deberia desapilar y devolver 12.0")
	require.NotPanics(t, func() { pila.Desapilar() }, "No deberia paniquear")
	require.Equal(t, 4.0, pila.Desapilar(), "Deberia desapilar y devolver 4.0")
	require.Equal(t, 0.0, pila.Desapilar(), "Deberia desapilar y devolver 0.0")
}

// Test que comprueba que la pila no se autodestruya al vaciar y querer apilar
func TestPilaNoSeDestruye(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float32]()
	require.True(t, pila.EstaVacia(), "La pila deberia estar vacia")
	pila.Apilar(2.0)
	require.False(t, pila.EstaVacia(), "La pila no deberia estar vacia")
	pila.Desapilar()
	require.True(t, pila.EstaVacia(), "La pila deberia estar vacia")
	pila.Apilar(48.0)
	require.False(t, pila.EstaVacia(), "La pila no deberia estar vacia")
}

// Test que comprueba que la pila se comporte como una pila vacia al apilar y desapilar todo
func TestComportamientoAlApilarDesapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia(), "La pila deberia estar vacia")
	pila.Apilar(2)
	pila.Apilar(4)
	require.False(t, pila.EstaVacia(), "La pila no deberia estar vacia")
	pila.Desapilar()
	pila.Desapilar()
	require.True(t, pila.EstaVacia(), "La pila deberia estar vacia")
	require.Panics(t, func() { pila.Desapilar() }, "La pila esta vacia por ende deberia retornar Panic")
	require.Panics(t, func() { pila.VerTope() }, "La pila esta vacia por ende deberia retornar Panic")
}

func TestVolumen(t *testing.T) {
	pilaVoluminosa := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pilaVoluminosa.EstaVacia(), "La pila deberia estar vacia")
	for i := 1; i <= _VOLUMEN; i++ {
		pilaVoluminosa.Apilar(i * 4)
		require.Equal(t, i*4, pilaVoluminosa.VerTope(), "Deberia devolver %d", i*4)
	}
	for j := _VOLUMEN; j >= _VOLUMEN-10; j-- {
		require.Equal(t, pilaVoluminosa.VerTope(), pilaVoluminosa.Desapilar())
		require.Equal(t, j*4-4, pilaVoluminosa.VerTope(), "El tope deberia actualizarse correctamente")
	}
	for !pilaVoluminosa.EstaVacia() {
		require.Equal(t, pilaVoluminosa.VerTope(), pilaVoluminosa.Desapilar(), "Deberia desapilar todo correctamente")
	}
	require.True(t, pilaVoluminosa.EstaVacia(), "La pila deberia estar vacia")
}
