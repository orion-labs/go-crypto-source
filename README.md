# go-crypto-source
Options for using `crypto/rand` to back a `math/rand.Rand`

## Slides
There is not much actual code here, but part of the intention for this repository is to be an educational tool, so there is a slide presentation available in `./.slides/`.

To run the slides:
```
$ go get -u golang.org/x/tools/present
$ cd $REPO_ROOT/.slides/ && present .
```

## Bechmarks
The benchmarks that are output from the `benchmark_test.go` are illuminating:
```
BenchmarkGlobal-4         	50000000	        37.7 ns/op
BenchmarkNative-4         	100000000	        22.7 ns/op
BenchmarkSeed-4           	50000000	        23.9 ns/op
BenchmarkCryptoRead-4     	 2000000	       735 ns/op
BenchmarkCryptoUnsafe-4   	 2000000	       793 ns/op
BenchmarkCryptoSafe-4     	 2000000	       867 ns/op
```

On a developer workstation, it takes nearly a full millisecond to read 8 bytes from `crypto/rand`, which is about 20x slower than the deterministic `math/rand`. This means that if you are looking to build code that needs random data in a super tight loop, perhaps the techniques used in this repository are not appropriate for your use case.
