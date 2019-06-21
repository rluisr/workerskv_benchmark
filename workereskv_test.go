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

/*
func BenchmarkReadKV_Parallel160(b *testing.B) {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	kv := NewKV()
	b.ResetTimer()

	sem := make(chan struct{}, cpus)
	for i := 0; i < 160; i++ {
		sem <- struct{}{}
		go func() {
			_, err := kv.read("key")
			if err != nil {
				b.Error(err)
			}
			<-sem
		}()
	}
}
*/

func BenchmarkReadKV_Parallel160(b *testing.B) {
	kv := NewKV()

	b.SetParallelism(10)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := kv.read("key")
			if err != nil {
				b.Error(err)
			}
		}
	})
}
