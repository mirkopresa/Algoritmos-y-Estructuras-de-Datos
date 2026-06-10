# Ejercicio 1: Este ejercicio es una bomba
# Se está modelando un juego en el que hay personajes en casilleros.
# Se quiere introducir un arma que actúa con una bomba, que cae en un casillero y afecta a todos las 
# casilleros que están a n pasos o menos del objetivo.

# Sabiendo que el nivel se modela como un grafo no pesado y no dirigido en el que cada casillero es un vértice, 
# implementar la función obtener_por_distancia que recibe un grafo, un identificar de un verfice objetivo y un valor n y 
# devuelva una lista con todos los vértices afectados por el impacto. 
# Indicar y justificar la complejidad del mismo.

def obtener_por_distancia(grafo, objetivo, n) -> list:
    # Operaciones O(1)
    afectados = []
    visitados = set()
    orden = {}
    cola = Cola()
    orden[objetivo] = 0
    visitados.add(objetivo)
    cola.encolar(objetivo)
    # Recorrido tipo BFS, complejidad O(V + E)
    # En el peor de los casos (n muy grande) entraran todos los vertices a la cola, y por cada uno de ellos, veremos todas las aristas
    while not cola.esta_vacia():
        v = cola.desencolar()
        if orden[v] > n:
            break
        afectados.append(v)
        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                orden[w] = orden[v] + 1
                cola.encolar(w)
    # Complejidad final: O(V + E)
    return afectados