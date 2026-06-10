# Ejercicio 4: Camino Mínimo más forzado que tu relación con tu ex
# Implementar un algoritmo que, dado un grafo pesado, un vértice v, uno w y otro z, nos devuelva el camino mínimo entre v y w, 
# pasando sí o sí por el vértice z.
# Indicar y justificar la complejidad del algoritmo.

def forzar_relacion_con_tu_ex(grafo, inicio, mitad, fin):
    _, padres_1 = dijkstra(grafo, inicio, mitad)
    _, padres_2 = dijkstra(grafo, mitad, fin)
    camino_1 = []
    actual = mitad
    while actual is not None:
        camino_1.append(actual)
        actual = padres_1[actual]
    camino_2 = []
    actual = fin
    while actual is not mitad:
        camino_2.append(actual)
        actual = padres_2[actual]
    camino_1.reverse()
    camino_2.reverse()
    return camino_1 + camino_2

def dijkstra(grafo, origen, destino):
    dist = {}
    padres = {}
    for v in grafo: dist[v] = float("inf")
    dist[origen] = 0
    padres[origen] = None
    heap = Heap() # considerar implementado
    heap.encolar((0, origen))
    while not heap.esta_vacio():
        _, v = heap.desencolar()
        if v == destino: break
        for w in grafo.adyacentes(v):
            dist_entre = dist[v] + grafo.obtener_peso(v, w)
            if dist_entre < dist[w]:
                dist[w] = dist_entre
                padres[w] = v
                heap.encolar((dist[w], w))
    return dist, padres