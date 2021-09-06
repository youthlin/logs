package kv

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

func TrimOdd(kvs []interface{}) []interface{} {
	count := len(kvs)
	if count < 2 {
		return nil
	}
	if count&1 != 0 { // 奇数, 去尾
		kvs = kvs[:count-1]
	}
	return kvs
}

func ToString(kvs []interface{}) [][]string {
	count := len(kvs)
	result := make([][]string, 0, count/2)
	for i := 0; i+1 < count; i += 2 {
		result = append(result, []string{valueToStr(kvs[i]), valueToStr(kvs[i+1])})
	}
	return result
}

func valueToStr(o interface{}) string {
	switch v := o.(type) {
	case nil:
		return ""
	case bool:
		if v {
			return "true"
		} else {
			return "false"
		}
	case string:
		return v
	case []byte:
		return string(v)
	case error:
		return v.Error()
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.Itoa(int(v))
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(int64(v), 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(uint64(v), 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', 3, 32)
	case float64:
		return strconv.FormatFloat(float64(v), 'f', 3, 64)
	case fmt.Stringer:
		value := reflect.ValueOf(o)
		if !value.IsValid() || value.Kind() == reflect.Ptr && value.IsNil() {
			return fmt.Sprintf("%#v", o)
		}
		return v.String()
	default:
		byteBuf := make([]byte, 0, 128)
		buf := (*jsonBuf)(unsafe.Pointer(&byteBuf))
		e := json.NewEncoder(buf)
		e.SetEscapeHTML(false)
		err := e.Encode(v)
		if err != nil {
			return fmt.Sprintf("%#v", o)
		}
		return string(byteBuf[:len(byteBuf)-1])
	}
}

type jsonBuf struct {
	buf []byte
}

func (b *jsonBuf) Write(p []byte) (n int, err error) {
	b.buf = append(b.buf, p...)
	return len(p), err
}
