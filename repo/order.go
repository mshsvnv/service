package repo

import (
	"context"

	"github.com/mshsvnv/service/model"
)

//go:generate mockery --name=IOrderrepo
type IOrderRepository interface {
	Create(ctx context.Context, order *model.Order) error
	Update(ctx context.Context, order *model.Order) error
	Remove(ctx context.Context, orderID int) error
	GetAllInProgressOrders(ctx context.Context) ([]*model.Order, error)
	GetAllOrders(ctx context.Context) ([]*model.Order, error)
	GetMyOrders(ctx context.Context, userID int) ([]*model.Order, error)
	GetOrderByID(ctx context.Context, orderID int) (*model.Order, error)
}
