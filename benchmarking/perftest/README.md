# To bench tests

## First test

```bash
$ cd benchmarking/perftest
$ go test -benchmem -bench=. 
```

## With CPU profiling (with disabling GC)

```bash
$ GOGC=off go test -bench=BenchmarkRegex -cpuprofile cpu.out
# change the name of file
$ go tool pprof perftest.test.exe cpu.out
~ web

```