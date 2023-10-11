package http

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var (
	postService service.PostService = service.NewPostService()
)

type PostController interface {
	GetAllCars(w http.ResponseWriter, r *http.Request)
	AddCars(w http.ResponseWriter, r *http.Request)
}

type Controller struct{}

func NewPostController() PostController {
	return &controller{}
}

func (*controller) GetAllCars(w http.ResponseWriter, r *http.Request) {

	corpo, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatalf("Erro ao ler os dados do  corpo: %v", erro)
		return
	}

	var body map[string]interface{}
	erro = json.Unmarshal(corpo, &body)
	if erro != nil {

	}
}
