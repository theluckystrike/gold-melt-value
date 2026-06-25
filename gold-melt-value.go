// Package goldmeltvalue computes precious-metal melt value from karat (purity),
// weight, and a live spot price. Pure, dependency-free math — the same engine
// behind the GoldGramPrice live gold calculator at https://goldgramprice.com/.
//
// Quick example:
//
//	spotPerOz := 3983.30           // USD per troy ounce
//	grams := 10.0                   // a 10 g chain
//	karat := uint32(14)
//	usd := goldmeltvalue.MeltValueGrams(grams, karat, spotPerOz)
//	// 10 g of 14k contains ~5.83 g pure gold -> melt ~ $747.05
//	_ = goldmeltvalue.Purity(karat) // 14/24 ~= 0.5833
package goldmeltvalue

// GramsPerTroyOz is the number of grams in one troy ounce (1 oz t = 31.1034768 g).
const GramsPerTroyOz = 31.1034768

// Common karat purities used in jewelry (US).
const (
	K24 = uint32(24)
	K22 = uint32(22)
	K18 = uint32(18)
	K14 = uint32(14)
	K10 = uint32(10)
)

// Purity returns the decimal purity for a karat value (karat / 24).
// Purity(24) == 1.0 (pure), Purity(18) == 0.75, Purity(14) ~= 0.5833.
func Purity(karat uint32) float64 {
	return float64(karat) / 24.0
}

// SpotPerGram returns the spot price per gram of pure metal, derived from a
// per-troy-ounce quote.
func SpotPerGram(spotPerTroyOz float64) float64 {
	return spotPerTroyOz / GramsPerTroyOz
}

// MeltValueGrams returns the melt value (in the spot-price currency) for a
// given weight in grams, karat, and spotPerTroyOz.
//
// Melt = grams * purity * (spot / 31.1034768).
func MeltValueGrams(grams float64, karat uint32, spotPerTroyOz float64) float64 {
	return grams * Purity(karat) * SpotPerGram(spotPerTroyOz)
}

// MeltValueDwt returns the melt value for a weight in pennyweight (dwt).
// 1 dwt = 1/20 troy oz.
func MeltValueDwt(dwt float64, karat uint32, spotPerTroyOz float64) float64 {
	troyOz := dwt / 20.0
	return troyOz * Purity(karat) * spotPerTroyOz
}

// MeltValueTroyOz returns the melt value for a weight in troy ounces
// (already in troy units).
func MeltValueTroyOz(troyOz float64, karat uint32, spotPerTroyOz float64) float64 {
	return troyOz * Purity(karat) * spotPerTroyOz
}
