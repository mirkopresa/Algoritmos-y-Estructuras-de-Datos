# Implementar una función que reciba un grafo no dirigido y pesado implementado con listas de adyacencia 
# (diccionario de diccionarios) y devuelva una matriz que sea equivalente a la representación de matriz de adyacencia del mismo grafo. 
# Indicar y justificar el orden del algoritmo implementado.

def convertir_a_matriz(grafo) -> tuple[list[list], list]:
    'Devolver la Matriz construida'
    'El arreglo de mapeo debe contener a todos los vértices, en el mismo orden en el que están representados en la matriz.'
    arreglo_mapeo = grafo.obtener_vertices()
    largo_f_c = len(arreglo_mapeo)
    # Inicializamos la matriz con todos 0, VxV
    matriz = []
    for _ in range(largo_f_c):
        fila = []
        for _ in range(largo_f_c):
            fila.append(0)
        matriz.append(fila)
    # O(V) recorremos cada vertice y le asignamos su posicion a cada uno
    posiciones = {}
    for i in range(len(arreglo_mapeo)):
        posiciones[arreglo_mapeo[i]] = i
    # O(V + E) pasamos por cada vertice y vemos sus adyacentes
    for v in grafo:
        i = posiciones[v]
        for w in grafo.adyacentes(v):
            j = posiciones[w]
            if i == j:
                continue
            else:
                peso = grafo.peso_arista(v, w)
                matriz[i][j] = peso
                matriz[j][i] = peso
    # Complejidad final O(V*V + V + V + E) -> V*V crece mucho mas rapido -> O(V²)
    return matriz, arreglo_mapeo