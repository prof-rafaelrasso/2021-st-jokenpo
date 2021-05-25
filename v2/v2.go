package v2

import "fmt"

// Nessa versão evoluimos nosso jogo par usar funções
func RunApplication() {
	fmt.Println("Iniciando JoKenPo: Versão 2")

	jogador1 := identificarJogador(1)
	jogador2 := identificarJogador(2)
	escolha1 := escolherSinal(jogador1)
	escolha2 := escolherSinal(jogador2)

	imprimirEscolha(jogador1, escolha1)
	imprimirEscolha(jogador2, escolha2)
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
	return escolha
}

func imprimirEscolha(jogador, escolha string) {
	fmt.Println(jogador, "escolheu", escolha)
}
