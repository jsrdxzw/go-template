package svc

import (
	"example/be/service/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config, env string) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
