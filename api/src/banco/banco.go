package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {

	db, err := sql.Open("mysql", config.StringConexaoBanco) // conecta com o banco
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil { // testa conexao
		db.Close()
		return nil, err
	}
	return db, nil

}
