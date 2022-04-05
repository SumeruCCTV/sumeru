package svcstat

type Status string

const (
	StatusHealthy   Status = "healthy"
	StatusUnhealthy Status = "unhealthy"
)

type ServiceStats interface {
	Status() Status
}
