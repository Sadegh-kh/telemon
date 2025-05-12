package shared

import (
	"errors"
	"fmt"
)

type UsageUnit string

const (
	UsageMB UsageUnit = "MB"
	UsageGB UsageUnit = "GB"
)

type Usage struct {
	amount float64
	unit   UsageUnit
}

func NewUsage(amount float64, unit UsageUnit) (Usage, error) {
	if amount < 0 {
		return Usage{}, errors.New("usage amount cannot be negative")
	}
	if unit != UsageMB && unit != UsageGB {
		return Usage{}, fmt.Errorf("invalid usage unit: %s", unit)
	}
	return Usage{amount: amount, unit: unit}, nil
}

func (u Usage) String() string {
	return fmt.Sprintf("%.2f %s", u.amount, u.unit)
}

func (u Usage) ToMB() float64 {
	if u.unit == UsageGB {
		return u.amount * 1024
	}
	return u.amount
}
func (u Usage) ToGB() float64 {
	if u.unit == UsageMB {
		return u.amount / 1024
	}
	return u.amount
}
