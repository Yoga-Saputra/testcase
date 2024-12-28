package helper

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/Yoga-Saputra/testcase/config"
	"github.com/Yoga-Saputra/testcase/internal/entity/std"
	"github.com/labstack/echo/v4"
)

func InArray(v interface{}, in interface{}) (ok bool) {
	val := reflect.Indirect(reflect.ValueOf(in))
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if ok = v == val.Index(i).Interface(); ok {
				return
			}
		}
	}
	return
}

func LoadLocation() time.Time {
	loc, err := time.LoadLocation(config.Of.App.TimeZone)
	if err != nil {
		fmt.Println("PANIC_LOGIC")
		panic(err)
	}

	return time.Now().In(loc)
}

func StructValidator(c echo.Context, p interface{}) *std.APIResponse {
	// Parsing params/payload to struct
	if err := c.Bind(p); err != nil {
		return std.APIResponseError(std.StatusBadRequest, errors.New("failed to parsing request params/payloads"))
	}

	// Validate the params/payloads
	if err := c.Validate(p); err != nil {
		return std.APIResponseError(std.StatusBadRequest, err)
	}

	return nil
}
