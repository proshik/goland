# Commands

## for pprof

```bash
$ go test -bench=BenchmarkUtf8Unescaped -benchmem -memprofile memprofile.out -cpuprofile profile.out
```

```bash
$ go tool pprof profile.out
```

or 

```bash
$ go tool pprof memprofile.out
```