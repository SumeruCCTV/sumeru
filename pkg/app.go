package pkg

import (
	"fmt"
	"github.com/SumeruCCTV/sumeru/pkg/config"
	"github.com/SumeruCCTV/sumeru/pkg/constants"
	"github.com/SumeruCCTV/sumeru/pkg/utils"
	"github.com/SumeruCCTV/sumeru/service"
	"os"
	"os/signal"
)

type Application struct {
	services    map[string]service.Service
	injectables map[string]interface{}
	cfg         *config.Config
	log         *utils.Logger
}

func New(cfg *config.Config, logger *utils.Logger) *Application {
	return &Application{
		services:    make(map[string]service.Service),
		injectables: make(map[string]interface{}),
		cfg:         cfg,
		log:         logger,
	}
}

func (app *Application) AddService(svc service.Service) {
	app.services[svc.Name()] = svc
}

func (app *Application) Services() map[string]service.Service {
	return app.services
}

func (app *Application) Inject(name string, i interface{}) {
	app.injectables[name] = i
}

func (app *Application) Run() error {
	app.log.Infof("Starting Sumeru version %s", constants.SumeruVersion)
	app.Inject("cfg", app.cfg) // inject config
	for name, svc := range app.services {
		app.Inject("log", app.createLogger(svc))
		if err := app.injectFields(name, svc); err != nil {
			return fmt.Errorf("failed to inject fields: %v", err)
		}
		app.log.Infof("Starting %s service", name)
		if err := svc.Start(); err != nil {
			return err
		}
		app.log.Infof("Started %s service", name)
	}

	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt, os.Kill)

	app.log.Info("Press CTRL+C to exit.")
	<-ch // wait...
	app.log.Info("Shutting down gracefully...")

	for name, svc := range app.services {
		app.log.Infof("Stopping %s service", name)
		if err := svc.Stop(); err != nil {
			return err
		}
		app.log.Infof("Stopped %s service", name)
	}
	return nil
}
