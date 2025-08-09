package modelos

import "time"

type Usuario struct {
	Id       uint64    `json:"id,omitempty"` // aqui no caso esse omitempty é para caso nao exista, nao passar 0, mas sim NAO PASSAR O ID.
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm"`
}
