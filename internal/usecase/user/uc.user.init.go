package user

import "github.com/androsyz/inventory/config"

type Usecase struct {
	cfg      *config.Config
	repoUser repoUserInterface
}

func New(cfg *config.Config, repoUser repoUserInterface) *Usecase {
	return &Usecase{
		cfg:      cfg,
		repoUser: repoUser,
	}
}
