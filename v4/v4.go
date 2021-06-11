package v4

import (
	"fmt"
	"os"
	"strings"
)

// Nessa versão evoluimos nossa aplicação par utilizar um laço de repetição
func RunApplication() {
	fmt.Println("Iniciando JoKenPo: Versão 4")

	for {

		jogador1 := identificarJogador(1)
		jogador2 := identificarJogador(2)
		escolha1 := escolherSinal(jogador1)
		escolha2 := escolherSinal(jogador2)

		definirVencedor(escolha1, escolha2, jogador1, jogador2)

		imprimirEscolha(jogador1, escolha1)
		imprimirEscolha(jogador2, escolha2)

		fmt.Println("Deseja Continuar? Sim(S) ou Não(N)")
		var decisao string
		fmt.Scan(&decisao)

		decisao = strings.ToLower(decisao)
		if decisao != "sim" && decisao[0:1] != "s" {
			fmt.Println("Finalizando o jogo! Bye Bye!")
			os.Exit(0)
		}

	}
}

func identificarJogador(numeroJogador int) string {
	fmt.Println("Informe o nome do Jogador", numeroJogador)
	var jogador string
	fmt.Scan(&jogador)
	return jogador
}

func escolherSinal(jogador string) (escolha string) {
	fmt.Println(jogador, "escolha entre Pedra, Papel e Tesoura!")
	fmt.Scan(&escolha)
	return validarEscolhaComSwitch(jogador, escolha)
}

func imprimirEscolha(jogador, escolha string) {
	fmt.Println(jogador, "escolheu", escolha)
}

func definirVencedor(escolha1 string, escolha2 string, jogador1 string, jogador2 string) {
	switch {
	case (escolha1 == escolha2):
		println("Deu empate!")
	case escolha1 == "pedra" && escolha2 == "tesoura":
		println(jogador1, "Venceu")
	case escolha1 == "tesoura" && escolha2 == "papel":
		println(jogador1, "Venceu")
	case escolha1 == "papel" && escolha2 == "pedra":
		println(jogador1, "Venceu")
	default:
		println(jogador2, "Venceu")
	}
}

func validarEscolhaComSwitch(jogador, escolha string) string {
	escolha = strings.ToLower(escolha)
	switch escolha {
	case "pedra", "papel", "tesoura":
		return escolha
	default:
		return escolherSinal(jogador)
	}
}

func validarEscolhaComIf(jogador, escolha string) string {
	escolha = strings.ToLower(escolha)
	if escolha == "pedra" || escolha == "papel" || escolha == "tesoura" {
		return escolha
	} else {
		fmt.Println("A opção escolhida é inválida.")
		return escolherSinal(jogador)
	}
}
