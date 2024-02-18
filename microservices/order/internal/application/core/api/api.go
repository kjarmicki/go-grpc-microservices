package api

import (
	"github.com/kjarmicki/go-grpc-microservices/microservices/order/internal/application/core/domain"
	"github.com/kjarmicki/go-grpc-microservices/microservices/order/internal/ports"
)

type Application struct {
	db ports.DbPort
}

func NewApplication(db ports.DbPort) *Application {
	return &Application{
		db: db,
	}
}

func (a *Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}
