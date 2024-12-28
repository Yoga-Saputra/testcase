package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Yoga-Saputra/testcase/internal/entity"
	"gorm.io/gorm"
)

type (
	ProductRepoDB struct {
		db  *gorm.DB
		sql *sql.DB
	}
)

// NewProductRepoDB create new DB repo for Product entity.
func NewProductRepoDB(db *gorm.DB, sql *sql.DB) *ProductRepoDB {
	if db != nil {
		return &ProductRepoDB{
			db:  db,
			sql: sql,
		}
	}

	return nil
}
func (prd *ProductRepoDB) FindByProductID(conds map[string]interface{}) (res entity.Product, rows int, err error) {
	// Find product records by given conditions and return product.
	tx := prd.db.Find(&res, conds)
	rows = int(tx.RowsAffected)
	err = tx.Error

	return
}

// GetListProductPagination select product datas with paginate.
func (prd *ProductRepoDB) GetListProductPagination(pa *PaginationArgs) (
	countTotal int64,
	res []entity.Product,
	err error,
) {

	// [1] Query of count total
	err = prd.db.Model(&entity.Product{}).
		Select("COUNT(1) as count").
		Count(&countTotal).
		Error
	if err != nil {
		return
	}

	// [2] Query of adjustment datas
	tx := prd.db.Model(&entity.Product{})

	// [2.1] Set-up query join
	for q, a := range pa.Joins {
		tx = tx.Joins(fmt.Sprintf("%v", q), a...)
	}

	// [2.2] Set-up query limit
	tx = tx.Limit(int(pa.Limit)).Offset(int(pa.Offset))

	// [3.] Set-up query order
	err = tx.Find(&res).Error
	if err != nil {
		return
	}
	return
}

// insert product
func (prd *ProductRepoDB) Insert(a *entity.Product) error {
	if a == nil {
		return errors.New("pointer argument cannot be nil")
	}

	// Insert
	if err := prd.db.Omit(
		"UpdatedAt",
	).Create(a).Error; err != nil {
		switch {
		case strings.Contains(err.Error(), "SQLSTATE 23505"):
			err = fmt.Errorf("product [%s] is already exists duplicate", a.NamaProduct)

		case strings.Contains(err.Error(), "SQLSTATE 23503"):
			err = fmt.Errorf("brand id [%v] doesnot exist", a.BrandID)
		}

		return err
	}

	return nil
}

// UpdateProduct by given update statements and conditions.
func (prd *ProductRepoDB) UpdateProduct(
	stats map[string]interface{},
	conds map[string]interface{},
) (int64, error) {
	// Update
	tx := prd.db.Model(&entity.Product{}).Where(conds).Updates(stats)

	// Check error
	if tx.Error != nil {
		switch {
		case strings.Contains(tx.Error.Error(), "SQLSTATE 23503"):
			return 0, fmt.Errorf("brand id [%v] doesnot exist", stats["brand_id"])
		default:
			return 0, tx.Error
		}
	}

	return tx.RowsAffected, nil
}

// delete product
func (prd *ProductRepoDB) Delete(brandID uint64) error {

	return prd.db.Delete(&entity.Product{}, brandID).Error
}
