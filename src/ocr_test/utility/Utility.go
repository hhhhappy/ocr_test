package utility

import (
	"github.com/satori/go.uuid"
	"strings"
	"strconv"
)

func InArrayString(needle string, haystack []string) bool {
	for _, unit := range haystack  {
		if unit == needle{
			return true
		}
	}
	return false
}

func InArrayInt(needle int, haystack []int) bool {
	for _, unit := range haystack  {
		if unit == needle{
			return true
		}
	}
	return false
}

func GetUUID() string {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	id = strings.Replace(id, "-", "", 4)
	return id
}

func InterfaceToInt(i interface{}) (int, bool) {
	if v, ok := i.(float32); ok {
		return int(v), true
	}
	if v, ok := i.(float64); ok {
		return int(v), true
	}
	if v, ok := i.(int64); ok {
		return int(v), true
	}
	if v, ok := i.(uint64); ok {
		return int(v), true
	}
	if v, ok := i.(int32); ok {
		return int(v), true
	}
	if v, ok := i.(uint32); ok {
		return int(v), true
	}
	if v, ok := i.(int); ok {
		return int(v), true
	}
	if v, ok := i.(uint); ok {
		return int(v), true
	}
	if v, ok := i.(int8); ok {
		return int(v), true
	}
	if v, ok := i.(uint8); ok {
		return int(v), true
	}

	type s interface {
		String() string
	}
	if v, ok := i.(s); ok {
		v1, _ := strconv.ParseInt(v.String(), 10, strconv.IntSize)
		return int(v1), true
	}
	if v, ok := i.(string); ok {
		v1, _ := strconv.ParseInt(v, 10, strconv.IntSize)
		return int(v1), true
	}
	return 0, false

}
