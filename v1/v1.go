package v1

import "fmt"

// Nessa versão aplicamos os conhecimentos de variável, println, endereçamento de memória e scan.
func RunApplication() {
	fmt.Println("Iniciando JoKenPo: Versão 1")

	var (
		jogador1 string
		jogador2 string
		escolha1 string
		escolha2 string
	)

	fmt.Println("Informe o nome do Jogador 1")
	fmt.Scan(&jogador1)

	fmt.Println("Informe o nome do Jogador 2")
	fmt.Scan(&jogador2)

	fmt.Println(jogador1, "escolha entre Pedra, Papel e Tesoura!")
	fmt.Scan(&escolha1)

	fmt.Println(jogador2, "escolha entre Pedra, Papel e Tesoura!")
	fmt.Scan(&escolha2)

	fmt.Println(jogador1, "escolheu", escolha1)
	fmt.Println(jogador2, "escolheu", escolha2)

}
