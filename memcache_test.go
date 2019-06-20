package main

import (
	"testing"
)

func BenchmarkWriteMem(b *testing.B) {
	m := NewMem()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := m.set("key", "value")
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkReadMem(b *testing.B) {
	m := NewMem()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := m.get("key")
		if err != nil {
			b.Error(err)
		}
	}
}
