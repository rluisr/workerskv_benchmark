workerskv_benchmark
===================
Benchmark [Cloudflare Workers KV](https://developers.cloudflare.com/workers/kv/) with official sdk. 
memcache server is [memcachier](https://www.memcachier.com/)'s free plan.

```
$ go test -bench . -benchmem
goos: windows
goarch: amd64
pkg: workerskv-benchmark
BenchmarkWriteMem-16                 200           8696971 ns/op              73 B/op          3 allocs/op
BenchmarkReadMem-16                  200           9917228 ns/op             152 B/op          4 allocs/op
BenchmarkWriteKV-16                    1        1191774700 ns/op          444056 B/op       7953 allocs/op
BenchmarkReadKV-16                     2        1196770400 ns/op           15292 B/op        108 allocs/op
PASS
ok      workerskv-benchmark     9.358s
```

About 1sec per operation.
