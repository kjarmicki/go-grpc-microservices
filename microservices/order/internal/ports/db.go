package ports

import "github.com/kjarmicki/go-grpc-microservices/microservices/order/internal/application/core/domain"

type DbPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
}
