package goldmeltvalue

import (
	"math"
	"testing"
)

func approx(a, b, tol float64) bool {
	return math.Abs(a-b) < tol
}

func TestPureGoldOneGramAtSpot(t *testing.T) {
	// 1 g of 24k at $3983.30/oz ~= $128.066
	v := MeltValueGrams(1.0, 24, 3983.30)
	if !approx(v, 128.066, 0.001) {
		t.Fatalf("got %v, want ~128.066", v)
	}
}

func TestTenGrams14k(t *testing.T) {
	// 10 g of 14k at $3983.30/oz ~= $747.05
	v := MeltValueGrams(10.0, 14, 3983.30)
	if !approx(v, 747.05, 0.05) {
		t.Fatalf("got %v, want ~747.05", v)
	}
}

func TestPennyweightMatchesTroy(t *testing.T) {
	spot := 3983.30
	// 20 dwt == 1 troy oz of 24k
	fromDwt := MeltValueDwt(20.0, 24, spot)
	fromOz := MeltValueTroyOz(1.0, 24, spot)
	if !approx(fromDwt, fromOz, 1e-6) {
		t.Fatalf("dwt %v != troy %v", fromDwt, fromOz)
	}
}

func TestPurityTable(t *testing.T) {
	if Purity(24) != 1.0 {
		t.Fatalf("Purity(24) = %v, want 1.0", Purity(24))
	}
	if !approx(Purity(18), 0.75, 1e-9) {
		t.Fatalf("Purity(18) = %v, want 0.75", Purity(18))
	}
	if !approx(Purity(14), 0.5833333, 1e-6) {
		t.Fatalf("Purity(14) = %v, want ~0.5833", Purity(14))
	}
}

func TestSpotPerGram(t *testing.T) {
	v := SpotPerGram(3983.30)
	if !approx(v, 128.066, 0.001) {
		t.Fatalf("SpotPerGram = %v, want ~128.066", v)
	}
}

func TestKaratConstants(t *testing.T) {
	if K24 != 24 || K22 != 22 || K18 != 18 || K14 != 14 || K10 != 10 {
		t.Fatalf("karat constants wrong: %d %d %d %d %d", K24, K22, K18, K14, K10)
	}
}
