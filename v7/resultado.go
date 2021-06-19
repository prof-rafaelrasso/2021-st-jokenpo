package v7

type Resultado struct {
	vencedor *Jogador
}

func DefinirResultado(jogador1, jogador2 *Jogador) Resultado {
	if jogador1.EscolheuAlgoQueVenceDaEscolhaDo(jogador2) {
		return Resultado{jogador1}
	}
	if jogador2.EscolheuAlgoQueVenceDaEscolhaDo(jogador1) {
		return Resultado{jogador2}
	}
	return Resultado{}
}

func (r *Resultado) ImprimirResultado() string {
	if r.EhEmpate() {
		return "Empatou!"
	} else {
		return "Quem venceu foi " + r.vencedor.nome
	}
}

func (r *Resultado) EhEmpate() bool {
	return r.vencedor == nil
}

func (r *Resultado) Jogador1EhOVencedor() bool {
	return !r.EhEmpate() && r.vencedor.numero == 1
}

func (r *Resultado) Jogador2EhOVencedor() bool {
	return !r.EhEmpate() && r.vencedor.numero == 2
}
