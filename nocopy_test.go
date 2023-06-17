package nocopy_test

import (
	"bytes"
	"fmt"
	"runtime"
	"testing"

	"github.com/colduction/nocopy"
)

const sampleText = "Hello World!"

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

func BenchmarkStringToByteSlice(b *testing.B) {
	benchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				nocopy.StringToByteSlice(sampleText)
			}
		})
	})
}

func BenchmarkNewBufferString(b *testing.B) {
	benchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				bytes.NewBufferString(sampleText)
			}
		})
	})
}

func BenchmarkByteSliceToString(b *testing.B) {
	benchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				nocopy.ByteSliceToString(sampleSlice)
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
