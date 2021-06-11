package v5

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	rodadas                                     = []map[string]string{}
	empates, vitoriasJogador1, vitoriasJogador2 int
)

// Nessa versão evoluimos nosso game para utilizar arrays e/ou slices e/ou maps
func RunApplication() {
	fmt.Println("Iniciando JoKenPo: Versão 5")

	for {

		jogador1 := identificarJogador(1)
		jogador2 := identificarJogador(2)
		escolha1 := escolherSinal(jogador1)
		escolha2 := escolherSinal(jogador2)

		resultadoRodada := definirVencedor(escolha1, escolha2, jogador1, jogador2)
		fmt.Println(resultadoRodada)

		contabilizarRodada(jogador1, jogador2, resultadoRodada)

		iniciarProximaRodadaOuFinalizarJogo()
	}
}

func iniciarProximaRodadaOuFinalizarJogo() {
	fmt.Println("Deseja Continuar? Sim(S) ou Não(N)")
	var decisao string
	fmt.Scan(&decisao)

	decisao = strings.ToLower(decisao)
	if decisao != "sim" && decisao[0:1] != "s" {
		imprmirResumo()
		fmt.Println("Finalizando o jogo! Bye Bye!")
		os.Exit(0)
	}
}

func imprmirResumo() {
	fmt.Println("Deseja visualizar o resumo das partidas? Sim(S)/Não(N)")
	var visualizarResumo string
	fmt.Scan(&visualizarResumo)
	visualizarResumo = strings.ToLower(visualizarResumo)
	if visualizarResumo == "sim" || visualizarResumo[0:1] == "s" {
		fmt.Println("Foi(ram) realizada(s)", len(rodadas), "rodada(s) com um total de \n",
			empates, "empate(s) \n",
			vitoriasJogador1, "vitória(s) para o jogador 1 e \n",
			vitoriasJogador2, "vitóri(s) para o jogador2")
		imprimirRodadas()
	}
}

func imprimirRodadas() {
	for _, rodada := range rodadas {
		fmt.Println("---------------------------------------------------------")
		fmt.Println("Rodada:", rodada["Rodada"])
		fmt.Println("- Jogador1:", rodada["Jogador1"])
		fmt.Println("- Jogador2:", rodada["Jogador2"])
		fmt.Println(":", rodada["Resultado"])
	}
}

func contabilizarRodada(jogador1, jogador2, resultadoRodada string) {
	rodada := map[string]string{
		"Rodada":    strconv.Itoa(len(rodadas) + 1),
		"Jogador1":  jogador1,
		"Jogador2":  jogador2,
		"Resultado": resultadoRodada,
	}
	rodadas = append(rodadas, rodada)
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

func definirVencedor(escolha1 string, escolha2 string, jogador1 string, jogador2 string) (resultado string) {
	switch {
	case (escolha1 == escolha2):
		resultado = "Deu empate!"
		empates++
		return resultado
	case escolha1 == "pedra" && escolha2 == "tesoura":
		resultado = jogador1
		vitoriasJogador1++
	case escolha1 == "tesoura" && escolha2 == "papel":
		resultado = jogador1
		vitoriasJogador1++
	case escolha1 == "papel" && escolha2 == "pedra":
		resultado = jogador1
		vitoriasJogador1++
	default:
		resultado = jogador2
		vitoriasJogador2++
	}
	return resultado + " Venceu"
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
