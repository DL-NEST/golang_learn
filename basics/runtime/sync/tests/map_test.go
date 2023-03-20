package tests_test

import (
	"sync"
	"testing"
)

func BenchmarkMap(b *testing.B) {
	m1 := make(map[int]any, 1000000)
	m2 := sync.Map{}

	for i := 0; i < 1000000; i++ {
		m1[i] = i
	}

	b.Run("map", func(b *testing.B) {
		var a any
		var ok bool
		for i := 0; i < b.N; i++ {
			a, ok = m1[956799]
			if ok {

			}
		}
		b.Log(a)
	})
	b.Run("sync.map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = m2.Load("ss")
		}
	})
}
