# Implementar una función que reciba un grafo no dirigido, y que compruebe la siguiente afirmación:
# "La cantidad de vértices de grado IMPAR es PAR". Indicar y justificar el orden del algoritmo si
# el grafo está implementado como matriz de adyacencia.

# O(3V + E) -> O(V + E)
# O(V²) con matriz de adyacencia
def comprobar_teorema(grafo):
    grados = {}
    for v in grafo: # O(V) - O(V) con matriz de adyacencia
        grados[v] = 0
    for v in grafo: # O(V + E) - O(V²) con matriz de adyacencia ya que por cada vertice V recorremos toda la fila de V vertices
        for w in grafo.adyacentes(v):
            grados[w] += 1
    contador = 0
    for v in grados: # O(V)
        if grados[v] % 2 != 0:
            contador += 1
    return contador % 2 == 0