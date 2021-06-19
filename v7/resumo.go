package v7

import (
	"fmt"
	"strings"
)

func ImprimirResumo(jogo *Jogo) {
	fmt.Println("Deseja visualizar resumo das partidas? Sim(S)/Não(N)")
	var visualizarResumo string
	fmt.Scan(&visualizarResumo)
	visualizarResumo = strings.ToLower(visualizarResumo)
	if visualizarResumo == "sim" || string(visualizarResumo[0]) == "s" {
		fmt.Println("-------------------------------------------------------")
		fmt.Println("Foi(ram) realizada(s)",
			jogo.TotalRodadas(), "rodada(s) com um total de \n",
			jogo.TotalEmpates(), "empate(s) \n",
			jogo.TotalJogador(1), "vitória(s) para o jogador 1 e \n",
			jogo.TotalJogador(2), "vitória(s) para o jogador 2")
		imprimirRodadas(jogo)
	}
}

func imprimirRodadas(jogo *Jogo) {
	for _, rodada := range jogo.rodadas {
		fmt.Println("-------------------------------------------------------")
		fmt.Println("Rodada", rodada.numero)
		fmt.Println("- Jogador1", rodada.jogador1.nome)
		fmt.Println("- Jogador2", rodada.jogador2.nome)
		fmt.Println(":", rodada.resultado.ImprimirResultado())
	}
}
