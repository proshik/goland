package main

import (
	"strconv"
	"sync"
	"unicode/utf8"
)

var (
	bytesPool = sync.Pool{
		New: func() interface{} {
			return make([]byte, 4094)
		},
	}
)

func ConvertInt64ToInt32(int64 int64) int32 {
	return int32(int64)
}

func ConvertStringToByteSlice(string string) []byte {
	return []byte(string)
}

func ConvertSliceByteToString(sliceByte []byte) string {
	s := string(sliceByte)
	return s
}

func Utf8Unescaped(input []byte) []byte {
	buf := bytesPool.Get().([]byte)[:0]

	var tmp [4]byte

	i, l := 0, len(input)
	for i < l {
		ch := input[i]

		// \u1234

		// в случае любых ошибок просто пропускаем один байт
		if ch != '\\' {
		} else if (i >= l-4) || input[i+1] != 'u' {
		} else if r, err := strconv.ParseUint(string(input[i+2:i+6]), 16, 64); err != nil {
		} else {
			n := utf8.EncodeRune(tmp[:], rune(r))
			buf = append(buf, tmp[:n]...)
			i += 6
			continue
		}

		buf = append(buf, ch)
		i++
	}

	res := append([]byte{}, buf...)

	bytesPool.Put(buf)

	return res
}
