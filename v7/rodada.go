package v7

type Rodada struct {
	numero    int
	jogador1  *Jogador
	jogador2  *Jogador
	resultado *Resultado
}

func IniciarRodada(numero int) *Rodada {
	return &Rodada{numero: numero}
}

func (r *Rodada) IdentificarJogadores() {
	r.jogador1 = IdentificarJogador(1)
	r.jogador2 = IdentificarJogador(2)
}

func (r *Rodada) EscolherSinais() {
	r.jogador1.EscolherSinal()
	r.jogador2.EscolherSinal()
}

func (r *Rodada) EncerrarRodada() *Resultado {
	resultado := DefinirResultado(r.jogador1, r.jogador2)
	r.resultado = &resultado
	return &resultado
}
