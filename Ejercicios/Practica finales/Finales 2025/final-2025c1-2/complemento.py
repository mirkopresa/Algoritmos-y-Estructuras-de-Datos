# Implementar un algoritmo que dado Grafo no dirigido nos devuelva su complemento. Es decir, un grafo en el que una arista (v, w)
# significa que v y w no son adyacentes en el grafo original. Indicar y justificar la complejidad del algoritmo implementado.


def complemento(grafo):
    complemento = Grafo(es_dirigido=False, vertices_init=grafo.obtener_vertices())
    for v in grafo:
        for w in grafo:
            if v != w and not grafo.estan_unidos(v, w):
                complemento.agregar_arista(v, w)
    return complemento
