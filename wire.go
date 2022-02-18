//go:build wireinject
// +build wireinject

package main

import (
	"gochicoba/handler"
	"gochicoba/repository"
	"gochicoba/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func BuyHandler(db *gorm.DB) handler.BuyHandler {
	wire.Build(repository.NewBuyRepository, service.NewBuyService, handler.NewBuyHandler)
	return handler.BuyHandler{}
}
