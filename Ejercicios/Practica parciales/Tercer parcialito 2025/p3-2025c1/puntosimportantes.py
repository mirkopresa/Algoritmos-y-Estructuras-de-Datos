# En una pequeña ciudad costera, todos los habitantes usan el AlgoMaps, una aplicación que les ayuda a decidir sus rutas de
# viaje diarias. Conociendo su funcionamiento, sabemos que siempre ofrece la ruta del camino más corto de un punto a otro
# dadas las rutas existentes, sabiendo que cuenta con la información de la distancia cubierta por cada camino entre punto y
# punto. En el futuro se planifican ciertas obras en diferentes puntos de la ciudad que seguramente afecten a la red de transporte.
# Para evitar afectar en demasía a los ciudadanos, se busca analizar la información disponible para planificar las obras. Sería
# deseable evitar bloquear el transporte de los puntos más importantes de la ciudad.
# a. Explicar cómo la información disponible se puede modelar con un grafo, y explicar sus características. ¿Qué información
# debemos solicitarle a AlgoMaps? (hecho en cuaderno)
# b. Desarrollar una función puntosImportantes(grafo, k) que reciba el grafo definido en el punto (a) y que devuelva los k
# puntos más importantes en la ciudad, según lo modelado. Indicar y justificar su complejidad.

def puntos_importantes(grafo, k) -> list:
    centralidad = {}
    for v in grafo:
        centralidad[v] = 0
    for v in grafo:
        dist, padres = dijkstra(grafo, v)
        centralidad_aux = {}
        for w in grafo:
            centralidad_aux[w] = 0
        vertices_ordenados = ordenar_vertices(grafo, dist)
        for w in vertices_ordenados:
            centralidad_aux[padres[w]] += 1 + centralidad_aux[w]
        for w in grafo:
            if v == w: continue
            centralidad[w] += centralidad_aux[w]
    lista = []
    for v in centralidad:
        lista.append(v)
    return topk(lista, k) # imaginar que esta implementado en Go

def dijkstra(grafo, origen):
    distancia = {}
    padres = {}
    for v in grafo:
        distancia[v] = float("inf")
    padres[origen] = None
    distancia[origen] = 0
    q = Heap()
    q.encolar((0, origen))
    while not q.esta_vacia():
        peso, v = q.desencolar()
        for w in grafo.adyacentes(v):
            peso_ady = grafo.obtener_peso(v, w)
            if distancia[v] + peso_ady < distancia[w]:
                distancia[w] = distancia[v] + peso_ady
                padres[w] = v
                q.encolar(distancia[w], w)
    return distancia, padres

def ordenar_vertices(grafo, distancias):
    vertices = []
    for v in grafo:
        if distancias[v] != float("inf"):
            vertices.append(v)
    return sorted(vertices, key=lambda x: distancias[x], reverse=True)