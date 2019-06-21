workerskv_benchmark
===================
Benchmark [Cloudflare Workers KV](https://developers.cloudflare.com/workers/kv/) with official sdk.  
You can compare own memcache server with WorkersKV. See [memcache.go](memcache.go).

```
BenchmarkWriteKV-16                            3         420761100 ns/op           58597 B/op        125 allocs/op
BenchmarkReadKV-16                             5         362196140 ns/op           12107 B/op        102 allocs/op
BenchmarkWriteKV_Parallel160-16                5        1077852620 ns/op           94801 B/op        722 allocs/op
BenchmarkReadKV_Parallel160-16                 5         247583660 ns/op           14038 B/op        143 allocs/op
```
