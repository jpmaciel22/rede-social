package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if err = usuario.Preparar("cadastro"); err != nil {
		respostas.ERRO(w, 400, err)
		return
	}

	bd, err := banco.Conectar()
	if err != nil {
		respostas.ERRO(w, 500, err)
		return
	}
	defer bd.Close()

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
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	bd, err := banco.Conectar()
	if err != nil {
		respostas.ERRO(w, 500, err)
		return
	}
	defer bd.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(bd) // so para poder ter os metodos no banco como "executador do metodo"
	usuarios, err := repositorio.Buscar(nomeOuNick)           // equivalente ao service createOne do node
	if err != nil {
		respostas.ERRO(w, 500, err)
		return
	}

	respostas.JSON(w, 201, usuarios)
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	userId, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		respostas.ERRO(w, 422, err)
		return
	}

	bd, err := banco.Conectar()
	if err != nil {
		respostas.ERRO(w, 500, err)
		return
	}
	defer bd.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(bd)
	usuario, err := repositorio.BuscarPorId(userId)
	if err != nil {
		respostas.ERRO(w, 500, err)
		return
	}

	respostas.JSON(w, 200, usuario)
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		respostas.ERRO(w, 400, err)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		respostas.ERRO(w, 400, err)
		return
	}

	var usuario modelos.Usuario

	if err = json.Unmarshal(body, &usuario); err != nil {
		respostas.ERRO(w, 400, err)
		return
	}

	if err = usuario.Preparar("edicao"); err != nil {
		respostas.ERRO(w, 400, err)
		return
	}

	bd, err := banco.Conectar()
	if err != nil {
		respostas.ERRO(w, 500, err)
		return
	}
	defer bd.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(bd)
	if err = repositorio.Atualizar(usuarioId, usuario); err != nil {
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

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		respostas.ERRO(w, 400, err)
		return
	}

	bd, err := banco.Conectar()
	if err != nil {
		respostas.ERRO(w, 500, err)
		return
	}
	defer bd.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(bd)
	if err = repositorio.Deletar(usuarioId); err != nil {
		respostas.ERRO(w, 500, err)
		return
	}

	respostas.JSON(w, 200, fmt.Sprintf("DELETANDO USUÁRIO DE ID: %d", usuarioId))
}
