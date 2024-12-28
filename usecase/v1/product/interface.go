package product

import (
	"github.com/Yoga-Saputra/testcase/internal/entity"
	"github.com/Yoga-Saputra/testcase/internal/repo"
)

type Repository interface {
	// GetListProductPagination(pa *repo.PaginationArgs) (res repo.RepoDatatableResponse, err error)
	GetListProductPagination(pa *repo.PaginationArgs) (
		countTotal int64,
		res []entity.Product,
		// res repo.RepoDatatableResponse,
		err error,
	)
	FindByProductID(conds map[string]interface{}) (res entity.Product, row int, err error)
	UpdateProduct(
		stats map[string]interface{},
		conds map[string]interface{},
	) (int64, error)

	Insert(a *entity.Product) error
	Delete(productID uint64) error
}
