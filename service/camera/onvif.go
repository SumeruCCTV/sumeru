package camera

import (
	"fmt"
	"github.com/SumeruCCTV/go-onvif"
	"github.com/SumeruCCTV/sumeru/pkg/utils"
)

// maybe switch to https://github.com/videonext/onvif
// go-onvif is a mess

const protocol = "http"

type ONVIFConnector struct {
	svc *Service
	log *utils.Logger

	data *ConnectorData
	rtsp *RTSPConnector
}

func (c *ONVIFConnector) TestConnection() error {
	dvc := c.device()
	if _, err := dvc.GetInformation(); err != nil {
		return err
	}
	return nil
}

func (c *ONVIFConnector) device() *onvif.Device {
	return &onvif.Device{
		XAddr:    fmt.Sprintf("%s://%s:%d/onvif/services", protocol, c.data.ipAddress, c.data.port),
		User:     c.data.credentials.Username,
		Password: c.data.credentials.Password,
	}
}

func NewONVIFConnector(svc *Service, data *ConnectorData) *ONVIFConnector {
	return &ONVIFConnector{
		svc:  svc,
		log:  svc.log.Named("onvif"),
		data: data,
	}
}
