package repo

import (
	"context"

	"github.com/mshsvnv/service/model"
)

type ISupplierRepository interface {
	Create(ctx context.Context, supplier *model.Supplier) error
	Update(ctx context.Context, supplier *model.Supplier) error
	Remove(ctx context.Context, email string) error
	GetSupplierByID(ctx context.Context, id int) (*model.Supplier, error)
	GetSupplierByEmail(ctx context.Context, email string) (*model.Supplier, error)
	GetAllSuppliers(ctx context.Context) ([]*model.Supplier, error)
}
