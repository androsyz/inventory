package handler

import (
	"github.com/androsyz/inventory/config"
)

type Handler struct {
	ucSupplier    ucSupplierInterface
	ucProduct     ucProductInterface
	ucUser        ucUserInterface
	ucTransaction ucTransactionInterface
	cfg           config.Config
}

func NewHandler(
	ucSupplier ucSupplierInterface,
	ucProduct ucProductInterface,
	ucUser ucUserInterface,
	ucTransaction ucTransactionInterface,
	cfg config.Config,
) *Handler {
	return &Handler{
		ucSupplier:    ucSupplier,
		ucProduct:     ucProduct,
		ucUser:        ucUser,
		ucTransaction: ucTransaction,
		cfg:           cfg}
}
