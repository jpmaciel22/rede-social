package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		respostas.ERRO(w, 422, err)
		return
	}

	var usuario modelos.Usuario

	if err = json.Unmarshal(body, &usuario); err != nil {
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
	usuarioAchado, err := repositorio.BuscarPorEmail(usuario.Email)
	if err != nil {
		respostas.ERRO(w, 500, err)
		return
	}

	if err = seguranca.VerificarSenha(usuarioAchado.Senha, usuario.Senha); err != nil {
		respostas.ERRO(w, 401, err) // senha errada
		return
	}

	token, err := autenticacao.CriarToken(usuario.Id)
	if err != nil {
		respostas.ERRO(w, 500, err)
	}
	respostas.JSON(w, 200, token)

}
