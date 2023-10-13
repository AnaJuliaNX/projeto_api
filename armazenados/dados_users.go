package http

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/AnaJuliaNX/projeto_api/banco"
	"github.com/AnaJuliaNX/projeto_api/tipos"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	db, erro := banco.BancoConectar()
	if erro != nil {
		log.Fatalf("Erro ao conectar no banco de dados %v", erro)
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select id, primeiro_nome, ultimo_nome from compradores")
	if erro != nil {
		log.Fatalf("Erro ao obter os dados dos carros %v", erro)
		return
	}
	defer linhas.Close()

	var compradores []tipos.Compradores
	var comprador tipos.Compradores
	for linhas.Next() {
		erro = linhas.Scan(&comprador.ID, &comprador.PrimeiroNome, &comprador.UltimoNome)
		if erro != nil {
			return
		}
		compradores = append(compradores, comprador)
	}
	erro = json.NewEncoder(w).Encode(compradores)
	if erro != nil {
		log.Fatalf("Erro ao converter para json %v", erro)
		return
	}
}

func AddUsers(w http.ResponseWriter, r *http.Request) {

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

	statement, erro := db.Prepare("Insert into compradores(primeiro_nome, ultimo_nome) values(?, ?)")
	if erro != nil {
		log.Fatalf("Erro ao criar o statement %v", erro)
		return
	}
	defer statement.Close()
	inserir, erro := statement.Exec(body["primeiro_nome"], body["ultimo_nome"])
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
