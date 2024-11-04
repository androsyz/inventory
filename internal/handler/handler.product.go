package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/androsyz/inventory/consts"
	"github.com/androsyz/inventory/internal/model"
	zlog "github.com/androsyz/inventory/internal/pkg/log"
	"github.com/androsyz/inventory/internal/response"
	"github.com/androsyz/inventory/internal/validator"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateProduct(c echo.Context) error {
	ctx := context.Background()
	payload := new(model.CreateProductReq)

	if err := c.Bind(payload); err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_BIND, err))
		return response.ErrorResponse(c, err.Error(), http.StatusBadRequest)
	}

	if err := validator.ValidateStruct(payload); err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_VALIDATE_BODY, err))
		return response.ErrorResponse(c, err.Error(), http.StatusBadRequest)
	}

	res, err := h.ucProduct.CreateProduct(ctx, payload)
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf("error when call CreateProduct, got %v", err))
		return response.ErrorResponse(c, err.Error(), http.StatusInternalServerError)
	}

	return response.SuccessResponse(c, res)
}

func (h *Handler) GetProducts(c echo.Context) error {
	ctx := context.Background()
	res, err := h.ucProduct.GetProducts(ctx)
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf("error when call GetProduct, got %v", err))
		return response.ErrorResponse(c, err.Error(), http.StatusInternalServerError)
	}

	return response.SuccessResponse(c, res)
}

func (h *Handler) UpdateProduct(c echo.Context) error {
	ctx := context.Background()

	productIDStr := c.Param("product_id")
	productID, _ := strconv.Atoi(productIDStr)

	payload := new(model.UpdateProductReq)
	payload.ID = productID

	if err := c.Bind(payload); err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_BIND, err))
		return response.ErrorResponse(c, err.Error(), http.StatusBadRequest)
	}

	if err := validator.ValidateStruct(payload); err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_VALIDATE_BODY, err))
		return response.ErrorResponse(c, err.Error(), http.StatusBadRequest)
	}

	err := h.ucProduct.UpdateProduct(ctx, payload)
	if err != nil && err.Error() == consts.ERR_PRODUCT_NOT_FOUND {
		zlog.Error(ctx, nil, fmt.Sprintf("product not found got %v", err))
		return response.ErrorResponse(c, err.Error(), http.StatusNotFound)
	}

	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf("error when call UpdateProduct, got %v", err))
		return response.ErrorResponse(c, err.Error(), http.StatusInternalServerError)
	}

	return response.SuccessResponse(c, nil)
}

func (h *Handler) GetSafetyStock(c echo.Context) error {
	ctx := context.Background()
	payload := new(model.SafetyStockReq)

	if err := c.Bind(payload); err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_BIND, err))
		return response.ErrorResponse(c, err.Error(), http.StatusBadRequest)
	}

	if err := validator.ValidateStruct(payload); err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_VALIDATE_BODY, err))
		return response.ErrorResponse(c, err.Error(), http.StatusBadRequest)
	}

	res, err := h.ucProduct.GetSafetyStock(ctx, payload)
	if err != nil && err.Error() == consts.ERR_PRODUCT_NOT_FOUND {
		zlog.Error(ctx, nil, fmt.Sprintf("product not found got %v", err))
		return response.ErrorResponse(c, err.Error(), http.StatusNotFound)
	}

	if err != nil && err.Error() == consts.ERR_SUPPLIER_NOT_FOUND {
		zlog.Error(ctx, nil, fmt.Sprintf("supplier not found got %v", err))
		return response.ErrorResponse(c, err.Error(), http.StatusNotFound)
	}

	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf("error when call GetSafetyStock, got %v", err))
		return response.ErrorResponse(c, err.Error(), http.StatusInternalServerError)
	}

	return response.SuccessResponse(c, res)
}
