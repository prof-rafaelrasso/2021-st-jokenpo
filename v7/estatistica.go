package v7

type Estatistica struct {
	empates          int
	vitoriasJogador1 int
	vitoriasJogador2 int
}

func IniciarEstatistica() *Estatistica {
	return &Estatistica{}
}

func (e *Estatistica) ContabilizarResultado(resultado *Resultado) {
	if resultado.EhEmpate() {
		e.empates++
	} else if resultado.Jogador1EhOVencedor() {
		e.vitoriasJogador1++
	} else if resultado.Jogador2EhOVencedor() {
		e.vitoriasJogador2++
	}
}
