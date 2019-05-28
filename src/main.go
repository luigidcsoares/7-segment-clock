package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"time"
)

var mapColors = map[string]string{
	"black":   ColorBlack,
	"red":     ColorRed,
	"green":   ColorGreen,
	"yellow":  ColorYellow,
	"blue":    ColorBlue,
	"magenta": ColorMagenta,
	"cyan":    ColorCyan,
	"white":   ColorWhite,
}

func printRealTime(segment chan int, pause chan bool) {
	// Valor inicial do tamanho do segmento.
	segSize := <-segment

	// Espaço entre as unidades.
	margin := 4

	// Número de linhas necessárias para mostrar o dígito.
	rows := segSize*2 + 1

	PrintClock(NowAsMatrix(), segSize, margin)
	for {
		select {
		case <-pause:
			segSize = <-segment
			rows = segSize*2 + 1
			PrintClock(NowAsMatrix(), segSize, margin)
		default:
			// Pausa o processo por um tempo pequeno, para garantir que nenhuma
			// anormalidade acontecerá na tela (Ctrl-C, por exemplo, durante a
			// movimentação do cursor).
			time.Sleep(time.Millisecond * 10)
			MoveCursorPreviousLine(rows + 1)
			PrintClock(NowAsMatrix(), segSize, margin)
		}
	}
}

func askParams(min, max int) (n, segSize int) {
	// Número de linhas printadas na tela.
	n = 0

	for {
		fmt.Print("## Entre com o tamanho do segmento [1..5]: ")
		n++

		fmt.Scanln(&segSize)
		fmt.Scan("\n")

		// Testa se o tamanho do segmento está na faixa especificada.
		if segSize >= min && segSize <= max {
			break
		}

		// Senão, pede o parâmetro para o usuário novamente.
		fmt.Println("## O valor deve ser entre 1 e 5!!!")
		n++
	}

	return
}

func main() {
	// Recupera os argumentos passados na linha de comando.
	// A primeira posição refere-se ao caminho do programa e, portanto, é
	// desconsiderada.
	color := mapColors["magenta"]

	if len(os.Args) > 1 {
		color = mapColors[os.Args[1]]
	}

	// Altera a cor do texto.
	fmt.Print(color)

	// Define a faixa de valores do parâmetro.
	min := 1
	max := 5

	// Requisita ao usuário o parâmetro.
	n, segSize := askParams(min, max)

	// Para redefinição do tamanho do segmento, usuário deve pressionar a tecla
	// ENTER; já para finalizar o programa, Ctrl-C.
	fmt.Println("## Pressione enter para redefinir os parâmetros")
	fmt.Println("## Pressione Ctrl-C para finalizar")

	// Aqui estamos lidando com a interrupção gerada pelo Ctrl-C, para
	// finalizar o programa da maneira mais adequada.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	go func() {
		// Recebe a interrupção do sistema operacional.
		<-sig

		// Mostra na tela a mensagem de saída do programa, fazendo com que um
		// possível caractere ^C seja apagado.
		fmt.Println("\rSee ya ;)")

		os.Exit(0)
	}()

	// Cria canais para comunicação entre processo principal e a rotina
	// responsável por mostrar o relógio em tempo real.
	pause := make(chan bool)
	segment := make(chan int)

	// Executa a rotina do relógio em uma nova goroutine.
	go printRealTime(segment, pause)

	// Envia o tamanho inicial do segmento.
	segment <- segSize

	// Cria um novo buffer para ler do stdin.
	// Usamos isso para ler uma entrada até atingir um caractere de quebra de
	// linha. Assim, sabemos que o usuário pressionou ENTER.
	reader := bufio.NewReader(os.Stdin)
	for {
		reader.ReadString('\n')

		// Se ENTER foi pressionado, pausamos a rotina responsável pelo
		// relógio, limpamos a tela e pedimos novamente o parâmetro ao usuário.
		pause <- true

		nrows := n + (segSize*2 + 1) + 4

		MoveCursorPreviousLine(nrows)
		ClearScreenEnd()
		Flush()

		n, segSize = askParams(min, max)

		fmt.Println("## Pressione enter para redefinir os parâmetros")
		fmt.Println("## Pressione Ctrl-C para finalizar")

		// Enviamos o novo valor para a rotina do relógio, fazendo com que
		// volte a executar.
		segment <- segSize
	}
}
