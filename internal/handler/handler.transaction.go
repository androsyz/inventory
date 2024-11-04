package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/androsyz/inventory/consts"
	"github.com/androsyz/inventory/internal/model"
	zlog "github.com/androsyz/inventory/internal/pkg/log"
	"github.com/androsyz/inventory/internal/response"
	"github.com/androsyz/inventory/internal/validator"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateTransaction(c echo.Context) error {
	ctx := context.Background()
	payload := new(model.CreateTransactionReq)

	if err := c.Bind(payload); err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_BIND, err))
		return response.ErrorResponse(c, err.Error(), http.StatusBadRequest)
	}

	if err := validator.ValidateStruct(payload); err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_VALIDATE_BODY, err))
		return response.ErrorResponse(c, err.Error(), http.StatusBadRequest)
	}

	res, err := h.ucTransaction.CreateTransaction(ctx, payload)
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf("error when call CreateTransaction, got %v", err))
		return response.ErrorResponse(c, err.Error(), http.StatusInternalServerError)
	}

	return response.SuccessResponse(c, res)
}
