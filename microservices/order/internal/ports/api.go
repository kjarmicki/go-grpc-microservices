package ports

import "github.com/kjarmicki/go-grpc-microservices/microservices/order/internal/application/core/domain"

type ApiPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
