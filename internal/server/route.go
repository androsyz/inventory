package server

import (
	handler "github.com/androsyz/inventory/internal/handler"
	repoProduct "github.com/androsyz/inventory/internal/repository/product"
	repoSupplier "github.com/androsyz/inventory/internal/repository/supplier"
	repoTransaction "github.com/androsyz/inventory/internal/repository/transaction"
	repoUser "github.com/androsyz/inventory/internal/repository/user"
	ucProduct "github.com/androsyz/inventory/internal/usecase/product"
	ucSupplier "github.com/androsyz/inventory/internal/usecase/supplier"
	ucTransaction "github.com/androsyz/inventory/internal/usecase/transaction"
	ucUser "github.com/androsyz/inventory/internal/usecase/user"
)

func (s *Server) ConfigureRoutes() {

	repoSupplier := repoSupplier.New(s.Config, s.Sql)
	ucSupplier := ucSupplier.New(s.Config, repoSupplier)

	repoProduct := repoProduct.New(s.Config, s.Sql)
	ucProduct := ucProduct.New(s.Config, repoProduct, repoSupplier)

	repoUser := repoUser.New(s.Config, s.Sql)
	ucUser := ucUser.New(s.Config, repoUser)

	repoTransaction := repoTransaction.New(s.Config, s.Sql)
	ucTransaction := ucTransaction.New(s.Config, s.Sql, repoTransaction, repoProduct)

	handler := handler.NewHandler(ucSupplier, ucProduct, ucUser, ucTransaction, *s.Config)

	s.Echo.POST("/supplier", handler.CreateSupplier)
	s.Echo.GET("/supplier", handler.GetSuppliers)
	s.Echo.POST("/product", handler.CreateProduct)
	s.Echo.PUT("/product/:product_id", handler.UpdateProduct)
	s.Echo.GET("/product", handler.GetProducts)
	s.Echo.POST("/safety-stock", handler.GetSafetyStock)
	s.Echo.POST("/user", handler.CreateUser)
	s.Echo.GET("/user", handler.GetUsers)
	s.Echo.POST("/transaction", handler.CreateTransaction)
}
