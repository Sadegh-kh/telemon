package domain

import "github.com/Sadegh-kh/telemon/internal/shared"

type PlanID string

type ServiceType string

const (
	ServiceADSL  ServiceType = "ADSL"
	ServiceVDSL  ServiceType = "VDSL"
	ServiceFiber ServiceType = "FIBER"
)

type ServicePlan struct {
	ID        PlanID
	Name      string
	Type      ServiceType
	Bandwidth shared.BandWidth
	BasePrice int
}
