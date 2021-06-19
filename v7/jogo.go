package v7

import (
	"fmt"
	"strings"
)

type Jogo struct {
	rodadas     []*Rodada
	estatistica *Estatistica
	ligado      bool
}

func IniciarJogo() *Jogo {
	fmt.Println("Starting Jokenpo Game v7")
	return &Jogo{estatistica: IniciarEstatistica(), ligado: true}
}

func (j *Jogo) Ligado() bool {
	return j.ligado
}

func (j *Jogo) RodadaAtual() int {
	return len(j.rodadas) + 1
}

func (j *Jogo) ContabilizarRodada(rodada *Rodada) {
	j.rodadas = append(j.rodadas, rodada)
	j.estatistica.ContabilizarResultado(rodada.resultado)
}

func (j *Jogo) ContinuarJogando() {
	fmt.Println("Deseja Continuar? Sim(S)/Não(N)")
	var continuarPartida string
	fmt.Scan(&continuarPartida)
	continuarPartida = strings.ToLower(continuarPartida)
	j.ligado = continuarPartida == "sim" || string(continuarPartida[0]) == "s"
}

func (j *Jogo) TotalRodadas() int {
	return len(j.rodadas)
}

func (j *Jogo) TotalEmpates() int {
	return j.estatistica.empates
}

func (j *Jogo) TotalJogador(numero int) (total int) {
	switch numero {
	case 1:
		total = j.estatistica.vitoriasJogador1
	case 2:
		total = j.estatistica.vitoriasJogador2
	}
	return
}
