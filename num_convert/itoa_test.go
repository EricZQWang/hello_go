package main
import (
	"fmt"
	"strconv"
	"testing"
)

var smallInt = 35
var bigInt = 999999999999999

func BenchmarkItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := strconv.Itoa(smallInt)
		_ = val
	}
}

func BenchmarkItoaFormatInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := strconv.FormatInt(int64(smallInt), 10)
		_ = val
	}
}

func BenchmarkItoaSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := fmt.Sprint(smallInt)
		_ = val
	}
}

func BenchmarkItoaSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := fmt.Sprintf("%d", smallInt)
		_ = val
	}
}

func BenchmarkItoaBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := strconv.Itoa(bigInt)
		_ = val
	}
}

func BenchmarkItoaFormatIntBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := strconv.FormatInt(int64(bigInt), 10)
		_ = val
	}
}

func BenchmarkItoaSprintBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := fmt.Sprint(bigInt)
		_ = val
	}
}

func BenchmarkItoaSprintfBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := fmt.Sprintf("%d", bigInt)
		_ = val
	}
}

var smallFloat =1.01

var bigFloat =1.01E10
func BenchmarkFloatItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := strconv.Itoa(int(smallFloat))
		_ = val
	}
}

func BenchmarkFloatItoaFormatInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := strconv.FormatInt(int64(smallFloat), 10)
		_ = val
	}
}

func BenchmarkStrconvFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := strconv.FormatFloat(smallFloat,'b',0,64)
		_ = val
	}
}

func BenchmarkFloatItoaSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := fmt.Sprint(smallFloat)
		_ = val
	}
}

func BenchmarkFloatItoaSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := fmt.Sprintf("%0.0f", smallFloat)
		_ = val
	}
}
func BenchmarkBigFloatItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := strconv.Itoa(int(bigFloat))
		_ = val
	}
}

func BenchmarkBigFloatItoaFormatInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := strconv.FormatInt(int64(bigFloat), 10)
		_ = val
	}
}

func BenchmarkStrconvBigFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := strconv.FormatFloat(bigFloat,'b',0,64)
		_ = val
	}
}

func BenchmarkBigFloatItoaSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := fmt.Sprint(bigFloat)
		_ = val
	}
}

func BenchmarkBigFloatItoaSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := fmt.Sprintf("%0.0f", bigFloat)
		_ = val
	}
}


