package service

import (
	"fmt"
	"net/http"
)

type CarService interface {
	BuscarDados()
}

var (
	CarDataChannel = make(chan *http.Response)
)

const (
	carServiceUrl = "localhost:8000/carros"
)

type BuscarCarsDadosService struct{}

func NewCarService() CarService {
	return &BuscarCarsDadosService{}
}

func (*BuscarCarsDadosService) BuscarDados() {
	cliente := http.Client{}
	fmt.Printf("Buscando a url %v", carServiceUrl)

	//Chamando a tabela dos carros
	resposta, _ := cliente.Get(carServiceUrl)

	//Escrever a resposta para o canal
	CarDataChannel <- resposta
}
