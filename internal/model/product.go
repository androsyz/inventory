package model

type Product struct {
	ID         int    `json:"id"`
	SupplierID int    `json:"supplier_id"`
	Name       string `json:"name"`
	Stock      int    `json:"stock"`
	Price      int    `json:"price"`
}

type CreateProductReq struct {
	SupplierID int    `json:"supplier_id" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Stock      int    `json:"stock" validate:"required"`
	Price      int    `json:"price" validate:"required"`
}

type UpdateProductReq struct {
	ID         int    `json:"id" validate:"required"`
	SupplierID int    `json:"supplier_id" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Stock      int    `json:"stock" validate:"required"`
	Price      int    `json:"price" validate:"required"`
}

type SafetyStockReq struct {
	ProductID  int `json:"product_id" validate:"required"`
	AverageReq int `json:"average_req" validate:"required"`
}

type ProductRes struct {
	ID         int    `json:"id"`
	SupplierID int    `json:"supplier_id"`
	Name       string `json:"name"`
	Stock      int    `json:"stock"`
	Price      int    `json:"price"`
}

type ProductListRes struct {
	Products []*ProductRes `json:"products"`
}

type SafetyStockRes struct {
	ProductID   int    `json:"product_id"`
	ProductName string `json:"product_name"`
	Stock       int    `json:"stock"`
	SafetyStock int    `json:"safety_stock"`
}
