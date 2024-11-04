package product

import "github.com/androsyz/inventory/config"

type Usecase struct {
	cfg          *config.Config
	repoProduct  repoProductInterface
	repoSupplier repoSupplierInterface
}

func New(cfg *config.Config, repoProduct repoProductInterface, repoSupplier repoSupplierInterface) *Usecase {
	return &Usecase{
		cfg:          cfg,
		repoProduct:  repoProduct,
		repoSupplier: repoSupplier,
	}
}
