package domain

import "context"

type CustomerCommandRepository interface {
	Save(ctx context.Context, customer *Customer) error
	Create(ctx context.Context, customer *Customer) error
	Delete(ctx context.Context, id CustomerID) error
}

type CustomerQueryRepository interface {
	FindByID(ctx context.Context, id CustomerID) (*Customer, error)
	FindByNationalID(ctx context.Context, nID NationalID) (*Customer, error)
}
