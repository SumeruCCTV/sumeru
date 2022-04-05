package service

import "github.com/SumeruCCTV/sumeru/pkg/svcstat"

type Service interface {
	Name() string

	Start() error
	Stop() error

	svcstat.ServiceStats
}
