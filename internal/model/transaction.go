package model

type Transaction struct {
	ID         int `json:"id"`
	UserID     int `json:"user_id"`
	ProductID  int `json:"product_id"`
	TotalPrice int `json:"total_price"`
	Quantity   int `json:"quantity"`
}

type CreateTransactionReq struct {
	UserID    int `json:"user_id" validate:"required"`
	ProductID int `json:"product_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required"`
}

type TransactionRes struct {
	ID         int `json:"id"`
	UserID     int `json:"user_id"`
	ProductID  int `json:"product_id"`
	TotalPrice int `json:"total_price"`
	Quantity   int `json:"quantity"`
}
