package domain

import "context"

type CustomerRepository interface {
	FindByID(ctx context.Context, id CustomerID) (*Customer, error)
	FindByNationalID(ctx context.Context, nID NationalID) (*Customer, error)
	Save(ctx context.Context, customer *Customer) error
	Create(ctx context.Context, customer *Customer) error
	Delete(ctx context.Context, id CustomerID) error
}
