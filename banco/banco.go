package banco

import (
	"database/sql"
	"errors"

	"github.com/AnaJuliaNX/projeto_api/tipos"
	_ "github.com/go-sql-driver/mysql"
)

func BancoConectar() (*sql.DB, error) {

	stringDeConexao := "carro:carrinhos@/carros?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringDeConexao)
	if erro != nil {
		return nil, erro
	}
	erro = db.Ping()
	if erro != nil {
		return nil, erro
	}
	return db, nil
}

func BuscarCarrros() ([]tipos.Carros, error) {
	db, erro := BancoConectar()
	if erro != nil {
		return nil, erro
	}
	defer db.Close()

	linhas, erro := db.Query("select id, carro, fabricante from carros")
	if erro != nil {
		return nil, errors.New("erro ao buscar os carros")
	}
	defer linhas.Close()

	var carros []tipos.Carros
	var carro tipos.Carros
	for linhas.Next() {
		erro := linhas.Scan(&carro.ID, &carro.Carro, &carro.Fabricante)
		if erro != nil {
			return nil, errors.New("erro ao scanear os carros")
		}
		carros = append(carros, carro)
	}
	return carros, nil
}
