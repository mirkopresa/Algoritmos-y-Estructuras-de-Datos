# Implementar un algoritmo que dado un grafo no dirigido determine si el mismo es conexo. Se pide implementar utilizando un recorrido
# DFS. Indicar y justificar la complejidad del algoritmo.


def es_conexo(grafo) -> bool:
    if len(grafo) == 0:
        return True
    origen = grafo.vertice_aleatorio()
    visitados = set()
    visitados.add(origen)
    dfs(grafo, origen, visitados)
    return len(visitados) == len(grafo)


def dfs(grafo, v, visitados) -> None:
    for w in grafo.adyacentes(v):
        if w not in visitados:
            visitados.add(w)
            dfs(grafo, w, visitados)
