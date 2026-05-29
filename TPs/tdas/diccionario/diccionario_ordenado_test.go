package diccionario_test

import (
	"cmp"
	"fmt"
	"math/rand"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMENABB = []int{1000, 5000, 12500, 25000, 50000, 100000}

func TestDiccionarioOrdenadoVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](cmp.Compare)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestUnElementoAbb(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[string, int](cmp.Compare)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestDiccionarioOrdenadoGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[string, string](cmp.Compare)
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestReemplazoDatoAbb(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](cmp.Compare)
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestDiccionarioOrdenadoBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](cmp.Compare)

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestConClavesNumericasAbb(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDADiccionario.CrearABB[int, string](cmp.Compare)
	clave := 10
	valor := "Gatito"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestConClavesStructsAbb(t *testing.T) {
	t.Log("Valida que tambien funcione con estructuras mas complejas")
	type basico struct {
		a string
		b int
	}
	type avanzado struct {
		w int
		x basico
		y basico
		z string
	}

	dic := TDADiccionario.CrearABB[avanzado, int](func(a, b avanzado) int {
		if a.w < b.w {
			return -1
		}
		if a.w > b.w {
			return 1
		}
		if a.z < b.z {
			return -1
		}
		if a.z > b.z {
			return 1
		}
		return 0
	})

	a1 := avanzado{w: 10, z: "hola", x: basico{a: "mundo", b: 8}, y: basico{a: "!", b: 10}}
	a2 := avanzado{w: 10, z: "aloh", x: basico{a: "odnum", b: 14}, y: basico{a: "!", b: 5}}
	a3 := avanzado{w: 10, z: "hello", x: basico{a: "world", b: 8}, y: basico{a: "!", b: 4}}

	dic.Guardar(a1, 0)
	dic.Guardar(a2, 1)
	dic.Guardar(a3, 2)

	require.True(t, dic.Pertenece(a1))
	require.True(t, dic.Pertenece(a2))
	require.True(t, dic.Pertenece(a3))
	require.EqualValues(t, 0, dic.Obtener(a1))
	require.EqualValues(t, 1, dic.Obtener(a2))
	require.EqualValues(t, 2, dic.Obtener(a3))
	dic.Guardar(a1, 5)
	require.EqualValues(t, 5, dic.Obtener(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))
	require.EqualValues(t, 5, dic.Borrar(a1))
	require.False(t, dic.Pertenece(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))

}

func TestClaveVaciaAbb(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccionario.CrearABB[string, string](cmp.Compare)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestValorNuloAbb(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionario.CrearABB[string, *int](cmp.Compare)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestIteradorInternoValoresAbb(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](cmp.Compare)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestIteradorInternoValoresConBorradosAbb(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](cmp.Compare)
	dic.Guardar(clave0, 7)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	dic.Borrar(clave0)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func ejecutarPruebaVolumenAbb(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, int](cmp.Compare)

	claves := make([]string, n)
	valores := make([]int, n)

	for i := 0; i < n; i++ {
		claves[i] = fmt.Sprintf("%08d", i)
		valores[i] = i
	}

	rand.Shuffle(n, func(i, j int) {
		claves[i], claves[j] = claves[j], claves[i]
		valores[i], valores[j] = valores[j], valores[i]
	})

	for i := 0; i < n; i++ {
		dic.Guardar(claves[i], valores[i])
	}
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que devuelva los valores correctos */
	ok := true
	for i := 0; i < n; i++ {
		ok = dic.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = dic.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	for i := 0; i < n; i++ {
		ok = dic.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
		ok = !dic.Pertenece(claves[i])
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Cantidad())
}

func BenchmarkDiccionarioOrdenado(b *testing.B) {
	b.Log("Prueba de stress del Diccionario. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves geeneradas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAMS_VOLUMENABB {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenAbb(b, n)
			}
		})
	}
}

func TestIterarDiccionarioOrdenadoVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, int](cmp.Compare)
	iter := dic.Iterador()
	require.False(t, iter.HayAlgoMas())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Avanzar() })
}

