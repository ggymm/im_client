package utils

import (
	"strconv"
)

func StrToUint(strNumber string, value interface{}) (err error) {
	var number interface{}
	number, err = strconv.ParseUint(strNumber, 10, 64)
	switch v := number.(type) {
	case uint64:
		switch d := value.(type) {
		case *uint64:
			*d = v
		case *uint:
			*d = uint(v)
		case *uint16:
			*d = uint16(v)
		case *uint32:
			*d = uint32(v)
		case *uint8:
			*d = uint8(v)
		}
	}
	return
}

func StrToInt(strNumber string, value interface{}) (err error) {
	var number interface{}
	number, err = strconv.ParseInt(strNumber, 10, 64)
	switch v := number.(type) {
	case int64:
		switch d := value.(type) {
		case *int64:
			*d = v
		case *int:
			*d = int(v)
		case *int16:
			*d = int16(v)
		case *int32:
			*d = int32(v)
		case *int8:
			*d = int8(v)
		}
	}
	return
}
