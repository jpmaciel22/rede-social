package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.ConfigurarAmbiente()
	r := router.Gerar()

	fmt.Printf("Api inicializada na porta %d.", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r)) // faz o servidor escutar a porta da config e utilizando as rotas do router configurado

}
