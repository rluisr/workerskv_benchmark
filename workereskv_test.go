package main

import (
	"testing"
)

func BenchmarkWriteKV(b *testing.B) {
	kv := NewKV()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := kv.write("key", "value")
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkReadKV(b *testing.B) {
	kv := NewKV()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := kv.read("key")
		if err != nil {
			b.Error(err)
		}
	}
}