func TestIteradorIndependienteYOrdenado(t *testing.T) {
	t.Log("Prueba que varios iteradores funcionan por separado y devuelven las claves en orden alfabético.")
	dic := TDADiccionario.CrearABB[string, string](cmp.Compare)
	claves := []string{"B", "A", "C"} // Las guardo desordenadas
	for _, k := range claves {
		dic.Guardar(k, "")
	}

	// Iterador que se crea y no se usa (no debe afectar a los demás)
	dic.Iterador()

	iter := dic.Iterador()

	require.True(t, iter.HayAlgoMas())
	act, _ := iter.VerActual()
	require.EqualValues(t, "A", act)
	iter.Avanzar()

	require.True(t, iter.HayAlgoMas())
	act, _ = iter.VerActual()
	require.EqualValues(t, "B", act)
	iter.Avanzar()

	require.True(t, iter.HayAlgoMas())
	act, _ = iter.VerActual()
	require.EqualValues(t, "C", act)
	iter.Avanzar()

	require.False(t, iter.HayAlgoMas())
}

func ejecutarPruebasVolumenIteradorAbb(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, *int](cmp.Compare)

	claves := make([]string, n)
	valores := make([]int, n)

	/* Inserta 'n' parejas en el arbol */
	for i := 0; i < n; i++ {
		claves[i] = fmt.Sprintf("%08d", i)
		valores[i] = i
	}

	rand.Shuffle(len(claves), func(i, j int) {
		claves[i], claves[j] = claves[j], claves[i]
		valores[i], valores[j] = valores[j], valores[i]
	})

	for i := 0; i < n; i++ {
		dic.Guardar(claves[i], &valores[i])
	}

	// Prueba de iteración sobre las claves almacenadas.
	iter := dic.Iterador()
	require.True(b, iter.HayAlgoMas())

	ok := true
	var i int
	var clave string
	var valor *int

	for i = 0; i < n; i++ {
		if !iter.HayAlgoMas() {
			ok = false
			break
		}
		c1, v1 := iter.VerActual()
		clave = c1
		if clave == "" {
			ok = false
			break
		}
		valor = v1
		if valor == nil {
			ok = false
			break
		}
		*valor = n
		iter.Avanzar()
	}
	require.True(b, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(b, n, i, "No se recorrió todo el largo")
	require.False(b, iter.HayAlgoMas(), "El iterador debe estar al final luego de recorrer")

	ok = true
	for i = 0; i < n; i++ {
		if valores[i] != n {
			ok = false
			break
		}
	}
	require.True(b, ok, "No se cambiaron todos los elementos")
}

func BenchmarkIteradorAbb(b *testing.B) {
	b.Log("Prueba de stress del Iterador del Diccionario. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) b.N elementos, iterarlos todos sin problemas. Se ejecuta cada prueba b.N veces para generar " +
		"un benchmark")
	for _, n := range TAMS_VOLUMENABB {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebasVolumenIteradorAbb(b, n)
			}
		})
	}
}

