package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(res http.ResponseWriter, statusCode int, dados any) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	if err := json.NewEncoder(res).Encode(dados); err != nil {
		log.Fatal(err)
	}
}

func ERRO(res http.ResponseWriter, statusCode int, err error) {
	JSON(res, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: err.Error(),
	})
}
