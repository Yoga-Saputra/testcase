package ucv1producthttp

import (
	"github.com/Yoga-Saputra/testcase/pkg/kemu"
	"github.com/Yoga-Saputra/testcase/usecase/v1/product"
	"github.com/labstack/echo/v4"
)

type domainService struct {
	s    product.Service
	kemu *kemu.Mutex
}

func RegisterRoute(v1 *echo.Group, s product.Service, k *kemu.Mutex) {
	// Setup domain service
	ds := &domainService{
		s:    s,
		kemu: k,
	}

	// Create root product group
	bg := v1.Group("/product")          // <- Route group (and also prefix) "product"
	bg.GET("/list", ds.List)            // <- create product
	bg.POST("/create", ds.Create)       // <- create product
	bg.PUT("/update/:id", ds.Update)    // <- create product
	bg.DELETE("/delete/:id", ds.Delete) // <- delete product

}
