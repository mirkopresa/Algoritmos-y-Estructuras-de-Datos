from typing import Callable, Any


class Heap:
    def __init__(self, cmp: Callable[[Any, Any], int]):
        self.datos = []
        self.cant = 0
        self.cmp = cmp

    def __str__(self) -> str:
        return f"{self.datos}"

    def esta_vacio(self) -> bool:
        return self.cant == 0

    def ver_max(self) -> Any:
        return self.datos[0]

    def cantidad(self) -> int:
        return self.cant

    def encolar(self, elem: Any) -> None:
        self.datos.append(elem)
        self.cant += 1
        self._upheap(self.cant - 1)

    def desencolar(self) -> Any:
        desencolado = self.ver_max()
        ultimo = self.datos.pop()
        self.cant -= 1
        if self.cant > 0:
            self.datos[0] = ultimo
            self._downheap(0)

        return desencolado

    def _upheap(self, pos: int) -> None:
        if pos == 0:
            return
        pos_padre = self._pos_padre(pos)
        comparacion = self.cmp(self.datos[pos_padre], self.datos[pos])
        if comparacion < 0:
            self._swap(pos_padre, pos)
            self._upheap(pos_padre)

    def _downheap(self, pos: int) -> None:
        pos_hijo_izq = self._pos_hijo_izq(pos)
        # Si no hay hijo izquierdo ni derecho
        if pos_hijo_izq >= self.cant:
            return
        hijo_max = pos_hijo_izq
        pos_hijo_der = self._pos_hijo_der(pos)
        if pos_hijo_der < self.cant:
            comparacion_hijos = self.cmp(
                self.datos[pos_hijo_izq], self.datos[pos_hijo_der]
            )
            if comparacion_hijos < 0:
                hijo_max = pos_hijo_der

        comparacion = self.cmp(self.datos[hijo_max], self.datos[pos])
        if comparacion > 0:
            self._swap(hijo_max, pos)
            self._downheap(hijo_max)

    def _pos_padre(self, pos: int) -> int:
        return (pos - 1) // 2

    def _pos_hijo_izq(self, pos: int) -> int:
        return pos * 2 + 1

    def _pos_hijo_der(self, pos: int) -> int:
        return pos * 2 + 2

    def _swap(self, a: int, b: int) -> None:
        self.datos[a], self.datos[b] = self.datos[b], self.datos[a]
