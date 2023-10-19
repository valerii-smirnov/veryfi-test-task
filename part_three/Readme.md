# Strings similarity algorithms

I used a Golang library that includes implementations of various well-known string similarity calculation algorithms.
Library: https://github.com/adrg/strutil

This library calculates string similarity in percentages using different algorithms.

For benchmarks, I've used the Veryfi OCR result. 
You can find these files [here](similarity/testdata)

# Benchmarks

### Standard benchmarks run
```shell
go test -bench=.
```

```text
goos: darwin
goarch: arm64
pkg: github.com/valerii-smirnov/veryfi-test-task/part_three/similarity
BenchmarkSimilarityLevenshtein-10                     69          16818857 ns/op
BenchmarkSimilarityHamming-10                     189316              5924 ns/op
BenchmarkSimilarityJaccard-10                       4656            251090 ns/op
BenchmarkSimilarityJaro-10                           992           1208409 ns/op
BenchmarkSimilarityJaroWinkler-10                    973           1215689 ns/op
BenchmarkSimilarityOverlapCoefficient-10            4706            246430 ns/op
BenchmarkSimilaritySmithWatermanGotoh-10              44          26325739 ns/op
BenchmarkSimilaritySorensenDice-10                  3988            257573 ns/op
```

### Benchmarking of every algorithm twice by 5 seconds
```shell
go test -bench=. -count 2 -benchtime=5s
```

```text
goos: darwin
goarch: arm64
pkg: github.com/valerii-smirnov/veryfi-test-task/part_three/similarity
BenchmarkSimilarityLevenshtein-10                    384          15508491 ns/op
BenchmarkSimilarityLevenshtein-10                    390          15682371 ns/op
BenchmarkSimilarityHamming-10                     998826              5825 ns/op
BenchmarkSimilarityHamming-10                    1047457              5923 ns/op
BenchmarkSimilarityJaccard-10                      23736            235285 ns/op
BenchmarkSimilarityJaccard-10                      24588            241693 ns/op
BenchmarkSimilarityJaro-10                          5001           1204003 ns/op
BenchmarkSimilarityJaro-10                          4980           1213049 ns/op
BenchmarkSimilarityJaroWinkler-10                   4701           1222746 ns/op
BenchmarkSimilarityJaroWinkler-10                   4641           1251287 ns/op
BenchmarkSimilarityOverlapCoefficient-10           21387            257740 ns/op
BenchmarkSimilarityOverlapCoefficient-10           24760            256993 ns/op
BenchmarkSimilaritySmithWatermanGotoh-10             230          27543468 ns/op
BenchmarkSimilaritySmithWatermanGotoh-10             232          27324058 ns/op
BenchmarkSimilaritySorensenDice-10                 20217            342820 ns/op
BenchmarkSimilaritySorensenDice-10                 22576            256757 ns/op
PASS
```