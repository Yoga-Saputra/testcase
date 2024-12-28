package ucv1brandhttp

import (
	"github.com/Yoga-Saputra/testcase/pkg/kemu"
	"github.com/Yoga-Saputra/testcase/usecase/v1/brand"
	"github.com/labstack/echo/v4"
)

type domainService struct {
	s    brand.Service
	kemu *kemu.Mutex
}

func RegisterRoute(v1 *echo.Group, s brand.Service, k *kemu.Mutex) {
	// Setup domain service
	ds := &domainService{
		s:    s,
		kemu: k,
	}

	// Create root brand group
	bg := v1.Group("/brand")            // <- Route group (and also prefix) "brand"
	bg.POST("/create", ds.Create)       // <- create brand
	bg.DELETE("/delete/:id", ds.Delete) // <- delete brand
}
