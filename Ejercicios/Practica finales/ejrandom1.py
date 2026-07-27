# Implementar un algoritmo que reciba un grafo no dirigido y detecte si el mismo es bipartito o no.
# Si lo es, devolver los dos subconjuntos de vertices que conforman cada particion
# Si no lo es, devolver None


def devolver_subconjuntos(grafo) -> tuple[list, list] | tuple[None, None]:
    subconjunto_1 = []
    subconjunto_2 = []
    visitados = set()
    orden = {}
    for v in grafo:
        if v not in visitados:
            bipartito = es_bipartito(grafo, v, orden, visitados)
            if not bipartito:
                return None, None
    cargar_subconjuntos(subconjunto_1, subconjunto_2, orden)
    return subconjunto_1, subconjunto_2


def es_bipartito(grafo, inicio, orden, visitados):
    orden[inicio] = 0
    visitados.add(inicio)
    cola = Cola()
    cola.encolar(inicio)
    while not cola.esta_vacia():
        v = cola.desencolar()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                orden[w] = 1 - orden[v]
                visitados.add(w)
                cola.encolar(w)
            else:
                if orden[v] == orden[w]:
                    return False
    return True


def cargar_subconjuntos(subconjunto_1, subconjunto_2, orden):
    for v in orden:
        if orden[v] == 0:
            subconjunto_1.append(v)
        else:
            subconjunto_2.append(v)
