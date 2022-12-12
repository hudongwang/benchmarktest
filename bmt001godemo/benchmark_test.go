package testing_sp

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkAdd(b *testing.B) {
	for i, n := 0, 0; i < b.N; i++ {
		n++
	}
}

func BenchmarkAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("i: %d", i)
	}
}

func BenchmarkSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(time.Millisecond * 10)
	}
}

func BenchmarkSleep2(b *testing.B) {
	time.Sleep(time.Millisecond * 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		time.Sleep(time.Millisecond * 10)
		b.StopTimer()
		time.Sleep(time.Millisecond * 10)
		b.StartTimer()
	}
}
