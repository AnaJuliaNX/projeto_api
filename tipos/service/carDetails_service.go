package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AnaJuliaNX/projeto_api/details"
	"github.com/AnaJuliaNX/projeto_api/tipos"
)

type CarDetailsService struct {
	GetDetails() details.CarDetails //Existe um erro nessa parte que não consegui resolver AINDA
}

var (
	carService      CarService  = NewCarService()
	userService     UserService = NewUserService() 
	carDataChannel              = make(chan *http.Response)
	userDataChannel             = make(chan *http.Response)
)

type service struct{}

func NewCarDetailsService() CarDetailsService{}

func (*service) GetDetails() details.carDetails {
	//Goroutine chamada endopoint 1, para pegar os dados da tabela modelo_carro
	go carService.BuscarDados()
	//Goroutine chamada endpoint 2, para pegar os dados da tabela compradores
	go userService.BuscarDados()
	//carChannel, canal para obter os dados do carro do endpoint 1
	carro, _ := getCarData()
	//userChannel, canal para obter os dados dos compradores do endpoint 2
	comprador, _ := getUserData()

	return details.CarDetails{
		ID_carro:     carro.ID,
		Carro:        carro.Carro,
		Fabricante:   carro.Fabricante,
		PrimeiroNome: comprador.PrimeiroNome,
		UltimoNome:   comprador.UltimoNome,
	}
}

func getCarData() (tipos.CarData, error) {
	r1 := <-carDataChannel
	var carro tipos.CarData
	erro := json.NewDecoder(r1.Body).Decode(&carro)
	if erro != nil {
		log.Fatalf("Erro ao converter para struct: %v", erro)
		return carro, erro
	}
	return carro, nil
}

func getUserData() (tipos.CompradoresData, error) {
	r2 := <-userDataChannel
	var comprador tipos.CompradoresData
	erro := json.NewDecoder(r2.Body).Decode(&comprador)
	if erro != nil {
		log.Fatalf("Erro ao converter para struct: %v", erro)
		return comprador, erro
	}
	return comprador, nil
}
