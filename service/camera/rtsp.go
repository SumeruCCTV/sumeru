package camera

import (
	"fmt"
	"github.com/SumeruCCTV/sumeru/pkg/utils"
	"github.com/SumeruCCTV/transcoder/ffmpeg"
	"net"
)

// TODO: see https://github.com/andrewlfw/joy4/tree/main/format/rtspv2
// maybe use this instead of ffmpeg?

type RTSPConnector struct {
	svc *Service
	log *utils.Logger

	data *ConnectorData
}

func (c *RTSPConnector) TestConnection() error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.data.ipAddress, c.data.port))
	if err != nil {
		return err
	}
	_ = conn.Close()

	_ = ffmpeg.Config{}

	return nil
}

func NewRTSPConnector(svc *Service, data *ConnectorData) *RTSPConnector {
	return &RTSPConnector{
		svc:  svc,
		log:  svc.log.Named("rtsp"),
		data: data,
	}
}
