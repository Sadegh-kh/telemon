package domain

import (
	"time"

	customer_domain "github.com/Sadegh-kh/telemon/internal/customer/domain"
	serviceplan_domain "github.com/Sadegh-kh/telemon/internal/serviceplan/domain"
	"github.com/Sadegh-kh/telemon/internal/shared"
)

type ServiceID string

type ServiceStatus string

const (
	ServiceActive  ServiceStatus = "ACTIVE"
	ServicePaused  ServiceStatus = "PAUSED"
	ServiceExpired ServiceStatus = "EXPIRED"
)

type InternetService struct {
	ID         ServiceID
	CustomerID customer_domain.CustomerID
	PlanID     serviceplan_domain.PlanID
	Status     ServiceStatus
	Usage      shared.Usage
	StartDate  time.Time
	ExpireDate time.Time
}

func (s InternetService) IsExpired() bool {
	return time.Now().After(s.ExpireDate)
}
