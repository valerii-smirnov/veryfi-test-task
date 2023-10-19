package similarity

import (
	"github.com/adrg/strutil/metrics"
	"os"
	"testing"
)

func BenchmarkSimilarityLevenshtein(b *testing.B) {
	first, second, err := getOCRStrings("testdata/1.txt", "testdata/2.txt")
	if err != nil {

	}

	for i := 0; i < b.N; i++ {
		Similarity(first, second, metrics.NewLevenshtein())
	}
}

func BenchmarkSimilarityHamming(b *testing.B) {
	first, second, err := getOCRStrings("testdata/1.txt", "testdata/2.txt")
	if err != nil {

	}

	for i := 0; i < b.N; i++ {
		Similarity(first, second, metrics.NewHamming())
	}
}

func BenchmarkSimilarityJaccard(b *testing.B) {
	first, second, err := getOCRStrings("testdata/1.txt", "testdata/2.txt")
	if err != nil {

	}

	for i := 0; i < b.N; i++ {
		Similarity(first, second, metrics.NewJaccard())
	}
}

func BenchmarkSimilarityJaro(b *testing.B) {
	first, second, err := getOCRStrings("testdata/1.txt", "testdata/2.txt")
	if err != nil {

	}

	for i := 0; i < b.N; i++ {
		Similarity(first, second, metrics.NewJaro())
	}
}

func BenchmarkSimilarityJaroWinkler(b *testing.B) {
	first, second, err := getOCRStrings("testdata/1.txt", "testdata/2.txt")
	if err != nil {

	}

	for i := 0; i < b.N; i++ {
		Similarity(first, second, metrics.NewJaroWinkler())
	}
}

func BenchmarkSimilarityOverlapCoefficient(b *testing.B) {
	first, second, err := getOCRStrings("testdata/1.txt", "testdata/2.txt")
	if err != nil {

	}

	for i := 0; i < b.N; i++ {
		Similarity(first, second, metrics.NewOverlapCoefficient())
	}
}

func BenchmarkSimilaritySmithWatermanGotoh(b *testing.B) {
	first, second, err := getOCRStrings("testdata/1.txt", "testdata/2.txt")
	if err != nil {

	}

	for i := 0; i < b.N; i++ {
		Similarity(first, second, metrics.NewSmithWatermanGotoh())
	}
}

func BenchmarkSimilaritySorensenDice(b *testing.B) {
	first, second, err := getOCRStrings("testdata/1.txt", "testdata/2.txt")
	if err != nil {

	}

	for i := 0; i < b.N; i++ {
		Similarity(first, second, metrics.NewSorensenDice())
	}
}

func getOCRStrings(a, b string) (string, string, error) {
	ba, err := os.ReadFile(a)
	if err != nil {
		return "", "", err
	}

	bb, err := os.ReadFile(b)
	if err != nil {
		return "", "", err
	}

	return string(ba), string(bb), nil
}
