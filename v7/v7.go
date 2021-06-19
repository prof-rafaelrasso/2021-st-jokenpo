package v7

import (
	"fmt"
	"os"
)

func RunApplication() {
	jogo := IniciarJogo()

	for jogo.Ligado() {
		rodada := IniciarRodada(jogo.RodadaAtual())
		rodada.IdentificarJogadores()
		rodada.EscolherSinais()
		resultado := rodada.EncerrarRodada()
		fmt.Println(resultado.ImprimirResultado())
		jogo.ContabilizarRodada(rodada)
		jogo.ContinuarJogando()
	}

	ImprimirResumo(jogo)
	fmt.Println("Finalizando o jogo! Bye Bye!")
	os.Exit(0)

}
