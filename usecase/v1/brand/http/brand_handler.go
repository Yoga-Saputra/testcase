package ucv1brandhttp

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/Yoga-Saputra/testcase/internal/entity/std"
	"github.com/Yoga-Saputra/testcase/internal/helper"
	"github.com/labstack/echo/v4"
)

// create brand
func (ds *domainService) Create(c echo.Context) error {
	var apiResp *std.APIResponse
	p := new(Brand__Reqp)

	// Validate
	if apiResp = helper.StructValidator(c, p); apiResp != nil {
		return c.JSON(int(apiResp.StatusCode), apiResp.Body)
	}

	// Log payload
	log.Printf("[CreateBrand-%s][Payload] - %v", p.BrandName, *p)

	if err := ds.s.CreateBrand(p.BrandName); err != nil {
		log.Printf("[CreateBrand] - (%s) Error: %s", p.BrandName, err.Error())

		if apiResp = std.APIResponseError(std.StatusServerError, err); apiResp != nil {
			return c.JSON(int(apiResp.StatusCode), apiResp.Body)
		} else {
			return c.NoContent(500)
		}
	}

	// Return success
	apiResp = std.APIResponseSuccess(fmt.Sprintf("Brand name [%s] successfully created", p.BrandName))
	return c.JSON(int(apiResp.StatusCode), apiResp.Body)
}

// delete brand
func (ds *domainService) Delete(c echo.Context) error {
	var apiResp *std.APIResponse

	brandId, err := strconv.ParseUint(c.Param("id"), 10, 16) // base 10, bit size 16
	if err != nil {
		log.Fatalf("Error parsing string to uint16: %v", err)
	}

	// Since ParseUint returns a uint64, we need to explicitly cast it to uint16
	BrandIdUint16 := uint16(brandId)

	// Log payload
	log.Printf("[DeleteBrand-%v][Payload]", BrandIdUint16)

	ok := ds.s.FindByBrandId(BrandIdUint16)

	if ok {
		if apiResp = std.APIResponseError(std.StatusForbidden, errors.New("brandID has been used")); apiResp != nil {
			return c.JSON(int(apiResp.StatusCode), apiResp.Body)
		} else {
			return c.NoContent(500)
		}
	}

	rows, err := ds.s.DeleteByBrandID(BrandIdUint16)
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
	apiResp = std.APIResponseSuccess(fmt.Sprintf("Brand ID [%v] successfully deleted", BrandIdUint16))
	return c.JSON(int(apiResp.StatusCode), apiResp.Body)
}
