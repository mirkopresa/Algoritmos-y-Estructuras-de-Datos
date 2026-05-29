package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const _VOLUMEN = 10000

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia")
	require.Panics(t, func() { lista.VerPrimero() }, "La lista esta vacia por ende deberia retornar Panic")
	require.Panics(t, func() { lista.VerUltimo() }, "La lista esta vacia por ende deberia retornar Panic")
}

func TestInsertarElementosPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(4)
	require.Equal(t, 4, lista.VerPrimero(), "Deberia retornar 4")
	require.Equal(t, 4, lista.VerUltimo(), "Deberia retornar 4")
	lista.InsertarPrimero(40)
	require.Equal(t, 40, lista.VerPrimero(), "Deberia retornar 40")
	lista.InsertarPrimero(70)
	require.Equal(t, 70, lista.VerPrimero(), "Deberia retornar 70")
	lista.InsertarPrimero(420)
	require.Equal(t, 420, lista.VerPrimero(), "Deberia retornar 420")
	require.Equal(t, 4, lista.VerUltimo(), "Deberia retornar 4")
}

func TestInsertarElementosUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[float64]()
	lista.InsertarUltimo(4.0)
	require.Equal(t, 4.0, lista.VerUltimo(), "Deberia retornar 4.0")
	require.Equal(t, 4.0, lista.VerPrimero(), "Deberia retornar 4.0")
	lista.InsertarUltimo(40.7)
	require.Equal(t, 40.7, lista.VerUltimo(), "Deberia retornar 40.7")
	lista.InsertarUltimo(6.234)
	require.Equal(t, 6.234, lista.VerUltimo(), "Deberia retornar 6.234")
	lista.InsertarUltimo(10.1)
	require.Equal(t, 10.1, lista.VerUltimo(), "Deberia retornar 10.1")
	require.Equal(t, 4.0, lista.VerPrimero(), "Deberia retornar 4.0")

}

func TestBorrarElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(40)
	lista.InsertarPrimero(70)
	lista.InsertarPrimero(420)
	require.Equal(t, 420, lista.VerPrimero(), "Deberia retornar 420")
	require.Equal(t, 420, lista.BorrarPrimero(), "Deberia retornar 420")
	require.Equal(t, 70, lista.VerPrimero(), "Deberia retornar 70")
	require.Equal(t, 70, lista.BorrarPrimero(), "Deberia retornar 70")
	require.Equal(t, 40, lista.VerPrimero(), "Deberia retornar 40")
	require.Equal(t, 40, lista.BorrarPrimero(), "Deberia retornar 40")
	require.Equal(t, 4, lista.VerPrimero(), "Deberia retornar 4")
	require.Equal(t, 4, lista.BorrarPrimero(), "Deberia retornar 4")
}

func TestComportamientoAlVaciarLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int32]()
	lista.InsertarPrimero(4)
	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia")
	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia")
	require.Panics(t, func() { lista.VerPrimero() }, "La lista esta vacia por ende deberia retornar Panic")
	require.Panics(t, func() { lista.VerUltimo() }, "La lista esta vacia por ende deberia retornar Panic")
}

func TestLargoLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	require.Equal(t, 0, lista.Largo(), "Deberia devolver 0")
	lista.InsertarPrimero("Algoritmos")
	lista.InsertarPrimero("y")
	lista.InsertarPrimero("Estructuras")
	lista.InsertarPrimero("de")
	lista.InsertarPrimero("Datos")
	require.Equal(t, 5, lista.Largo(), "Deberia devolver 5")
	lista.BorrarPrimero()
	require.Equal(t, 4, lista.Largo(), "Deberia devolver 4")
}

func TestIteradorInterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 30; i++ {
		lista.InsertarUltimo(i)
	}
	suma := 0
	lista.Iterar(func(dato int) bool {
		suma += dato
		return true
	})
	require.Equal(t, 435, suma, "Deberia devolver 435")
}

func TestIteradorInternoCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 30; i++ {
		lista.InsertarUltimo(i)
	}
	suma := 0
	lista.Iterar(func(dato int) bool {
		if dato == 15 {
			return false
		}
		suma += dato
		return true
	})
	require.Equal(t, 105, suma, "Deberia devolver 105")
}

func TestInsertarPrimeroyUltimoEnIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(4)
	require.Equal(t, 4, iter.VerActual(), "Deberia devolver 4")
	iter.Avanzar()
	iter.Insertar(40)
	require.Equal(t, 4, lista.VerPrimero(), "Deberia devolver 4")
	require.Equal(t, 40, lista.VerUltimo(), "Deberia devolver 40")
}

func TestInsertarEnMedioEnIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(40)
	lista.InsertarUltimo(70)
	iter := lista.Iterador()
	// Avanzamos el iterador hasta el 40 (la mitad)
	iter.Avanzar()
	iter.Insertar(6)
	iter.Avanzar()
	require.Equal(t, 40, iter.VerActual(), "Deberia devolver 40")
	require.Equal(t, 4, lista.VerPrimero(), "Deberia devolver 4")
}

func TestBorrarPrimerElementoEnIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(40)
	lista.InsertarUltimo(70)
	iter := lista.Iterador()
	require.Equal(t, 4, iter.Borrar(), "Deberia devolver 4")
	require.Equal(t, 40, lista.VerPrimero(), "Deberia devolver 40")
}

func TestBorrarUltimoElementoEnIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(40)
	lista.InsertarUltimo(70)
	iter := lista.Iterador()
	iter.Avanzar()
	iter.Avanzar()
	require.Equal(t, 70, iter.Borrar(), "Deberia devolver 70")
	require.Equal(t, 40, lista.VerUltimo(), "Deberia devolver 40")
}

func TestBorrarElementoEnMedioEnIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(40)
	lista.InsertarUltimo(70)
	iter := lista.Iterador()
	iter.Avanzar()
	require.Equal(t, 40, iter.Borrar(), "Deberia devolver 40")
	for iter := lista.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
		require.NotEqual(t, 40, iter.VerActual(), "No deberia devolver 40")
	}
	require.Equal(t, 2, lista.Largo(), "Deberia quedar 2 elementos")
}

func TestIteradorExternoBorrarImpares(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < _VOLUMEN; i++ {
		lista.InsertarUltimo(i)
	}
	iter := lista.Iterador()
	for iter.HayAlgoMas() {
		if iter.VerActual()%2 != 0 {
			require.Equal(t, iter.VerActual(), iter.Borrar(), "Deberia borrar el numero impar")
		} else {
			iter.Avanzar()
		}
	}
	require.Equal(t, _VOLUMEN/2, lista.Largo(), "Deberia quedar la mitad de elementos")
	for iter := lista.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
		require.Equal(t, 0, iter.VerActual()%2, "Deberia quedar solo numeros pares")
	}
}

func TestVolumenIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	for i := 0; i <= _VOLUMEN; i++ {
		iter.Insertar(i * 2)
		require.Equal(t, i*2, iter.VerActual(), "Deberia retornar %d", i*2)
		iter.Avanzar()
	}
	require.Equal(t, _VOLUMEN+1, lista.Largo(), "Deberia quedar %d elementos", _VOLUMEN)
	require.Equal(t, 0, lista.VerPrimero(), "Deberia retornar 0")
	require.Equal(t, _VOLUMEN*2, lista.VerUltimo(), "Deberia retornar %d", _VOLUMEN*2)
}