func TestVolumenIteradorCorteAbb(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDADiccionario.CrearABB[int, int](cmp.Compare)

	/* Inserta 'n' parejas en el arbol */
	for i := 0; i < 10000; i++ {
		dic.Guardar(i, i)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.Iterar(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%100 == 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

func TestAbbOrden(t *testing.T) {
	t.Log("Valida que el iterador recorra las claves en orden ascendente")
	dic := TDADiccionario.CrearABB[string, int](cmp.Compare)

	claves := []string{"D", "A", "C", "B", "E"}
	ordenEsperado := []string{"A", "B", "C", "D", "E"}

	for _, c := range claves {
		dic.Guardar(c, 0)
	}

	iter := dic.Iterador()
	for _, claveEsperada := range ordenEsperado {
		require.True(t, iter.HayAlgoMas(), "El iterador debería tener más elementos")
		c, _ := iter.VerActual()
		require.Equal(t, claveEsperada, c)
		iter.Avanzar()
	}
	require.False(t, iter.HayAlgoMas(), "El iterador debería haber terminado")
}

func TestAbbOrdenado(t *testing.T) {
	t.Log("Prueba con un árbol que es solo una rama derecha (ordenado)")
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	for i := 1; i <= 5; i++ {
		dic.Guardar(i, i)
	}

	desde, hasta := 2, 4
	iter := dic.IteradorRango(&desde, &hasta)

	// Debería devolver 2, 3, 4 exactamente
	for i := 2; i <= 4; i++ {
		require.True(t, iter.HayAlgoMas())
		c, _ := iter.VerActual()
		require.Equal(t, i, c)
		iter.Avanzar()
	}
	require.False(t, iter.HayAlgoMas())
}

func TestAbbRango(t *testing.T) {
	t.Log("Valida la iteración por rango con límites que no existen en el árbol")
	dic := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })

	for i := 0; i <= 10; i += 2 {
		dic.Guardar(i, "valor")
	}
	// El árbol tiene: 0, 2, 4, 6, 8, 10

	t.Run("Rango intermedio", func(t *testing.T) {
		// Rango [3, 7] -> Debería devolver 4, 6
		desde, hasta := 3, 7
		iter := dic.IteradorRango(&desde, &hasta)

		require.True(t, iter.HayAlgoMas(), "El iterador debería tener más elementos")
		c1, _ := iter.VerActual()
		require.Equal(t, 4, c1)
		iter.Avanzar()

		require.True(t, iter.HayAlgoMas())
		c2, _ := iter.VerActual()
		require.Equal(t, 6, c2)
		iter.Avanzar()

		require.False(t, iter.HayAlgoMas())
	})
}

func TestAbbIteradorInternoRango(t *testing.T) {
	t.Log("Valida que el iterador interno respete el rango y el corte (visitar)")
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	for i := 1; i <= 10; i++ {
		dic.Guardar(i, i)
	}

	desde, hasta := 3, 8
	contador := 0

	// Solo queremos sumar los primeros 2 elementos que encontremos en el rango [3, 8]
	dic.IterarRango(&desde, &hasta, func(c int, v int) bool {
		contador++
		if contador == 2 {
			return false // Corte prematuro
		}
		return true
	})

	// Debería haber visitado el 3 y el 4, y luego frenar.
	require.Equal(t, 2, contador)
}

func TestAbbIteradorRangoVacio(t *testing.T) {
	t.Log("Valida que el iterador de rango sobre un árbol vacío no tenga elementos")
	dic := TDADiccionario.CrearABB[int, int](cmp.Compare)
	desde, hasta := 1, 10
	iter := dic.IteradorRango(&desde, &hasta)
	require.False(t, iter.HayAlgoMas())
}

func TestAbbRangoInvertido(t *testing.T) {
	t.Log("Valida que un rango donde 'desde' es mayor que 'hasta' devuelva un iterador vacío")
	dic := TDADiccionario.CrearABB[int, int](cmp.Compare)
	for i := 1; i <= 10; i++ {
		dic.Guardar(i, i)
	}
	desde, hasta := 10, 5
	iter := dic.IteradorRango(&desde, &hasta)
	require.False(t, iter.HayAlgoMas())
}

func TestAbbRangoAbierto(t *testing.T) {
	t.Log("Prueba rango desde nil hasta una clave")
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	for i := 1; i <= 10; i++ {
		dic.Guardar(i, i)
	}

	var desde *int = nil
	hasta := 3

	iter := dic.IteradorRango(desde, &hasta)
	// Debería devolver 1, 2, 3
	contador := 0
	for iter.HayAlgoMas() {
		contador++
		iter.Avanzar()
	}
	require.Equal(t, 3, contador)
}

func TestAbbBorrarUnHijo(t *testing.T) {
	t.Log("Valida el borrado de un nodo que tiene un solo hijo (izquierdo o derecho)")
	dic := TDADiccionario.CrearABB[int, int](cmp.Compare)

	// Nodo con solo hijo derecho
	dic.Guardar(10, 10)
	dic.Guardar(15, 15)
	dic.Borrar(10)
	require.True(t, dic.Pertenece(15))
	require.EqualValues(t, 1, dic.Cantidad())

	// Nodo con solo hijo izquierdo
	dic.Guardar(20, 20)
	dic.Guardar(18, 18)
	dic.Borrar(20)
	require.True(t, dic.Pertenece(18))
	require.EqualValues(t, 2, dic.Cantidad())
}

func TestAbbBorradoDosHijos(t *testing.T) {
	t.Log("Valida el borrado de un nodo con dos hijos y que el árbol siga siendo válido")
	dic := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	claves := []int{10, 5, 15, 2, 7, 12, 20}
	for _, k := range claves {
		dic.Guardar(k, "valor")
	}

	// Borramos la raíz (10) que tiene dos hijos.
	valor := dic.Borrar(10)
	require.Equal(t, "valor", valor)
	require.False(t, dic.Pertenece(10))
	require.Equal(t, 6, dic.Cantidad())

	// Verificamos que el orden se mantenga intacto
	ordenEsperado := []int{2, 5, 7, 12, 15, 20}
	iter := dic.Iterador()
	for _, exp := range ordenEsperado {
		require.True(t, iter.HayAlgoMas())
		c, _ := iter.VerActual()
		require.Equal(t, exp, c)
		iter.Avanzar()
	}
}
