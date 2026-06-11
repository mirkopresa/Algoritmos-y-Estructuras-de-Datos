# Contamos con un grafo dirigido que modela un ecosistema. En dicho grafo, cada vértice es una especie, y cada arista (v, w) 
# indica que v es depredador natural de w. Considerando la horrible tendencia del ser humano por llevar a la extinción especies, 
# algo que nos puede interesar es saber si existe alguna especie que, si llegara a desaparecer, rompería todo el ecosistema: 
# quienes la depredan no tienen un sustituto (y, por ende, pueden desaparecer también) y/o quienes eran depredados por esta ya no tienen amenazas, 
# por lo que crecerán descontroladamente. 
# Implementar un algoritmo que reciba un grafo de dichas características y devuelva una lista de todas las 
# especies que cumplan lo antes mencionado. Indicar y justificar la complejidad del algoritmo implementado.

def amenazados(grafo):
    'Devolver una lista con los vértices que cumplen la condición'
    grados_s, grados_e = grados_salida_entrada(grafo)
    resultado = []
    especies = set()
    for v in grafo:
        for w in grafo.adyacentes(v):
            if grados_e[w] == 1 and v not in especies:
                resultado.append(v)
                especies.add(v)
        if grados_s[v] == 1:
            for w in grafo.adyacentes(v):
                if w not in especies:
                    resultado.append(w)
                    especies.add(w)
    return resultado

def grados_salida_entrada(grafo):
    grados_e = {}
    grados_s = {}
    for v in grafo: 
        grados_e[v] = 0
    for v in grafo:
        adyacentes = grafo.adyacentes(v)
        grados_s[v] = len(adyacentes)
        for w in adyacentes:
            grados_e[w] += 1
    return grados_s, grados_e

