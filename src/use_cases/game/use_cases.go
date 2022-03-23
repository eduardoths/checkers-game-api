package game

import (
	"github.com/eduardoths/checkers-game/src/interfaces"
	"github.com/eduardoths/checkers-game/src/repositories"
)

type GameUseCases struct {
	repository interfaces.GameRepository
}

func NewGameUseCases(repos repositories.RepositoriesContainer) *GameUseCases {
	return &GameUseCases{
		repository: repos.GameRepository,
	}
}
