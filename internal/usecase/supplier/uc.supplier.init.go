package supplier

import "github.com/androsyz/inventory/config"

type Usecase struct {
	cfg          *config.Config
	repoSupplier repoSupplierInterface
}

func New(cfg *config.Config, repoSupplier repoSupplierInterface) *Usecase {
	return &Usecase{
		cfg:          cfg,
		repoSupplier: repoSupplier,
	}
}
