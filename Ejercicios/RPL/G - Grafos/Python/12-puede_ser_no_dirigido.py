# Implementar una función que permita determinar si un grafo puede ser no dirigido. 
# Determinar el orden del algoritmo implementado.

def puede_ser_no_dirigido(grafo):
    for v in grafo: # O(V + E), pasamos por todos los vertices, y todas las aristas de ese vertice, y luego chequeamos (suponiendo que adyacentes devuelve un dict es O(1))
        for w in grafo.adyacentes(v): 
            if v not in grafo.adyacentes(w):
                return False
    return True
