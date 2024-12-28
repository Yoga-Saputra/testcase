package ucv1producthttp

type (
	product__Reqp struct {
		BrandID     uint16  `json:"brand_id" form:"brand_id" xml:"brand_id" validate:"required"`
		NamaProduct string  `json:"nama_product" form:"nama_product" xml:"nama_product" validate:"required"`
		Harga       float64 `json:"harga" form:"harga" xml:"harga" validate:"required"`
		Qty         int16   `json:"qty" form:"qty" xml:"qty" validate:"required"`
	}

	list__Reqp struct {
		/// The dataTable search request.
		Draw   int `json:"draw" form:"draw" xml:"draw"`       /// The dataTable draw flag.
		Length int `json:"length" form:"length" xml:"length"` /// The dataTable offset request.
		Offset int `json:"offset" form:"offset" xml:"offset"` // The dataTable length or limit request.
	}
)
