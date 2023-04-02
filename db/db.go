package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=pgadmin dbname=testdb password=secure_password host=127.0.0.1 sslmode=disable" // parâmetros para teste
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
