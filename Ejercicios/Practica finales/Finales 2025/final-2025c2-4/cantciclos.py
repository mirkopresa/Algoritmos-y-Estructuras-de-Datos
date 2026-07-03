# Implementar un algoritmo que determine si en un grafo hay ciclos de cantidad impar de vértices, o si únicamente tiene de cantidad
# par (si los tuviera). Indicar y justificar la complejidad del algoritmo implementado.

from collections import deque


# asumiendo no dirigido
def ciclos_impares_o_pares(grafo) -> bool:
    visitados = set()
    orden = {}
    padres = {}
    for v in grafo:
        if v not in visitados:
            orden[v] = 0
            padres[v] = None
            visitados.add(v)
            ciclo_impar, ciclo_par = es_bipartito_o_ciclo_par(
                grafo, v, padres, orden, visitados
            )
    # faltaria saber que devolver, pero como no se especifica... lo dejo asi


def es_bipartito_o_ciclo_par(grafo, v, padres, orden, visitados: set):
    ciclo_par = False
    ciclo_impar = False
    q = deque()
    q.append(v)
    while q:
        v = q.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                orden[w] = orden[v] + 1
                padres[w] = v
                visitados.add(w)
                q.append(w)
            elif padres[v] != w:
                if orden[v] == orden[w]:
                    ciclo_impar = True
                else:
                    ciclo_par = True
    return ciclo_impar, ciclo_par
