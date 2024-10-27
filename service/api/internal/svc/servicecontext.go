package svc

import (
	"example/be/service/app"
)

type ServiceContext struct {
	Config app.Config
}

func NewServiceContext(c app.Config, env string) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
