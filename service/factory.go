package service

import (
	"github.com/drugs-4-3/ingrid_task/osrm"
	"sync"
)

var serviceImplOnce sync.Once
var serviceImpl *Service

func GetService() *Service {
	serviceImplOnce.Do(func() {
		serviceImpl = &Service{
			client: osrm.NewClient(),
		}
	})

	return serviceImpl
}
