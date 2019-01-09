package main

import (
	"testing"
)

var (
	sourceString           = "Test string for a benchmark"
	sourceSliceByte        = []byte("Test string for a benchmark")
	sourceForUtf8Unescaped = []byte(`\u0441\u0432\u043e\u0431\u043e\u0434\u043d\u044b`)
)

var result interface{}

func BenchmarkConvertInt64ToInt32(b *testing.B) {
	var sourceInt64 = int64(5675353422)
	var nI int32

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		nI = ConvertInt64ToInt32(sourceInt64)
	}
	result = nI
}

func BenchmarkConvertInt32ToInt64(b *testing.B) {
	var sourceInt32 = int32(56753422)
	var nI int64

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		nI = int64(sourceInt32)
	}
	result = nI
}

func BenchmarkConvertStringToByteSlice(b *testing.B) {
	var slice []byte
	for i := 0; i < b.N; i++ {
		slice = []byte(sourceString)
	}
	result = slice
}

func BenchmarkConvertSliceByteToString(b *testing.B) {
	var r string
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r = string(sourceSliceByte)
	}
	result = r
}

func BenchmarkUtf8Unescaped(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Utf8Unescaped(sourceForUtf8Unescaped)
	}
}
