package nocopy_test

import (
	"bytes"
	"fmt"
	"runtime"
	"testing"

	"github.com/colduction/nocopy"
)

const (
	sampleText string = "Hello World!"
	sampleByte byte   = 80
)

var sampleSlice []byte = []byte{}

func benchPerCoreConfigs(b *testing.B, f func(b *testing.B)) {
	b.Helper()
	coreConfigs := []int{1, 2, 4, 8}
	for _, n := range coreConfigs {
		name := fmt.Sprintf("%d cores", n)
		b.Run(name, func(b *testing.B) {
			runtime.GOMAXPROCS(n)
			f(b)
		})
	}
}

// StringToByteSlice
func BenchmarkStringToByteSlice(b *testing.B) {
	benchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				_ = nocopy.StringToByteSlice(sampleText)
			}
		})
	})
}

func BenchmarkNewBufferStringBytes(b *testing.B) {
	benchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				_ = bytes.NewBufferString(sampleText).Bytes()
			}
		})
	})
}

// ByteSliceToString
func BenchmarkByteSliceToString(b *testing.B) {
	benchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				_ = nocopy.ByteSliceToString(sampleSlice)
			}
		})
	})
}

func BenchmarkNewBuffer(b *testing.B) {
	benchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				_ = bytes.NewBuffer(sampleSlice).String()
			}
		})
	})
}

// ByteToByteSlice
func BenchmarkByteToByteSlice(b *testing.B) {
	benchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				_ = nocopy.ByteToByteSlice(sampleByte)
			}
		})
	})
}

func BenchmarkSafeByteToByteSlice(b *testing.B) {
	benchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				_ = []byte{sampleByte}
			}
		})
	})
}

// StringToStringSlice
func BenchmarkStringToStringSlice(b *testing.B) {
	benchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				_ = nocopy.StringToStringSlice(sampleText)
			}
		})
	})
}

func BenchmarkSafeStringToStringSlice(b *testing.B) {
	benchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				_ = []string{sampleText}
			}
		})
	})
}
