package http

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/AnaJuliaNX/projeto_api/banco"
	"github.com/AnaJuliaNX/projeto_api/tipos"
)

func GetAllCars(w http.ResponseWriter, r *http.Request) {

	db, erro := banco.BancoConectar()
	if erro != nil {
		log.Fatalf("Erro ao conectar no banco de dados %v", erro)
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select id, carro, fabricante from modelo_carro")
	if erro != nil {
		log.Fatalf("Erro ao obter os dados dos carros %v", erro)
		return
	}
	defer linhas.Close()

	var carros []tipos.Carros
	var carro tipos.Carros
	for linhas.Next() {
		erro = linhas.Scan(&carro.ID, &carro.Carro, &carro.Fabricante)
		if erro != nil {
			return
		}
		carros = append(carros, carro)
	}
	erro = json.NewEncoder(w).Encode(carros)
	if erro != nil {
		log.Fatalf("Erro ao converter para json %v", erro)
		return
	}
}

func AddCars(w http.ResponseWriter, r *http.Request) {

	corpo, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatalf("Errro ao ler os dados do corpo %v", erro)
		return
	}
	var body map[string]interface{}
	erro = json.Unmarshal(corpo, &body)
	if erro != nil {
		log.Fatalf("Erro ao converter para struct %v", erro)
		return
	}

	db, erro := banco.BancoConectar()
	if erro != nil {
		log.Fatalf("Erro ao conectar com o banco de dados %v", erro)
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("Insert into modelo_carro(carro, fabricante) values(?, ?)")
	if erro != nil {
		log.Fatalf("Erro ao criar o statement %v", erro)
		return
	}
	defer statement.Close()
	inserir, erro := statement.Exec(body["carro"], body["fabricante"])
	if erro != nil {
		log.Fatalf("Erro ao executar o statement %v", erro)
		return
	}
	_, erro = inserir.LastInsertId()
	if erro != nil {
		log.Fatalf("Erro ao obter o ID inserido %v", erro)
		return
	}
}
