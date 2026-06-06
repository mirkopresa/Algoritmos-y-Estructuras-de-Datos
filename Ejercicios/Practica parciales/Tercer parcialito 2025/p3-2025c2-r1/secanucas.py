# Un famoso ladrón (ladrón en serio) apodado como El Seca Nucas, se encuentra planificando un gran robo en el barrio de
# Barracas. La Policía se encuentra haciendo un gran operativo para atraparlo. Luego de tantos robos están finalmente un paso
# adelante suyo. Tienen información de cuando va a ser el próximo asalto, pero no saben exactamente en qué comercio o banco
# será. Se tiene:

# Un grafo no dirigido de la ciudad donde los vértices representan las esquinas de la misma, y las aristas, las cuadras
# que conectan esas esquinas. Cada una tiene un peso que indica la distancia entre ellas. Por extrañas razones, ¡todos los
# comercios/bancos están sobre esquinas en este barrio!

# Una lista de esquinas de fuga F por las cuales si pasa El Seca Nucas se sabe que ya se puede dar a la fuga.

# Implementar un algoritmo que obtenga el top 3 de lugares en las cuales El Seca Nucas podría ejecutar su asalto. Una esquina
# es la mejor para realizar un asalto si la suma de las distancias de los caminos mínimos contra todas las esquinas de Fuga es la
# menor. Inclusive, alguna de estas esquinas también podría ser parte del top 3. Indicar y justificar la complejidad del mismo.

def seca_nucas(grafo, lista_fugas: list) -> list:
    distancias_global = {}
    for v in grafo: distancias_global[v] = 0
    for v in lista_fugas:
        dist, _ = dijkstra(grafo, v)
        for v in dist:
            distancias_global[v] += dist[v]
    esquinas = []
    for v in distancias_global:
        esquinas.append((distancias_global[v], v))
    return topk(esquinas, 3)

def dijkstra(grafo, origen):
    dist = {}
    padres = {}
    for v in grafo: dist[v] = float("inf")
    padres[origen] = None
    dist[origen] = 0
    heap = Heap()
    heap.encolar(origen)
    while not heap.esta_vacio():
        _, v = heap.desencolar()
        for w in grafo.adyacentes(v):
            peso_adj = grafo.obtener_peso(v, w)
            if dist[v] + peso_adj < dist[w]:
                dist[w] = dist[v] + peso_adj
                padres[w] = v
                heap.encolar(dist[w], w)
    return dist, padres

# ver .go para el topk