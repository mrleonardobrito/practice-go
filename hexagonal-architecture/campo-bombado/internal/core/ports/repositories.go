package ports

import (
	"campo-bombado/internal/core/domain"
)

// porta que irá servir a um ator secundário. Ex: "preciso salvar o jogo em um banco de dados"
type GamesRepository interface {
	Get(id string) (domain.Game, error)
	Save(game domain.Game) error
}
