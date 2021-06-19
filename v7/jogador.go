package v7

import "fmt"

type Jogador struct {
	numero         int
	nome           string
	sinalEscolhido Sinal
}

func IdentificarJogador(numeroJogador int) *Jogador {
	var nomeInformado string
	fmt.Println("Informe o nome do Jogador", numeroJogador)
	fmt.Scan(&nomeInformado)
	return &Jogador{numero: numeroJogador, nome: nomeInformado}
}

func (j *Jogador) EscolherSinal() {
	fmt.Println(j.nome, "escolha entre Pedra, Papel ou Tesoura!")
	sinal, err := EscolherSinal()
	if err != nil {
		fmt.Println(err)
		j.EscolherSinal()
	} else {
		j.sinalEscolhido = sinal
	}

}

func (j *Jogador) EscolheuOMesmoQue(outroJogador *Jogador) bool {
	return j.sinalEscolhido == outroJogador.sinalEscolhido
}

func (j *Jogador) EscolheuAlgoQueVenceDaEscolhaDo(outroJogador *Jogador) bool {
	return j.sinalEscolhido.GanhaDe(outroJogador.sinalEscolhido)
}
