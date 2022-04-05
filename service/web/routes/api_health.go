package routes

import (
	"fmt"
	"github.com/SumeruCCTV/sumeru"
	"github.com/SumeruCCTV/sumeru/pkg/svcstat"
	"github.com/SumeruCCTV/sumeru/service/web"
	"github.com/gofiber/fiber/v2"
	"runtime"
)

type _svcStatus struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type healthResponse struct {
	Services []_svcStatus `json:"services"`
	Mem      string       `json:"mem"`
	Next     string       `json:"next"`
}

func init() {
	web.Register(func(svc *web.Service, app *fiber.App) {
		app.Get("/health", func(ctx *fiber.Ctx) error {
			res := new(healthResponse)

			for name, s := range sumeru.App.Services() {
				if status, ok := s.(svcstat.ServiceStats); ok {
					res.Services = append(res.Services, _svcStatus{
						Name:   name,
						Status: string(status.Status()),
					})
				}
			}

			var mem runtime.MemStats
			runtime.ReadMemStats(&mem)
			res.Mem = byteCount(mem.Alloc)
			res.Next = byteCount(mem.NextGC)

			return ctx.JSON(res)
		})
	})
}

// Thanks: https://programming.guide/go/formatting-byte-size-to-human-readable-format.html
func byteCount(b uint64) string {
	const unit uint64 = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := unit, 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}
