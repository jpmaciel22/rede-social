package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		respostas.ERRO(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario modelos.Usuario
	if err = json.Unmarshal(body, &usuario); err != nil {
		respostas.ERRO(w, http.StatusBadRequest, err)
		return
	}

	bd, err := banco.Conectar()
	if err != nil {
		respostas.ERRO(w, 500, err)
		return
	}

	repositorio := repositorios.NovoRepositorioDeUsuarios(bd) // so para poder ter os metodos no banco como "executador do metodo"
	usuarioId, err := repositorio.Criar(usuario)              // equivalente ao service createOne do node
	if err != nil {
		respostas.ERRO(w, 500, err)
		return
	}

	linha, err := bd.Query("SELECT * FROM usuarios where id = ?", usuarioId) // parecido com o prepare statement
	if err != nil {
		respostas.ERRO(w, 500, err)
		return
	}
	defer linha.Close()

	usuario.Id = usuarioId

	respostas.JSON(w, 201, usuario)
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
