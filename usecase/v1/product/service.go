package product

import (
	"errors"
	"fmt"

	"github.com/Yoga-Saputra/testcase/internal/entity"
	"github.com/Yoga-Saputra/testcase/internal/helper"
	"github.com/Yoga-Saputra/testcase/internal/repo"
	"github.com/Yoga-Saputra/testcase/pkg/kemu"
	"github.com/jinzhu/copier"
)

// Service represent Republish Kafka services interface
type (
	Service struct {
		repo Repository
		kemu *kemu.Mutex
	}

	PProduct repo.PaginationArgs
)

// NewService creates new Republish Kafka services
func NewService(
	kemu *kemu.Mutex,
	r Repository,
	callback ...func(s string),
) *Service {
	if len(callback) > 0 {
		callback[0]("Registering Product List Domain Entity...")
	}

	svc := &Service{
		repo: r,
		kemu: kemu,
	}

	return svc
}

// CreateBrand create new Brand
func (s *Service) CreateProduct(
	brandID uint16,
	namaProduct string,
	harga float64,
	qty int16,
) error {

	a := &entity.Product{
		BrandID:     uint64(brandID),
		NamaProduct: namaProduct,
		Harga:       harga,
		Quantity:    qty,
		CreatedAt:   helper.LoadLocation(),
	}

	defer func() {
		a = nil
	}()

	if err := s.repo.Insert(a); err != nil {
		return err
	}

	return nil
}

// CreateBrand create new Brand
func (s *Service) UpdateProduct(
	productID uint64,
	brandID uint16,
	namaProduct string,
	harga float64,
	qty int16,
) (int64, error) {
	stmt := map[string]interface{}{
		"brand_id":     brandID,
		"nama_product": namaProduct,
		"harga":        harga,
		"quantity":     qty,
		"updated_at":   helper.LoadLocation(),
	}

	cond := map[string]interface{}{"id": productID}

	return s.repo.UpdateProduct(stmt, cond)
}

// GetAdjustmentListPagination adjustment's GetListPagination wrapper.
func (s *Service) GetProductListPagination(pp *PProduct) (
	countTotal int64,
	res []entity.ProductDataTable,
	err error,
) {
	if pp == nil {
		err = errors.New("arguments cannot be nil")
		return
	}

	if pp.Joins == nil {
		pp.Joins = map[interface{}][]interface{}{"Brand": nil}
	}
	total, x, err := s.repo.GetListProductPagination((*repo.PaginationArgs)(pp))
	if err != nil {
		return
	}

	countTotal = total
	copier.Copy(&res, &x)
	return
}

// FindByBrandId
func (s *Service) FindByBrandId(productID interface{}) (*entity.Product, int, error) {
	// Prepare query conditions
	conds := make(map[string]interface{})
	switch {
	case productID != nil:
		conds["id"] = productID

	default:
		return nil, 0, fmt.Errorf("argument productID cannot be empty -> productID: %d", productID)
	}

	brand, rows, err := s.repo.FindByProductID(conds)

	switch {
	// Always check error appear first
	case err != nil:
		return nil, rows, err

	case rows == 0:
		return nil, rows, nil

	default:
		return &brand, rows, nil
	}
}

// // delete brandID
func (s *Service) DeleteByProductID(productID uint64) error {

	return s.repo.Delete(productID)
}
