package entity

import "time"

type (
	Product struct {
		ID          uint64    `gorm:"primaryKey" json:"id"`
		BrandID     uint64    `gorm:"not null;" json:"BrandId"`
		Brand       Brand     `gorm:"foreignKey:BrandID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"brand"`
		NamaProduct string    `gorm:"not null;size:25;uniqueIndex:product_brand_id_idx;" json:"namaProduct"`
		Harga       float64   `gorm:"type:numeric(30,6);default:0;not null" json:"harga"`
		Quantity    int16     `gorm:"not null;" json:"quantity"`
		CreatedAt   time.Time `gorm:"not null" json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
	}

	ProductDataTable struct {
		ID uint64 `json:"id" copier:"must"`

		BrandID     uint64  `json:"brand_id" copier:"must"`
		NamaProduct string  `json:"nama_product" copier:"must"`
		Harga       float64 `json:"harga" copier:"must"`
		Quantity    int16   `json:"quantity" copier:"must"`
		BrandName   string  `json:"brand_name"`

		CreatedAtDtStr string `json:"created_at"`
		UpdatedAtDtStr string `json:"updated_at"`
	}
)

func (p *Product) BrandName() string {
	return p.Brand.NamaBrand
}

func (p *Product) CreatedAtDtStr() string {
	return p.CreatedAt.Format(time.RFC3339)
}

func (p *Product) UpdatedAtDtStr() string {
	return p.UpdatedAt.Format(time.RFC3339)
}
