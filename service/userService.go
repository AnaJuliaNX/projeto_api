package service

import (
	"fmt"
	"net/http"
)

type UserService interface {
	BuscarDados()
}

var (
	UserDataChannel = make(chan *http.Response)
)

const (
	userServiceUrl = "localhost:8000/compradores"
)

type buscarUsersDadosService struct{}

func NewUserService() UserService {
	return &buscarUsersDadosService{}
}

func (*buscarUsersDadosService) BuscarDados() {
	cliente := http.Client{}
	fmt.Printf("Buscando a url %v", userServiceUrl)

	//Chamando a tabela dos carros
	resposta, _ := cliente.Get(userServiceUrl)

	//Escrever a resposta para o canal
	UserDataChannel <- resposta
}
