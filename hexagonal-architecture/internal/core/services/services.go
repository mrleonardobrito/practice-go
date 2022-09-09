package services

import (
	"campo-bombado/internal/core/domain"
	"campo-bombado/internal/core/ports"
	"errors"
)

type service struct {
	gamesRepository ports.GameRepository
}

func New(gamesRepository ports.GameRepository) *service {
	return &service{
		gamesRepository: gamesRepository,
	}
}

func (srv *service) Get(id string) (domain.Game, error) {
	game, err := srv.gamesRepository.Get(id)

	if err != nil {
		return domain.Game{}, errors.New("get game from a repository failed")
	}

	return game, nil
}

// func (srv *service) Create() (domain.Game, error) {

// }
