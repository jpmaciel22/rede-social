package modelos

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	Id       uint64    `json:"id,omitempty"` // aqui no caso esse omitempty é para caso nao exista, nao passar 0, mas sim NAO PASSAR O ID.
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm"`
}

func (usuario *Usuario) Preparar() error {
	if err := usuario.validar(); err != nil {
		return err
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New(" O nome é obrigatório e não pode estar em branco")
	}
	if usuario.Nick == "" {
		return errors.New(" O nick é obrigatório e não pode estar em branco")
	}
	if usuario.Email == "" {
		return errors.New(" O email é obrigatório e não pode estar em branco")
	}
	if usuario.Senha == "" {
		return errors.New(" A senha é obrigatória e não pode estar em branco")
	}
	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
}
