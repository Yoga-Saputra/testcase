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
	BrandRepoDB struct {
		db  *gorm.DB
		sql *sql.DB
	}
)

// NewProductRepoDB create new DB repo for Product entity.
func NewBrandRepoDB(db *gorm.DB, sql *sql.DB) *BrandRepoDB {
	if db != nil {
		return &BrandRepoDB{
			db:  db,
			sql: sql,
		}
	}

	return nil
}

func (brd *BrandRepoDB) Insert(a *entity.Brand) error {
	if a == nil {
		return errors.New("pointer argument cannot be nil")
	}

	// Insert
	if err := brd.db.Omit(
		"UpdatedAt",
	).Create(a).Error; err != nil {
		switch {
		case strings.Contains(err.Error(), "SQLSTATE 23505"):
			err = fmt.Errorf("brand [%s] is already exists duplicate", a.NamaBrand)
		}

		return err
	}

	return nil
}

// FindByBrandID to select product by brandID.
func (brd *BrandRepoDB) FindByBrandID(brandID uint16) (e entity.Product, ok bool) {
	// Query
	result := brd.db.Where("brand_id = ?", brandID).First(&e)
	if result.RowsAffected <= 0 {
		return
	}

	// Finish
	ok = true
	return
}

// delete log internal transaction
func (brd *BrandRepoDB) Delete(brandID uint16) (
	row int64,
	err error,
) {

	// Prepare transaction session
	tx := brd.db.Delete(&entity.Brand{}, brandID)
	row = tx.RowsAffected
	err = tx.Error

	return
}
