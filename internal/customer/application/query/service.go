package query

import "github.com/Sadegh-kh/telemon/internal/customer/domain"

type CusotmerQueryService struct {
	repo domain.CustomerQueryRepository
}

func NewCustomerQueryService(repo domain.CustomerQueryRepository) *CusotmerQueryService {
	return &CusotmerQueryService{
		repo: repo,
	}
}
