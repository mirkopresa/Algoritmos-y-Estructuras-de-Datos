# Implementar un algoritmo que reciba un grafo no dirigido y determine si para cada par de vértices del grafo existe y es único el
# camino entre ellos. El algoritmo implementado debe ser lineal en la cantidad de vértices y aristas.
# Justificar la complejidad del algoritmo.


def existe(grafo) -> bool:
    # importante
    if len(grafo) == 0:
        return True
    origen = grafo.vertice_aleatorio()
    padres = {}
    visitados = set()
    padres[origen] = None
    visitados.add(origen)
    ciclo = dfs(grafo, origen, padres, visitados)
    return len(visitados) == len(grafo) and not ciclo


def dfs(grafo, v, padres, visitados) -> bool:
    for w in grafo.adyacentes(v):
        if w not in visitados:
            padres[w] = v
            visitados.add(w)
            if dfs(grafo, w, padres, visitados):
                return True
        elif padres[v] != w:
            return True
    return False
