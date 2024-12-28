package app

import (
	"github.com/Yoga-Saputra/testcase/internal/repo"
	"github.com/Yoga-Saputra/testcase/pkg/kemu"
	"github.com/Yoga-Saputra/testcase/usecase/v1/brand"
	ucv1brandhttp "github.com/Yoga-Saputra/testcase/usecase/v1/brand/http"
	"github.com/Yoga-Saputra/testcase/usecase/v1/product"
	ucv1producthttp "github.com/Yoga-Saputra/testcase/usecase/v1/product/htpp"
)

// Helper table brand function will return entity repository that using gorm
func getRepoBrandGorm() *repo.BrandRepoDB {
	return repo.NewBrandRepoDB(DBA.DB, DBA.SQL)
}

// Helper table product function will return entity repository that using gorm
func getRepoProductGorm() *repo.ProductRepoDB {
	return repo.NewProductRepoDB(DBA.DB, DBA.SQL)
}

// DoEnV1Register register domain entity handler version 1 into the app
func doEntV1Register(args *AppArgs) {
	kemu := kemu.New()

	if HardMaintenance == "false" {
		printOutUp("Registering domain entity handler...")

		// Brand
		ucv1brandSvc := brand.NewService(
			kemu,
			getRepoBrandGorm(),
			printOutUp,
		)
		ucv1brandhttp.RegisterRoute(API.RouteGroup["v1"], *ucv1brandSvc, kemu)

		// Product
		ucv1ProductSvc := product.NewService(
			kemu,
			getRepoProductGorm(),
			printOutUp,
		)
		ucv1producthttp.RegisterRoute(API.RouteGroup["v1"], *ucv1ProductSvc, kemu)
	}
}
