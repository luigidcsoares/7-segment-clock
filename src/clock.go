package main

import "sync"

// Cada linha da matriz abaixo descreve como o algarismo será mostrado na tela.
// Os dígitos são compostos de 7 segmentos, sendo 3 na horizontal (topo, meio e
// base) e 4 na vertical (esquerda/cima, esquerda/baixo, direita/cima,
// direita/baixo). Para cada um dos algarismos, um segmento pode ter seu valor
// definido como verdadeiro ou falso. Nesse caso, o verdadeiro pode ser tomado
// como um led aceso. Os segmentos estão definidos da seguinte forma:
//
//    __1__
// 0 |     | 2
//   |__3__|
// 4 |     | 6
//   |__5__|
//
var digitSegments = [][]bool{
	{true, true, true, false, true, true, true},     // 0
	{false, false, true, false, false, false, true}, // 1
	{false, true, true, true, true, true, false},    // 2
	{false, true, true, true, false, true, true},    // 3
	{true, false, true, true, false, false, true},   // 4
	{true, true, false, true, false, true, true},    // 5
	{true, true, false, true, true, true, true},     // 6
	{false, true, true, false, false, false, true},  // 7
	{true, true, true, true, true, true, true},      // 8
	{true, true, true, true, false, true, true},     // 9
}

func buildDigit(digitValue, segSize int) (digit [][]byte) {
	// Existem 2 segmentos de cada lado no eixo vertical. Assim, a quantidade
	// de linhas necessárias para representar o algarismo é dada por 2 * o
	// tamanho do segmento + 1 (a primeira linha é referente ao segmento
	// vertical do topo).
	digit = make([][]byte, 2*segSize+1)

	for i := range digit {
		// Segmentos verticais são representados por "|" (pipe) de
		// tamanho igual ao tamanho do segmento passada por parâmetro.
		// Por outro lado, segmentos horizontais (topo, meio, base) são
		// representados por 2 símbolos "_" (underscore). Assim, para um
		// tamanho de segmento = 1, na horizontal serão utilizados 2
		// underscores; já para um tamanho = 2, 4 underscores.
		// Dessa forma, a quantidade de colunas é dada por 2 * tamanho do
		// segmento + 2 (primeira e última colunas).
		digit[i] = make([]byte, 2+2*segSize)

		for j := range digit[i] {
			// Inicializa cada posição com espaço em branco (led apagado).
			digit[i][j] = ' '
		}
	}

	// Como segmentos horizontais e verticais não são correlacionados, é
	// possível construí-los de forma concorrente. Para tanto, é preciso
	// sincronizar as duas goroutines ao final do método, para que terminem de
	// executar corretamente.
	var wg sync.WaitGroup
	wg.Add(2)
	defer wg.Wait()

	go func() {
		defer wg.Done()

		// Segmentos horizontais (1, 3, 5).
		segment := 1
		for i := 0; i < 3; i++ {
			row := i * segSize

			for j := 1; j < len(digit[row])-1; j++ {
				if digitSegments[digitValue][segment] {
					digit[row][j] = '_'
				}
			}

			segment += 2
		}
	}()

	go func() {
		defer wg.Done()

		// Segmentos verticais (0, 2, 4, 6)
		// Esquerda: 0, 4 (col = 0)
		// Direita: 2, 6 (col = last)
		for i := 0; i <= 6; i += 2 {
			// Segmentos 0, 2 --> segmentos superiores
			// Segmentos 4, 6 --> segmentos inferiores
			add := 0
			if i >= 4 {
				add = segSize
			}

			row := 1 + add

			// 0; 4 % 4 = 0 --> Segmentos à esquerda
			// 2; 6 % 4 = 2 --> Segmentos à direita
			// 0 >> 1 = 0; 2 >> 1 = 1
			col := (len(digit[row]) - 1) * (i % 4 >> 1)

			for j := row; j < row+segSize; j++ {
				if digitSegments[digitValue][i] {
					digit[j][col] = '|'
				}
			}
		}
	}()

	return
}

// Mostra na tela um único dígito.
func printDigit(digit [][]byte, margin int) {
	for i := range digit {
		MoveCursorForward(margin)

		for j := range digit[i] {
			Printf("%c", digit[i][j])
		}

		Println()
	}

	Println()
}

// Mostra na tela uma unidade de tempo (hora, minuto, segundo).
func printClockUnit(unit [2]int, segSize, margin int) {
	rows := segSize*2 + 1
	cols := segSize*2 + 2

	for _, d := range unit {
		digit := buildDigit(d, segSize)
		printDigit(digit, margin)

		margin += cols + 1
		MoveCursorUp(rows + 1)
	}
}

// PrintClock recebe uma matriz representando hora:minuto:segundo. Cada linha
// da matriz refere-se a uma das unidades de tempo; cada coluna, a um dos
// algarismos. Por fim, mostra-se na tela o tempo no formato HH:MM:SS.
func PrintClock(time [3][2]int, segSize, margin int) {
	// Define o número de linhas e colunas.
	rows := segSize*2 + 1
	cols := segSize*2 + 2

	// Define a coluna inicial (em que se dará o início do print).
	currColumn := margin

	// Percorre cada unidade (hora, minuto, segundo) da matriz e mostra na tela
	// os algarismos.
	for i, unit := range time {
		printClockUnit(unit, segSize, currColumn)

		// São mostrados dois dígitos por iteração. Além disso, existe uma
		// margem passada por parâmetro que influencia não somente na distância
		// entre os dígitos, mas também no offset do primeiro dígito. Por esse
		// motivo, é preciso incrementar a coluna levando em consideração os 2
		// dígitos e a margem.
		currColumn += 3*cols + margin

		// Se a unidade atual é hora ou minuto, é preciso printar o separador
		// (:).
		if i < 2 {
			// Print do primeiro ponto (superior).
			// Além da margem (porque nesse ponto já estamos na posição do
			// próximo algarismo), subtraimos tamanho do segmento - 1 para
			// centralizar o ponto entre os dígitos.
			MoveCursorForward(currColumn - margin - (segSize - 1))
			MoveCursorDown(rows / 2)
			Printf("%s", "○")

			// Print do segundo ponto (inferior).
			MoveCursorBack(1)
			MoveCursorDown(1)
			Printf("%s", "○")

			// Retorna para posição inicial.
			// É preciso adicionar 1, uma vez que, ao printar, o cursor move
			// uma posição para a direita.
			MoveCursorBack(currColumn - margin + 1)
			MoveCursorUp(rows/2 + 1)
		}
	}

	MoveCursorDown(rows + 1)
	Flush()
}
