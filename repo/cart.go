package repo

import (
	"context"

	"github.com/mshsvnv/service/dto"
	"github.com/mshsvnv/service/model"
)

//go:generate mockery --name=ICartrepo
type ICartRepository interface {
	Create(ctx context.Context, cart *model.Cart) error
	Update(ctx context.Context, cart *model.Cart) error
	Remove(ctx context.Context, userID int) error
	AddRacket(ctx context.Context, req *dto.AddRacketCartReq) error
	RemoveRacket(ctx context.Context, req *dto.RemoveRacketCartReq) error
	GetCartByID(ctx context.Context, userID int) (*model.Cart, error)
}
