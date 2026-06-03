# Implementar una función que devuelva una lista de tuplas que represente todas las aristas del grafo

# a. Para un grafo dirigido

# b. Para un grafo no dirigido, evitando devolver aristas repetidas


def obtener_aristas_dirigido(grafo):
    res = []
    for v in grafo:
        for w in grafo.adyacentes(v):
            res.append((v, w))
    return res

def obtener_aristas_no_dirigido(grafo):
    res = []
    visitados = set()
    for v in grafo:
        for w in grafo.adyacentes(v):
            if w not in visitados:
                res.append((v,w))
        visitados.add(v)
    return res