package repo

import (
	"context"

	"github.com/mshsvnv/service/model"
)

//go:generate mockery --name=IRacketrepo
type IRacketRepository interface {
	Create(ctx context.Context, racket *model.Racket) error
	Update(ctx context.Context, racket *model.Racket) error
	Remove(ctx context.Context, id int) error
	GetRacketByID(ctx context.Context, id int) (*model.Racket, error)
	GetAllAvaliableRackets(ctx context.Context) ([]*model.Racket, error)
	GetAllRackets(ctx context.Context) ([]*model.Racket, error)
}
