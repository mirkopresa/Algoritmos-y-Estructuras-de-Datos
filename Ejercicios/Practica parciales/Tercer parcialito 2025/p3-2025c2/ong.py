# Se quiere modelar un sitio que conecte ONGs con personas y se cuenta con la siguiente información:
# Para cada persona a qué ONGs donó.
# Para cada persona a qué otras personas sigue.
# Se considera que si una persona sigue a otra la distancia entre ellas es 1; y que si dos personas donaron a la misma ONG la
# distancia entre ellas también es 1.
# a) Modelar un grafo que represente estos datos (hecho en cuaderno)
# b) Dadas dos personas y el grafo modelado en el punto anterior, encontrar el camino más corto que las conecta.

def dijkstra(grafo, origen):
    dist = {}
    padres = {}
    for v in grafo: dist[v] = float("inf")
    dist[origen] = 0
    padres[origen] = None
    heap = Heap() # considerar como implementado
    heap.encolar((0, origen))
    while not heap.esta_vacia():
        _, v = heap.desencolar()
        for w in grafo.adyacentes(v):
            dist_entre = dist[v] + grafo.obtener_peso(v, w)
            if dist_entre < dist[w]:
                padres[w] = v
                dist[w] = dist_entre
                heap.encolar((dist[w], w))
    return dist, padres

def encontrar_personas(grafo, persona_1, persona_2):
    dist, padres = dijkstra(grafo, persona_1)
    if dist[persona_2] == float("inf"):
        return None # no se pudo encontrar
    camino = []
    actual = persona_2
    while actual is not None:
        camino.append(actual)
        actual = padres[actual]
    camino.reverse()
    return camino