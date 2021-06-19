package v7

import (
	"errors"
	"fmt"
	"strings"
)

var sinaisPermitidos []Sinal = []Sinal{{"pedra"}, {"papel"}, {"tesoura"}}

type Sinal struct {
	valor string
}

func EscolherSinal() (sin Sinal, err error) {
	var sinalRecebido string
	fmt.Scan(&sinalRecebido)
	sinalRecebido = strings.ToLower(sinalRecebido)
	for _, sinalPermitido := range sinaisPermitidos {
		if sinalPermitido.ehIgual(sinalRecebido) {
			sin = sinalPermitido
			return
		}
	}
	err = errors.New("Sinal Inválido")
	return
}

func (s *Sinal) ehIgual(valor string) bool {
	return s.valor == valor
}

func (s *Sinal) GanhaDe(outroSinal Sinal) bool {
	return (s.valor == "pedra" && outroSinal.valor == "tesoura") ||
		(s.valor == "tesoura" && outroSinal.valor == "papel") ||
		(s.valor == "papel" && outroSinal.valor == "pedra")
}
