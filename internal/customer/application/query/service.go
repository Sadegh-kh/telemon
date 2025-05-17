package query

import "github.com/Sadegh-kh/telemon/internal/customer/domain"

type CusotmerQueryService struct {
	repo domain.CustomerRepository
}

func NewCustomerQueryService(repo domain.CustomerRepository) *CusotmerQueryService {
	return &CusotmerQueryService{
		repo: repo,
	}
}
