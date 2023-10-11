package banco

import (
	"database/sql"

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
