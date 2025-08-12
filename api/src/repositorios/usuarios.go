package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

// cria um repositorio de usuarios no banco
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, err := repositorio.db.Prepare("INSERT INTO usuarios (nome, nick, email, senha) values (?,?,?,?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	usuarioInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(usuarioInserido), nil
}

func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%
	linhas, err := repositorio.db.Query("SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE nome OR nick LIKE ?", nomeOuNick)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if err = linhas.Scan(&usuario.Id, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil

}

func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linhas, err := repositorio.db.Query("SELECT id, senha from usuarios where email = ?", email)
	if err != nil {
		return modelos.Usuario{}, err
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if err = linhas.Scan(&usuario.Id, &usuario.Senha); err != nil {
			return modelos.Usuario{}, err
		}
	}
	return usuario, nil
}

func (repositorio Usuarios) BuscarPorId(userId uint64) (modelos.Usuario, error) {
	linhas, err := repositorio.db.Query("SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE id = ?", userId)
	if err != nil {
		return modelos.Usuario{}, err
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if err = linhas.Scan(&usuario.Id, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); err != nil {
			return modelos.Usuario{}, err
		}
	}
	return usuario, nil
}

func (repositorio Usuarios) Atualizar(userId uint64, usuario modelos.Usuario) error {
	statement, err := repositorio.db.Prepare("update usuarios set nome = ? , nick = ? , email = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, userId); err != nil {
		return err
	}
	return nil
}

func (repositorio Usuarios) Deletar(userId uint64) error {
	statement, err := repositorio.db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userId); err != nil {
		return err
	}
	return nil
}
