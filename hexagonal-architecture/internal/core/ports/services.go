package ports

import "campo-bombado/internal/core/domain"

// porta que irá servir a um ator primário. Ex: "uma api que tentar entrar no sistema"
type GameService interface {
	Get(id string) error
	Create(name string, size uint, bombs uint) (domain.Game, error)
	Reveal(id string, row uint, col uint) (domain.Game, error)
}
