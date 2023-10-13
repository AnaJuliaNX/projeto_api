package main

import (
	"fmt"
	"log"
	"net/http"

	armazenados "github.com/AnaJuliaNX/projeto_api/armazenados"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/carros", armazenados.AddCars).Methods(http.MethodPost)
	router.HandleFunc("/carros", armazenados.GetAllCars).Methods(http.MethodGet)

	router.HandleFunc("/compradores", armazenados.AddUsers).Methods(http.MethodPost)
	router.HandleFunc("/compradores", armazenados.GetAllUsers).Methods(http.MethodGet)

	fmt.Println("Executando na porta 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
