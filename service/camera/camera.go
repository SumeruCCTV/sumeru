package camera

import (
	"github.com/SumeruCCTV/sumeru/pkg/svcstat"
	"github.com/SumeruCCTV/sumeru/pkg/utils"
	"github.com/SumeruCCTV/sumeru/service/database"
	"github.com/SumeruCCTV/sumeru/service/database/models"
)

type Service struct {
	log      *utils.Logger
	database *database.Service

	// queue to add cameras
	q      chan *models.Camera
	closed bool

	// map of currently active cameras by uuid
	cameras map[string]Connector
}

func (Service) Name() string {
	return "camera"
}

func (svc *Service) Start() error {
	svc.q = make(chan *models.Camera)
	go svc.startQueue()

	svc.cameras = make(map[string]Connector)
	return nil
}

func (svc *Service) Stop() error {
	close(svc.q)
	return nil
}

func (svc *Service) Status() svcstat.Status {
	if !svc.closed {
		return svcstat.StatusHealthy
	}
	return svcstat.StatusUnhealthy
}

func (svc *Service) AddCamera(cam *models.Camera) {
	svc.q <- cam
}

func (svc *Service) startQueue() {
	svc.closed = false
	for {
		cam, ok := <-svc.q
		if !ok {
			svc.closed = true
			return
		}
		go svc.accept(cam)
	}
}
