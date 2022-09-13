package services

import (
	"campo-bombado/internal/core/domain"
	"campo-bombado/internal/core/ports"
	"campo-bombado/pkg/apperrors"
	"campo-bombado/pkg/uidgen"

	"github.com/matiasvarela/errors"
)

type service struct {
	gamesRepository ports.GameRepository
	uidgen          uidgen.UIDGen
}

func New(gamesRepository ports.GameRepository, uidgen uidgen.UIDGen) *service {
	return &service{
		gamesRepository: gamesRepository,
		uidgen:          uidgen,
	}
}

func (srv *service) Get(id string) (domain.Game, error) {
	game, err := srv.gamesRepository.Get(id)
	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			return domain.Game{}, errors.New(apperrors.NotFound, err, "game not found", "")
		}

		return domain.Game{}, errors.New(apperrors.Internal, err, "get game from repository has failed", "")
	}

	game.Board = game.Board.HideBombs()

	return game, nil
}

func (srv *service) Create(name string, size uint, bombs uint) (domain.Game, error) {
	if bombs >= size*size {
		return domain.Game{}, errors.New(apperrors.InvalidInput, nil, "the number of bombs is too high", "")
	}

	newGame := domain.NewGame(srv.uidgen.New(), name, size, bombs)

	if err := srv.gamesRepository.Save(newGame); err != nil {
		return domain.Game{}, errors.New(apperrors.Internal, err, "server failed to save the new game", "")
	}

	newGame.Board = newGame.Board.HideBombs()

	return newGame, nil
}

func (srv *service) RevealCell(id string, row uint, col uint) (domain.Game, error) {
	gameFound, err := srv.Get(id)

	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			return domain.Game{}, errors.New(apperrors.NotFound, err, "game not found", "")
		}

		return domain.Game{}, errors.New(apperrors.Internal, err, "get game from repository has failed", "")
	}

	if !gameFound.Board.IsValidPosition(row, col) {
		return domain.Game{}, errors.New(apperrors.InvalidInput, nil, "invalid position required", "")
	}

	if gameFound.IsOver() {
		return domain.Game{}, errors.New(apperrors.IllegalOperation, nil, "the game is over", "")
	}

	if gameFound.Board.Contains(row, col, domain.CELL_BOMB) {
		gameFound.State = domain.GAME_STATE_LOST
	} else {
		gameFound.Board.Set(row, col, domain.CELL_REVEALED)

		if !gameFound.Board.HasEmptyCells() {
			gameFound.State = domain.GAME_STATE_WON
		}
	}

	if err := srv.gamesRepository.Save(gameFound); err != nil {
		return domain.Game{}, errors.New(apperrors.Internal, err, "update game into repository has failed", "")
	}

	gameFound.Board = gameFound.Board.HideBombs()

	return gameFound, nil
}
