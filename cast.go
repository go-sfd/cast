package cast

import (
	"fmt"
	"strconv"
)

const (
	TrueNum  = 1
	FalseNum = 0
	TrueStr  = "true"
	FalseStr = "false"
)

type BaseType interface {
	~bool | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~uint | ~int8 | ~int16 | ~int32 | ~int64 | ~int | ~float32 | ~float64 | ~string
}

func To[OUT BaseType](in any) (out OUT, b bool) {
	if v, ok := in.(bool); ok {
		switch any(out).(type) {
		case bool:
			out, b = any(v).(OUT)
		case uint8:
			if v {
				out, b = any(uint8(TrueNum)).(OUT)
			} else {
				out, b = any(uint8(FalseNum)).(OUT)
			}
		case uint16:
			if v {
				out, b = any(uint16(TrueNum)).(OUT)
			} else {
				out, b = any(uint16(FalseNum)).(OUT)
			}
		case uint32:
			if v {
				out, b = any(uint32(TrueNum)).(OUT)
			} else {
				out, b = any(uint32(FalseNum)).(OUT)
			}
		case uint64:
			if v {
				out, b = any(uint64(TrueNum)).(OUT)
			} else {
				out, b = any(uint64(FalseNum)).(OUT)
			}
		case uintptr:
			if v {
				out, b = any(uintptr(TrueNum)).(OUT)
			} else {
				out, b = any(uintptr(FalseNum)).(OUT)
			}
		case int8:
			if v {
				out, b = any(int8(TrueNum)).(OUT)
			} else {
				out, b = any(int8(FalseNum)).(OUT)
			}
		case int16:
			if v {
				out, b = any(int16(TrueNum)).(OUT)
			} else {
				out, b = any(int16(FalseNum)).(OUT)
			}
		case int32:
			if v {
				out, b = any(int32(TrueNum)).(OUT)
			} else {
				out, b = any(int32(FalseNum)).(OUT)
			}
		case int64:
			if v {
				out, b = any(int64(TrueNum)).(OUT)
			} else {
				out, b = any(int64(FalseNum)).(OUT)
			}
		case uint:
			if v {
				out, b = any(uint(TrueNum)).(OUT)
			} else {
				out, b = any(uint(FalseNum)).(OUT)
			}
		case int:
			if v {
				out, b = any(TrueNum).(OUT)
			} else {
				out, b = any(FalseNum).(OUT)
			}
		case float32:
			if v {
				out, b = any(float32(TrueNum)).(OUT)
			} else {
				out, b = any(float32(FalseNum)).(OUT)
			}
		case float64:
			if v {
				out, b = any(float64(TrueNum)).(OUT)
			} else {
				out, b = any(float64(FalseNum)).(OUT)
			}
		case string:
			if v {
				out, b = any(TrueStr).(OUT)
			} else {
				out, b = any(FalseStr).(OUT)
			}
		}

		return
	}

	if v, ok := in.(uint8); ok {
		switch any(out).(type) {
		case bool:
			out, b = any(v != 0).(OUT)
		case uint8:
			out, b = any(v).(OUT)
		case uint16:
			out, b = any(uint16(v)).(OUT)
		case uint32:
			out, b = any(uint32(v)).(OUT)
		case uint64:
			out, b = any(uint64(v)).(OUT)
		case uintptr:
			out, b = any(uintptr(v)).(OUT)
		case int8:
			out, b = any(int8(v)).(OUT)
		case int16:
			out, b = any(int16(v)).(OUT)
		case int32:
			out, b = any(int32(v)).(OUT)
		case int64:
			out, b = any(int64(v)).(OUT)
		case uint:
			out, b = any(uint(v)).(OUT)
		case int:
			out, b = any(int(v)).(OUT)
		case float32:
			out, b = any(float32(v)).(OUT)
		case float64:
			out, b = any(float64(v)).(OUT)
		case string:
			out, b = any(strconv.FormatUint(uint64(v), 10)).(OUT)
		}

		return
	}

	if v, ok := in.(uint16); ok {
		switch any(out).(type) {
		case bool:
			out, b = any(v != 0).(OUT)
		case uint8:
			out, b = any(uint8(v)).(OUT)
		case uint16:
			out, b = any(v).(OUT)
		case uint32:
			out, b = any(uint32(v)).(OUT)
		case uint64:
			out, b = any(uint64(v)).(OUT)
		case uintptr:
			out, b = any(uintptr(v)).(OUT)
		case int8:
			out, b = any(int8(v)).(OUT)
		case int16:
			out, b = any(int16(v)).(OUT)
		case int32:
			out, b = any(int32(v)).(OUT)
		case int64:
			out, b = any(int64(v)).(OUT)
		case uint:
			out, b = any(uint(v)).(OUT)
		case int:
			out, b = any(int(v)).(OUT)
		case float32:
			out, b = any(float32(v)).(OUT)
		case float64:
			out, b = any(float64(v)).(OUT)
		case string:
			out, b = any(strconv.FormatUint(uint64(v), 10)).(OUT)
		}

		return
	}

	if v, ok := in.(uint32); ok {
		switch any(out).(type) {
		case bool:
			out, b = any(v != 0).(OUT)
		case uint8:
			out, b = any(uint8(v)).(OUT)
		case uint16:
			out, b = any(uint16(v)).(OUT)
		case uint32:
			out, b = any(v).(OUT)
		case uint64:
			out, b = any(uint64(v)).(OUT)
		case uintptr:
			out, b = any(uintptr(v)).(OUT)
		case int8:
			out, b = any(int8(v)).(OUT)
		case int16:
			out, b = any(int16(v)).(OUT)
		case int32:
			out, b = any(int32(v)).(OUT)
		case int64:
			out, b = any(int64(v)).(OUT)
		case uint:
			out, b = any(uint(v)).(OUT)
		case int:
			out, b = any(int(v)).(OUT)
		case float32:
			out, b = any(float32(v)).(OUT)
		case float64:
			out, b = any(float64(v)).(OUT)
		case string:
			out, b = any(strconv.FormatUint(uint64(v), 10)).(OUT)
		}

		return
	}

	if v, ok := in.(uint64); ok {
		switch any(out).(type) {
		case bool:
			out, b = any(v != 0).(OUT)
		case uint8:
			out, b = any(uint8(v)).(OUT)
		case uint16:
			out, b = any(uint16(v)).(OUT)
		case uint32:
			out, b = any(uint32(v)).(OUT)
		case uint64:
			out, b = any(v).(OUT)
		case uintptr:
			out, b = any(uintptr(v)).(OUT)
		case int8:
			out, b = any(int8(v)).(OUT)
		case int16:
			out, b = any(int16(v)).(OUT)
		case int32:
			out, b = any(int32(v)).(OUT)
		case int64:
			out, b = any(int64(v)).(OUT)
		case uint:
			out, b = any(uint(v)).(OUT)
		case int:
			out, b = any(int(v)).(OUT)
		case float32:
			out, b = any(float32(v)).(OUT)
		case float64:
			out, b = any(float64(v)).(OUT)
		case string:
			out, b = any(strconv.FormatUint(v, 10)).(OUT)
		}

		return
	}

	if v, ok := in.(uintptr); ok {
		switch any(out).(type) {
		case bool:
			out, b = any(v != 0).(OUT)
		case uint8:
			out, b = any(uint8(v)).(OUT)
		case uint16:
			out, b = any(uint16(v)).(OUT)
		case uint32:
			out, b = any(uint32(v)).(OUT)
		case uint64:
			out, b = any(uint64(v)).(OUT)
		case uintptr:
			out, b = any(v).(OUT)
		case int8:
			out, b = any(int8(v)).(OUT)
		case int16:
			out, b = any(int16(v)).(OUT)
		case int32:
			out, b = any(int32(v)).(OUT)
		case int64:
			out, b = any(int64(v)).(OUT)
		case uint:
			out, b = any(uint(v)).(OUT)
		case int:
			out, b = any(int(v)).(OUT)
		case float32:
			out, b = any(float32(v)).(OUT)
		case float64:
			out, b = any(float64(v)).(OUT)
		case string:
			out, b = any(strconv.FormatUint(uint64(v), 10)).(OUT)
		}

		return
	}

	if v, ok := in.(int8); ok {
		switch any(out).(type) {
		case bool:
			out, b = any(v != 0).(OUT)
		case uint8:
			out, b = any(uint8(v)).(OUT)
		case uint16:
			out, b = any(uint16(v)).(OUT)
		case uint32:
			out, b = any(uint32(v)).(OUT)
		case uint64:
			out, b = any(uint64(v)).(OUT)
		case uintptr:
			out, b = any(uintptr(v)).(OUT)
		case int8:
			out, b = any(v).(OUT)
		case int16:
			out, b = any(int16(v)).(OUT)
		case int32:
			out, b = any(int32(v)).(OUT)
		case int64:
			out, b = any(int64(v)).(OUT)
		case uint:
			out, b = any(uint(v)).(OUT)
		case int:
			out, b = any(int(v)).(OUT)
		case float32:
			out, b = any(float32(v)).(OUT)
		case float64:
			out, b = any(float64(v)).(OUT)
		case string:
			out, b = any(strconv.FormatInt(int64(v), 10)).(OUT)
		}

		return
	}

	if v, ok := in.(int16); ok {
		switch any(out).(type) {
		case bool:
			out, b = any(v != 0).(OUT)
		case uint8:
			out, b = any(uint8(v)).(OUT)
		case uint16:
			out, b = any(uint16(v)).(OUT)
		case uint32:
			out, b = any(uint32(v)).(OUT)
		case uint64:
			out, b = any(uint64(v)).(OUT)
		case uintptr:
			out, b = any(uintptr(v)).(OUT)
		case int8:
			out, b = any(int8(v)).(OUT)
		case int16:
			out, b = any(v).(OUT)
		case int32:
			out, b = any(int32(v)).(OUT)
		case int64:
			out, b = any(int64(v)).(OUT)
		case uint:
			out, b = any(uint(v)).(OUT)
		case int:
			out, b = any(int(v)).(OUT)
		case float32:
			out, b = any(float32(v)).(OUT)
		case float64:
			out, b = any(float64(v)).(OUT)
		case string:
			out, b = any(strconv.FormatInt(int64(v), 10)).(OUT)
		}

		return
	}

	if v, ok := in.(int32); ok {
		switch any(out).(type) {
		case bool:
			out, b = any(v != 0).(OUT)
		case uint8:
			out, b = any(uint8(v)).(OUT)
		case uint16:
			out, b = any(uint16(v)).(OUT)
		case uint32:
			out, b = any(uint32(v)).(OUT)
		case uint64:
			out, b = any(uint64(v)).(OUT)
		case uintptr:
			out, b = any(uintptr(v)).(OUT)
		case int8:
			out, b = any(int8(v)).(OUT)
		case int16:
			out, b = any(int16(v)).(OUT)
		case int32:
			out, b = any(v).(OUT)
		case int64:
			out, b = any(int64(v)).(OUT)
		case uint:
			out, b = any(uint(v)).(OUT)
		case int:
			out, b = any(int(v)).(OUT)
		case float32:
			out, b = any(float32(v)).(OUT)
		case float64:
			out, b = any(float64(v)).(OUT)
		case string:
			out, b = any(strconv.FormatInt(int64(v), 10)).(OUT)
		}

		return
	}

	if v, ok := in.(int64); ok {
		switch any(out).(type) {
		case bool:
			out, b = any(v != 0).(OUT)
		case uint8:
			out, b = any(uint8(v)).(OUT)
		case uint16:
			out, b = any(uint16(v)).(OUT)
		case uint32:
			out, b = any(uint32(v)).(OUT)
		case uint64:
			out, b = any(uint64(v)).(OUT)
		case uintptr:
			out, b = any(uintptr(v)).(OUT)
		case int8:
			out, b = any(int8(v)).(OUT)
		case int16:
			out, b = any(int16(v)).(OUT)
		case int32:
			out, b = any(int32(v)).(OUT)
		case int64:
			out, b = any(v).(OUT)
		case uint:
			out, b = any(uint(v)).(OUT)
		case int:
			out, b = any(int(v)).(OUT)
		case float32:
			out, b = any(float32(v)).(OUT)
		case float64:
			out, b = any(float64(v)).(OUT)
		case string:
			out, b = any(strconv.FormatInt(v, 10)).(OUT)
		}

		return
	}

	if v, ok := in.(int); ok {
		switch any(out).(type) {
		case bool:
			out, b = any(v != 0).(OUT)
		case uint8:
			out, b = any(uint8(v)).(OUT)
		case uint16:
			out, b = any(uint16(v)).(OUT)
		case uint32:
			out, b = any(uint32(v)).(OUT)
		case uint64:
			out, b = any(uint64(v)).(OUT)
		case uintptr:
			out, b = any(uintptr(v)).(OUT)
		case int8:
			out, b = any(int8(v)).(OUT)
		case int16:
			out, b = any(int16(v)).(OUT)
		case int32:
			out, b = any(int32(v)).(OUT)
		case int64:
			out, b = any(int64(v)).(OUT)
		case uint:
			out, b = any(uint(v)).(OUT)
		case int:
			out, b = any(v).(OUT)
		case float32:
			out, b = any(float32(v)).(OUT)
		case float64:
			out, b = any(float64(v)).(OUT)
		case string:
			out, b = any(strconv.FormatInt(int64(v), 10)).(OUT)
		}

		return
	}

	if v, ok := in.(uint); ok {
		switch any(out).(type) {
		case bool:
			out, b = any(v != 0).(OUT)
		case uint8:
			out, b = any(uint8(v)).(OUT)
		case uint16:
			out, b = any(uint16(v)).(OUT)
		case uint32:
			out, b = any(uint32(v)).(OUT)
		case uint64:
			out, b = any(uint64(v)).(OUT)
		case uintptr:
			out, b = any(uintptr(v)).(OUT)
		case int8:
			out, b = any(int8(v)).(OUT)
		case int16:
			out, b = any(int16(v)).(OUT)
		case int32:
			out, b = any(int32(v)).(OUT)
		case int64:
			out, b = any(int64(v)).(OUT)
		case uint:
			out, b = any(v).(OUT)
		case int:
			out, b = any(int(v)).(OUT)
		case float32:
			out, b = any(float32(v)).(OUT)
		case float64:
			out, b = any(float64(v)).(OUT)
		case string:
			out, b = any(strconv.FormatUint(uint64(v), 10)).(OUT)
		}

		return
	}

	if v, ok := in.(float32); ok {
		switch any(out).(type) {
		case bool:
			out, b = any(v != 0).(OUT)
		case uint8:
			out, b = any(uint8(v)).(OUT)
		case uint16:
			out, b = any(uint16(v)).(OUT)
		case uint32:
			out, b = any(uint32(v)).(OUT)
		case uint64:
			out, b = any(uint64(v)).(OUT)
		case uintptr:
			out, b = any(uintptr(v)).(OUT)
		case int8:
			out, b = any(int8(v)).(OUT)
		case int16:
			out, b = any(int16(v)).(OUT)
		case int32:
			out, b = any(int32(v)).(OUT)
		case int64:
			out, b = any(int64(v)).(OUT)
		case uint:
			out, b = any(uint(v)).(OUT)
		case int:
			out, b = any(int(v)).(OUT)
		case float32:
			out, b = any(v).(OUT)
		case float64:
			out, b = any(float64(v)).(OUT)
		case string:
			out, b = any(strconv.FormatFloat(float64(v), 'g', -1, 32)).(OUT)
		}

		return
	}

	if v, ok := in.(float64); ok {
		switch any(out).(type) {
		case bool:
			out, b = any(v != 0).(OUT)
		case uint8:
			out, b = any(uint8(v)).(OUT)
		case uint16:
			out, b = any(uint16(v)).(OUT)
		case uint32:
			out, b = any(uint32(v)).(OUT)
		case uint64:
			out, b = any(uint64(v)).(OUT)
		case uintptr:
			out, b = any(uintptr(v)).(OUT)
		case int8:
			out, b = any(int8(v)).(OUT)
		case int16:
			out, b = any(int16(v)).(OUT)
		case int32:
			out, b = any(int32(v)).(OUT)
		case int64:
			out, b = any(int64(v)).(OUT)
		case uint:
			out, b = any(uint(v)).(OUT)
		case int:
			out, b = any(int(v)).(OUT)
		case float32:
			out, b = any(float32(v)).(OUT)
		case float64:
			out, b = any(v).(OUT)
		case string:
			out, b = any(strconv.FormatFloat(v, 'g', -1, 64)).(OUT)
		}

		return
	}

	if vv, ok := in.(fmt.Stringer); ok {
		v := vv.String()
		switch any(out).(type) {
		case bool:
			var e error
			b, e = strconv.ParseBool(v)
			if e == nil {
				out, b = any(b).(OUT)
			}
		case uint8:
			out, b = any(uint8(tryParseUnsignedNumber(v))).(OUT)
		case uint16:
			out, b = any(uint16(tryParseUnsignedNumber(v))).(OUT)
		case uint32:
			out, b = any(uint32(tryParseUnsignedNumber(v))).(OUT)
		case uint64:
			out, b = any(tryParseUnsignedNumber(v)).(OUT)
		case uintptr:
			out, b = any(uintptr(tryParseUnsignedNumber(v))).(OUT)
		case int8:
			out, b = any(int8(tryParseNumber(v))).(OUT)
		case int16:
			out, b = any(int16(tryParseNumber(v))).(OUT)
		case int32:
			out, b = any(int32(tryParseNumber(v))).(OUT)
		case int64:
			out, b = any(tryParseNumber(v)).(OUT)
		case uint:
			out, b = any(uint(tryParseUnsignedNumber(v))).(OUT)
		case int:
			out, b = any(int(tryParseNumber(v))).(OUT)
		case float32:
			vv, _ := strconv.ParseFloat(v, 64)
			out, b = any(float32(vv)).(OUT)
		case float64:
			vv, _ := strconv.ParseFloat(v, 64)
			out, b = any(vv).(OUT)
		case string:
			out, b = any(v).(OUT)
		}

		return
	}

	if v, ok := in.(string); ok {
		switch any(out).(type) {
		case bool:
			var e error
			b, e = strconv.ParseBool(v)
			if e == nil {
				out, b = any(b).(OUT)
			}
		case uint8:
			out, b = any(uint8(tryParseUnsignedNumber(v))).(OUT)
		case uint16:
			out, b = any(uint16(tryParseUnsignedNumber(v))).(OUT)
		case uint32:
			out, b = any(uint32(tryParseUnsignedNumber(v))).(OUT)
		case uint64:
			out, b = any(tryParseUnsignedNumber(v)).(OUT)
		case uintptr:
			out, b = any(uintptr(tryParseUnsignedNumber(v))).(OUT)
		case int8:
			out, b = any(int8(tryParseNumber(v))).(OUT)
		case int16:
			out, b = any(int16(tryParseNumber(v))).(OUT)
		case int32:
			out, b = any(int32(tryParseNumber(v))).(OUT)
		case int64:
			out, b = any(tryParseNumber(v)).(OUT)
		case uint:
			out, b = any(uint(tryParseUnsignedNumber(v))).(OUT)
		case int:
			out, b = any(int(tryParseNumber(v))).(OUT)
		case float32:
			vv, _ := strconv.ParseFloat(v, 64)
			out, b = any(float32(vv)).(OUT)
		case float64:
			vv, _ := strconv.ParseFloat(v, 64)
			out, b = any(vv).(OUT)
		case string:
			out, b = any(v).(OUT)
		}

		return
	}

	unexpectedInput := fmt.Sprintf("%v", in)
	switch any(out).(type) {
	case bool:
		var e error
		b, e = strconv.ParseBool(unexpectedInput)
		if e == nil {
			out, b = any(b).(OUT)
		}
	case uint8:
		out, b = any(uint8(tryParseUnsignedNumber(unexpectedInput))).(OUT)
	case uint16:
		out, b = any(uint16(tryParseUnsignedNumber(unexpectedInput))).(OUT)
	case uint32:
		out, b = any(uint32(tryParseUnsignedNumber(unexpectedInput))).(OUT)
	case uint64:
		out, b = any(tryParseUnsignedNumber(unexpectedInput)).(OUT)
	case uintptr:
		out, b = any(uintptr(tryParseUnsignedNumber(unexpectedInput))).(OUT)
	case int8:
		out, b = any(int8(tryParseNumber(unexpectedInput))).(OUT)
	case int16:
		out, b = any(int16(tryParseNumber(unexpectedInput))).(OUT)
	case int32:
		out, b = any(int32(tryParseNumber(unexpectedInput))).(OUT)
	case int64:
		out, b = any(tryParseNumber(unexpectedInput)).(OUT)
	case uint:
		out, b = any(uint(tryParseUnsignedNumber(unexpectedInput))).(OUT)
	case int:
		out, b = any(int(tryParseNumber(unexpectedInput))).(OUT)
	case float32:
		vv, _ := strconv.ParseFloat(unexpectedInput, 64)
		out, b = any(float32(vv)).(OUT)
	case float64:
		vv, _ := strconv.ParseFloat(unexpectedInput, 64)
		out, b = any(vv).(OUT)
	case string:
		out, b = any(unexpectedInput).(OUT)
	}

	return
}

func tryParseNumber(v string) int64 {
	vv, e := strconv.ParseInt(v, 10, 64)
	if e == nil {
		return vv
	}

	vvv, _ := strconv.ParseFloat(v, 64)
	return int64(vvv)
}

func tryParseUnsignedNumber(v string) uint64 {
	vv, e := strconv.ParseUint(v, 10, 64)
	if e == nil {
		return vv
	}

	vvv, _ := strconv.ParseFloat(v, 64)
	return uint64(vvv)
}
