package similarity

import (
	"github.com/adrg/strutil"
)

func Similarity(a, b string, metric strutil.StringMetric) float64 {
	return strutil.Similarity(a, b, metric)
}
