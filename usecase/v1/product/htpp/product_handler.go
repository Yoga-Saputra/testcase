package ucv1producthttp

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/Yoga-Saputra/testcase/internal/entity/std"
	"github.com/Yoga-Saputra/testcase/internal/helper"
	"github.com/Yoga-Saputra/testcase/internal/repo"
	"github.com/Yoga-Saputra/testcase/usecase/v1/product"
	"github.com/labstack/echo/v4"
)

// create brand
func (ds *domainService) List(c echo.Context) error {
	var apiResp *std.APIResponse
	p := new(list__Reqp)

	// Validate
	if apiResp = helper.StructValidator(c, p); apiResp != nil {
		return c.JSON(int(apiResp.StatusCode), apiResp.Body)
	}

	// Prepare response
	resp := &repo.RepoDatatableResponse{Draw: p.Draw + 1}
	defer func() {
		resp = nil
	}()

	pp := &product.PProduct{
		Limit:  int64(p.Length),
		Offset: int64(p.Offset),
	}

	countTotoal, list, err := ds.s.GetProductListPagination(pp)
	if err != nil {
		log.Printf("[ListProduct] - Error: %s", err.Error())

		if apiResp = std.APIResponseError(std.StatusServerError, err); apiResp != nil {
			return c.JSON(int(apiResp.StatusCode), apiResp.Body)
		} else {
			return c.NoContent(500)
		}
	}

	resp.RecordsTotal = countTotoal
	resp.Data = list

	// Return success
	return c.JSON(int(std.StatusOK), resp)
}

// create brand
func (ds *domainService) Create(c echo.Context) error {
	var apiResp *std.APIResponse
	p := new(product__Reqp)

	// Validate
	if apiResp = helper.StructValidator(c, p); apiResp != nil {
		return c.JSON(int(apiResp.StatusCode), apiResp.Body)
	}

	// Log payload
	log.Printf("[CreateProduct-%s][Payload] - %v", p.NamaProduct, *p)

	if err := ds.s.CreateProduct(p.BrandID, p.NamaProduct, p.Harga, p.Qty); err != nil {
		log.Printf("[CreateProduct] - (%s) Error: %s", p.NamaProduct, err.Error())

		if apiResp = std.APIResponseError(std.StatusServerError, err); apiResp != nil {
			return c.JSON(int(apiResp.StatusCode), apiResp.Body)
		} else {
			return c.NoContent(500)
		}
	}

	// Return success
	apiResp = std.APIResponseSuccess(fmt.Sprintf("Product [%s] successfully created", p.NamaProduct))
	return c.JSON(int(apiResp.StatusCode), apiResp.Body)
}

// Update brand
func (ds *domainService) Update(c echo.Context) error {
	var apiResp *std.APIResponse
	p := new(product__Reqp)

	// Validate
	if apiResp = helper.StructValidator(c, p); apiResp != nil {
		return c.JSON(int(apiResp.StatusCode), apiResp.Body)
	}

	pID, err := strconv.ParseUint(c.Param("id"), 10, 64) // base 10, bit size 64
	if err != nil {
		log.Fatalf("Error parsing string to uint16: %v", err)
	}

	// Since ParseUint returns a uint64, we need to explicitly cast it to uint64
	pIDUint64 := uint64(pID)

	// Log payload
	log.Printf("[UpdateProduct[%s]-pID[%d]][Payload] - %v", p.NamaProduct, pIDUint64, *p)

	// Update record
	_, err = ds.s.UpdateProduct(pIDUint64, p.BrandID, p.NamaProduct, p.Harga, p.Qty)
	if err != nil {
		log.Printf("[UpdateProduct] - (%s) Error: %s", p.NamaProduct, err.Error())

		if apiResp = std.APIResponseError(std.StatusServerError, err); apiResp != nil {
			return c.JSON(int(apiResp.StatusCode), apiResp.Body)
		} else {
			return c.NoContent(500)
		}
	}

	// Return success
	apiResp = std.APIResponseSuccess("Data product successfully updated")
	return c.JSON(int(apiResp.StatusCode), apiResp.Body)
}

// delete product
func (ds *domainService) Delete(c echo.Context) error {
	var apiResp *std.APIResponse

	pID, err := strconv.ParseUint(c.Param("id"), 10, 64) // base 10, bit size 64
	if err != nil {
		log.Fatalf("Error parsing string to uint16: %v", err)
	}

	// Since ParseUint returns a uint64, we need to explicitly cast it to uint64
	pIDUint64 := uint64(pID)

	rows, err := ds.s.DeleteByProductID(pIDUint64)

	switch {
	// Always check error appear first
	case err != nil:
		if apiResp = std.APIResponseError(std.StatusServerError, err); apiResp != nil {
			return c.JSON(int(apiResp.StatusCode), apiResp.Body)
		} else {
			return c.NoContent(500)
		}

	case rows == 0:
		if apiResp = std.APIResponseError(std.StatusNotFound, errors.New("record not found")); apiResp != nil {
			return c.JSON(int(apiResp.StatusCode), apiResp.Body)
		} else {
			return c.NoContent(500)
		}
	}

	// Return success
	apiResp = std.APIResponseSuccess(fmt.Sprintf("Brand ID [%v] successfully deleted", pIDUint64))
	return c.JSON(int(apiResp.StatusCode), apiResp.Body)
}
