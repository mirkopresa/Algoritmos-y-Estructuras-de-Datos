# Proponer una función para calcular el grafo traspuesto G^T de un grafo dirigido G. 
# El grafo traspuesto G^T posee los mismos vértices que G, pero con todas sus aristas invertidas (por cada arista (v, w) en G,
# existe una arista (w, v) en G^T). Indicar la complejidad para un grafo implementado con:

# a. lista de adyancencias

# b. matriz de adyacencias

from grafo import Grafo # yo no lo tengo pero el corrector si je

def grafo_traspuesto(grafo):
    grafo_transpuesto = Grafo(True, grafo.obtener_vertices())
    for v in grafo:
        for w in grafo.adyacentes(v):
            grafo_transpuesto.agregar_arista(w, v, 1)
    return grafo_transpuesto