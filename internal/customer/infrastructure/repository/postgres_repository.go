package repository

import (
	"context"
	"database/sql"

	"github.com/Sadegh-kh/telemon/internal/customer/domain"
)

type PostgresCustomerRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresCustomerRepository {
	return &PostgresCustomerRepository{db: db}
}

func (r *PostgresCustomerRepository) Create(ctx context.Context, c *domain.Customer) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO customers (id, name, phone_number, national_id, address, region_id, created_at, updated_at,is_deleted)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		c.ID,
		c.Name,
		c.PhoneNumber,
		c.NationalID,
		c.Address,
		c.RegionID,
		c.CreatedAt,
		c.UpdatedAt,
		false,
	)
	return err
}

func (r *PostgresCustomerRepository) Save(ctx context.Context, c *domain.Customer) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE customers
		SET name = $1, phone_number = $2, national_id = $3, address = $4, region_id = $5, updated_at = $6
		WHERE id = $7`,
		c.Name,
		c.PhoneNumber,
		c.NationalID,
		c.Address,
		c.RegionID,
		c.UpdatedAt,
		c.ID,
	)
	return err
}

func (r *PostgresCustomerRepository) Delete(ctx context.Context, id domain.CustomerID) error {
	const query = `UPDATE customers SET is_deleted = $1 WHERE id = $2`

	result, err := r.db.ExecContext(ctx, query, true, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *PostgresCustomerRepository) FindByID(ctx context.Context, id domain.CustomerID) (*domain.Customer, error) {
	var customer domain.Customer
	err := r.db.QueryRowContext(ctx, "SELECT id, national_id, name, phone_number, address, region_id, created_at, updated_at FROM customers WHERE id = $1 AND is_deleted = false", id).
		Scan(&customer.ID, &customer.NationalID, &customer.Name, &customer.PhoneNumber, &customer.Address, &customer.RegionID, &customer.CreatedAt, &customer.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *PostgresCustomerRepository) FindByNationalID(ctx context.Context, nID domain.NationalID) (*domain.Customer, error) {
	var customer domain.Customer
	err := r.db.QueryRowContext(ctx, "SELECT id, national_id, name, phone_number, address, region_id, created_at, updated_at FROM customers WHERE national_id = $1 AND is_deleted = false", nID).
		Scan(&customer.ID, &customer.NationalID, &customer.Name, &customer.PhoneNumber, &customer.Address, &customer.RegionID, &customer.CreatedAt, &customer.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
