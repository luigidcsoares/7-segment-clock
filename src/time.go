package main

import (
	"time"
)

// NowAsMatrix retorna o tempo atual na forma de uma matriz de 3 linhas e 2
// colunas. Nesta, cada linha se refere a uma unidade de tempo (hora, minuto ou
// segundo), e cada coluna se refere a um algarismo da unidade.
func NowAsMatrix() (now [3][2]int) {
	h, m, s := time.Now().Clock()
	now[0] = [2]int{h / 10, h % 10}
	now[1] = [2]int{m / 10, m % 10}
	now[2] = [2]int{s / 10, s % 10}
	return
}
