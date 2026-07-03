# Implementar un algoritmo que, dado un grafo no dirigido, conexo, y sin puentes (es decir, sin ninguna arista que al
# quitarla formaría más de una componente conexa), determine una dirección para cada arista, para que el grafo dirigido
# resultante sea fuertemente conexo (es decir, haya una única componente fuertemente conexa). Indicar y justificar la
# complejidad del algoritmo.


def obtener_grafo_fuertemente_conexo(grafo_no_dirigido):
    grafo_dirigido = Grafo(
        es_dirigido=True, vertices_init=grafo_no_dirigido.obtener_vertices()
    )
    visitados = set()
    padres = {}
    inicio = grafo_no_dirigido.vertice_aleatorio()
    visitados.add(inicio)
    padres[inicio] = None
    dfs(grafo_no_dirigido, grafo_dirigido, inicio, visitados, padres)
    return grafo_dirigido


def dfs(grafo_no_dirigido, grafo_dirigido, v, visitados, padres):
    for w in grafo_no_dirigido.adyacentes(v):
        if w not in visitados:
            visitados.add(w)
            grafo_dirigido.agregar_arista(v, w)
            padres[w] = v
            dfs(grafo_no_dirigido, grafo_dirigido, w, visitados, padres)
        elif padres[v] != w:
            grafo_dirigido.agregar_arista(v, w)
