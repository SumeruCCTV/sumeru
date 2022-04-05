package camera

import (
	"github.com/SumeruCCTV/sumeru/pkg/svcstat"
	"github.com/SumeruCCTV/transcoder/ffmpeg"
)

type Service struct {
}

func (Service) Name() string {
	return "camera"
}

func (svc *Service) Start() error {
	_ = ffmpeg.Config{}

	return nil
}

func (svc *Service) Stop() error {
	return nil
}

func (svc *Service) Status() svcstat.Status {
	return svcstat.StatusHealthy
}
