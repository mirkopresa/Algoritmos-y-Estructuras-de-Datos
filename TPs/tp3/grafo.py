import random


class Grafo:
    def __init__(self, es_dirigido: bool = False, vertices_init: list = []) -> None:
        self.vertices = {v: {} for v in vertices_init}
        self.dirigido = es_dirigido

    def __str__(self) -> str:
        if self.dirigido:
            tipo_grafo = "Grafo dirigido"
        else:
            tipo_grafo = "Grafo no dirigido"
        cant_aristas = 0
        visitados = set()
        for v in self.vertices:
            adyacentes = self.adyacentes(v)
            for w in adyacentes:
                if w in visitados and not self.dirigido:
                    continue
                cant_aristas += 1
            visitados.add(v)
        return (
            f"{tipo_grafo} con {len(self.vertices)} vertices y {cant_aristas} aristas"
        )

    def __iter__(self):
        return iter(self.vertices)

    def __len__(self):
        return len(self.vertices)

    def agregar_vertice(self, v) -> None:
        if v not in self.vertices:
            self.vertices[v] = {}
        else:
            raise vertice_ya_existe(f"El vertice {v} ya esta agregado")

    def borrar_vertice(self, v) -> None:
        if v in self.vertices:
            del self.vertices[v]
            for w in self.vertices:
                if v in self.vertices[w]:
                    del self.vertices[w][v]
        else:
            raise vertice_inexistente(f"El vertice {v} no existe en el grafo")

    def agregar_arista(self, v, w, peso: int = 1) -> None:
        if v in self.vertices and w in self.vertices:
            self.vertices[v][w] = peso
            if not self.dirigido:
                self.vertices[w][v] = peso
        if v not in self.vertices:
            raise vertice_inexistente(f"El vertice {v} no existe en el grafo")
        if w not in self.vertices:
            raise vertice_inexistente(f"El vertice {w} no existe en el grafo")

    def borrar_arista(self, v, w) -> None:
        if self.estan_unidos(v, w):
            del self.vertices[v][w]
            if not self.dirigido and v != w:
                del self.vertices[w][v]
        else:
            raise arista_inexistente(f"La arista ({v}, {w}) no existe en el grafo")

    def estan_unidos(self, v, w) -> bool:
        return v in self.vertices and w in self.vertices[v]

    def peso_arista(self, v, w) -> int:
        if self.estan_unidos(v, w):
            return self.vertices[v][w]
        raise arista_inexistente(f"La arista ({v}, {w}) no existe en el grafo")

    def obtener_vertices(self) -> list:
        return [v for v in self.vertices]

    def vertice_aleatorio(self):
        return random.choice(self.obtener_vertices())

    def adyacentes(self, v) -> list:
        if v not in self.vertices:
            raise vertice_inexistente(f"El vertice {v} no existe en el grafo")
        return [w for w in self.vertices[v]]


class grafo_error(Exception):
    pass


class vertice_inexistente(grafo_error):
    pass


class arista_inexistente(grafo_error):
    pass


class arista_ya_existe(grafo_error):
    pass


class vertice_ya_existe(grafo_error):
    pass
