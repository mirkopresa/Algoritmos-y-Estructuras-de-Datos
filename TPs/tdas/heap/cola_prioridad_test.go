package cola_prioridad_test

import (
	"strings"
	TDAHeap "tdas/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

const _VOLUMEN = 100000

// Funciones de comparacion para los tests
func cmpInt(a, b int) int {
	return a - b
}

func cmpFloat(a, b float64) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func TestHeapVacio(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](cmpInt)
	require.True(t, heap.EstaVacia(), "El heap deberia estar vacio")
	require.Panics(t, func() { heap.VerMax() }, "El heap esta vacio por ende deberia retornar Panic")
	require.Panics(t, func() { heap.Desencolar() }, "El heap esta vacio por ende deberia retornar Panic")
}

func TestCantidad(t *testing.T) {
	heap := TDAHeap.CrearHeap[string](strings.Compare)
	require.Equal(t, 0, heap.Cantidad(), "La cantidad deberia ser 0")
	heap.Encolar("Hola")
	require.Equal(t, 1, heap.Cantidad(), "La cantidad deberia ser 1")
	heap.Encolar("Mundo")
	require.Equal(t, 2, heap.Cantidad(), "La cantidad deberia ser 2")
	heap.Desencolar()
	require.Equal(t, 1, heap.Cantidad(), "La cantidad deberia ser 1")
	heap.Desencolar()
	require.Equal(t, 0, heap.Cantidad(), "La cantidad deberia ser 0")
}

func TestVerMax(t *testing.T) {
	heap := TDAHeap.CrearHeap[string](strings.Compare)
	heap.Encolar("Hola")
	require.Equal(t, "Hola", heap.VerMax(), "Deberia retornar 'Hola'")
	heap.Encolar("Mundo")
	require.Equal(t, "Mundo", heap.VerMax(), "Deberia retornar 'Mundo'	")
	heap.Encolar("30")
	require.Equal(t, "Mundo", heap.VerMax(), "Deberia retornar 'Mundo'")
}

func TestEncolarMaximos(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](cmpInt)
	heap.Encolar(5)
	require.Equal(t, 5, heap.VerMax(), "Deberia retornar 5")
	heap.Encolar(10)
	require.Equal(t, 10, heap.VerMax(), "Deberia retornar 10")
	heap.Encolar(15)
	require.Equal(t, 15, heap.VerMax(), "Deberia retornar 15")
	heap.Encolar(20)
	require.Equal(t, 20, heap.VerMax(), "Deberia retornar 20")
}

func TestDesencolar(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](cmpInt)
	for i := 1; i <= 10; i++ {
		heap.Encolar(i)
	}
	for i := 10; i >= 1; i-- {
		require.Equal(t, i, heap.Desencolar(), "Deberia desencolar %i", i)
	}
}

func TestEncolarDesencolar(t *testing.T) {
	heap := TDAHeap.CrearHeap[float64](cmpFloat)
	heap.Encolar(4.5)
	heap.Encolar(40.2)
	heap.Encolar(30.1)
	require.Equal(t, 40.2, heap.Desencolar(), "Deberia retornar 40.2")
	require.Equal(t, 30.1, heap.Desencolar(), "Deberia retornar 30.1")
	require.Equal(t, 4.5, heap.Desencolar(), "Deberia retornar 4.5")
}

func TestHeapify(t *testing.T) {
	elementos := []int{5, 3, 8, 1, 2}
	heap := TDAHeap.CrearHeapArr(elementos, cmpInt)
	require.Equal(t, 8, heap.VerMax(), "El maximo deberia ser 8")
	heap.Desencolar()
	require.Equal(t, 5, heap.VerMax(), "El maximo deberia ser 5")
	heap.Desencolar()
	require.Equal(t, 3, heap.VerMax(), "El maximo deberia ser 3")
	heap.Desencolar()
	require.Equal(t, 2, heap.VerMax(), "El maximo deberia ser 2")
	heap.Desencolar()
	require.Equal(t, 1, heap.VerMax(), "El maximo deberia ser 1")
}

func TestHeapSort(t *testing.T) {
	elementos := []int{5, 3, 8, 1, 2}
	TDAHeap.HeapSort(elementos, cmpInt)
	require.Equal(t, []int{1, 2, 3, 5, 8}, elementos, "El arreglo deberia estar ordenado")
	elementosFloat := []float64{4.5, 40.2, 30.1, 2.3, 3.4, 420.14}
	TDAHeap.HeapSort(elementosFloat, cmpFloat)
	require.Equal(t, []float64{2.3, 3.4, 4.5, 30.1, 40.2, 420.14}, elementosFloat, "El arreglo deberia estar ordenado")
}

func TestHeapComportamientoAlVaciarHeap(t *testing.T) {
	heap := TDAHeap.CrearHeap[string](strings.Compare)
	heap.Encolar("4")
	heap.Encolar("Algoritmos")
	heap.Encolar("y")
	heap.Encolar("Estructuras")
	heap.Encolar("de")
	heap.Encolar("Datos")
	require.False(t, heap.EstaVacia(), "El heap no deberia estar vacio")
	require.Equal(t, "y", heap.VerMax(), "Deberia retornar 'y'")
	for i := 0; i < 6; i++ {
		heap.Desencolar()
	}
	require.True(t, heap.EstaVacia(), "El heap deberia estar vacio")
	require.Panics(t, func() { heap.VerMax() }, "El heap esta vacio por ende deberia retornar Panic")
	require.Panics(t, func() { heap.Desencolar() }, "El heap esta vacio por ende deberia retornar Panic")
}

func TestHeapVolumen(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](cmpInt)
	for i := 0; i < _VOLUMEN; i++ {
		heap.Encolar(i)
	}
	require.Equal(t, _VOLUMEN-1, heap.VerMax(), "El maximo deberia ser el ultimo elemento en ser encolado")
	for i := _VOLUMEN - 1; i >= 0; i-- {
		require.Equal(t, i, heap.Desencolar(), "Deberia retornar el maximo actual")
	}
	require.True(t, heap.EstaVacia(), "El heap deberia estar vacio")
}
