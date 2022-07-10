package camera

import "github.com/SumeruCCTV/sumeru/service/database/models"

type Connector interface {
	TestConnection() error
}

type ConnectorData struct {
	uuid        string
	ipAddress   string
	port        int
	credentials models.CameraCredentials
}

func NewConnector(cameraType models.CameraType, svc *Service, data *ConnectorData) Connector {
	switch cameraType {
	case models.CameraTypeONVIF:
		return NewONVIFConnector(svc, data)
	case models.CameraTypeRTSP:
		return NewRTSPConnector(svc, data)
	default:
		return nil
	}
}
