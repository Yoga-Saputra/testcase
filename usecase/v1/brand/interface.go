package brand

import (
	"github.com/Yoga-Saputra/testcase/internal/entity"
)

type Repository interface {
	Insert(a *entity.Brand) error
	FindByBrandID(brandID uint16) (e entity.Product, ok bool)
	Delete(brandID uint16) (row int64, err error)
}
