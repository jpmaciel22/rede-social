package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var usuario modelos.Usuario
	if err = json.Unmarshal(body, &usuario); err != nil {
		log.Fatal(err)
	}

	bd, err := banco.Conectar()
	if err != nil {
		log.Fatal(err)
	}

	repositorio := repositorios.NovoRepositorioDeUsuarios(bd) // so para poder ter os metodos no banco como "executador do metodo"
	usuarioId, err := repositorio.Criar(usuario)              // equivalente ao service createOne do node
	if err != nil {
		log.Fatal(err)
	}

	linha, err := bd.Query("SELECT * FROM usuarios where id = ?", usuarioId) // parecido com o prepare statement
	if err != nil {
		w.Write([]byte("FAILED TO GET USUARIOS"))
		return
	}
	defer linha.Close()

	var usuarioRetornado modelos.Usuario
	if linha.Next() { // para cada linha executada ele faz uma iteracao
		if err := linha.Scan(&usuarioRetornado.Id, &usuarioRetornado.Nome, &usuarioRetornado.Nick, &usuarioRetornado.Email, &usuarioRetornado.Senha, &usuarioRetornado.CriadoEm); err != nil {
			w.Write([]byte("FAIL ON SCANNING USER"))
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	if err := json.NewEncoder(w).Encode(usuarioRetornado); err != nil {
		log.Fatal("Error retrieving usuarioRetornado")
	}
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode("USERS FOUND")
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode("USER FOUND")
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode("USER UPDATED")
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode("USER DELETED")
}
