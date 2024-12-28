package convert

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// AlwaysFloat64 to convert all possible given argument to float64
func AlwaysFloat64(p interface{}) (float64, error) {
	switch p.(type) {
	case int:
		return float64(p.(int)), nil
	case int16:
		return float64(p.(int16)), nil
	case int32:
		return float64(p.(int32)), nil
	case int64:
		return float64(p.(int64)), nil
	case string:
		val, err := strconv.ParseFloat(fmt.Sprintf("%v", p), 64)
		if err != nil {
			return 0, err
		}

		return val, nil
	case float64:
		return float64(p.(float64)), nil
	case float32:
		return float64(p.(float32)), nil
	default:
		return 0, errors.New("unknown type of arguments")
	}
}

// Float64ToUint64 to convert float64 value to uint64 value
func Float64ToUint64(f float64) uint64 {
	var val uint64 = uint64(f)

	return val
}

// RoundDownDecimal2Places rounding down given float to float with 2 decimal places
func RoundDownDecimal2Places(f float64) float64 {
	return math.Floor(f*100) / 100
}

// InterfaceToMap convert given interface to map[string]interface{}
func InterfaceToMap(i interface{}) (map[string]interface{}, error) {
	// Marshal interface to []byte
	b, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}

	// Unmarshal []byte to map[string]interface{}
	var result map[string]interface{}
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CaseInsensitiveReplace replace in-sensitive string
func CaseInsensitiveReplace(subject, search, replace string) string {
	searchRegex := regexp.MustCompile("(?i)" + search)
	return searchRegex.ReplaceAllString(subject, replace)
}

// Reverse string
func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}

	return
}

func GetBranchIdFromPlayId(s string) (branchId uint64) {
	// Prepare BranchID
	mc := s
	idx := strings.Index(Reverse(mc), "u")
	if idx < 0 {
		branchId = 0
		err := fmt.Errorf("cannot find `u` on playId: %s", s)
		log.Printf("[GetBranchIdFromPlayId] Failed to parse branchId: %s", err.Error())

		return
	}

	mc = mc[len(mc)-idx:]
	bid, err := strconv.Atoi(mc)
	if err != nil {
		log.Printf("[GetBranchIdFromPlayId] Failed to parse branchId: %s", err.Error())
	}

	branchId = uint64(bid)
	return
}
