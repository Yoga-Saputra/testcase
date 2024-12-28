package brand

import (
	"github.com/Yoga-Saputra/testcase/internal/entity"
	"github.com/Yoga-Saputra/testcase/internal/helper"
	"github.com/Yoga-Saputra/testcase/pkg/kemu"
)

// Service represent Republish Kafka services interface
type Service struct {
	repo Repository
	kemu *kemu.Mutex
}

// NewService creates new Republish Kafka services
func NewService(
	kemu *kemu.Mutex,
	r Repository,
	callback ...func(s string),
) *Service {
	if len(callback) > 0 {
		callback[0]("Registering Brand List Domain Entity...")
	}

	svc := &Service{
		repo: r,
		kemu: kemu,
	}

	return svc
}

// CreateBrand create new Brand
func (s *Service) CreateBrand(BrandName string) error {

	a := &entity.Brand{
		NamaBrand: BrandName,
		CreatedAt: helper.LoadLocation(),
	}

	defer func() {
		a = nil
	}()

	if err := s.repo.Insert(a); err != nil {
		return err
	}

	return nil
}

// FindByBrandId
func (s *Service) FindByBrandId(BrandID uint16) bool {

	_, ok := s.repo.FindByBrandID(BrandID)

	if ok {
		return ok
	}

	return false
}

// delete brandID
func (s *Service) DeleteByBrandID(BrandID uint16) (row int64, err error) {

	return s.repo.Delete(BrandID)
}
