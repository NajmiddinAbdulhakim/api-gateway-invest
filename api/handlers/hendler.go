package hendlers

import (
	"github.com/NajmiddinAbdulhakim/iman/api-gateway/config"
	"github.com/NajmiddinAbdulhakim/iman/api-gateway/service"
)

type handler struct {
	cfg config.Config
	serviceManager service.IServiceManager
}

type HandlerConfig struct {
	Cfg config.Config
	ServiceManager service.IServiceManager
}

func New(c *HandlerConfig) *handler {
	return &handler{
		serviceManager: c.ServiceManager,
		cfg: c.Cfg,
	}
}