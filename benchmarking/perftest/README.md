# To bench tests

## First test

```bash
$ cd benchmarking/perftest
$ go test -bench=.
```

## With CPU profiling (with disabling GC)

```bash
$ GOGC=off go test -bench=BenchmarkRegex -cpuprofile cpu.out
$ go tool pprof perftest.test.exe cpu.out
~ web

```