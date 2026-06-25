# gold-melt-value

[![Go Reference](https://pkg.go.dev/badge/github.com/theluckystrike/gold-melt-value.svg)](https://pkg.go.dev/github.com/theluckystrike/gold-melt-value)

Compute precious-metal **melt value** from karat (purity), weight, and a live
spot price. Pure, dependency-free Go math — the same engine behind the
[GoldGramPrice](https://goldgramprice.com/) live gold gram price calculator.

## Install

```sh
go get github.com/theluckystrike/gold-melt-value
```

## Quick example

```go
package main

import (
	"fmt"

	goldmeltvalue "github.com/theluckystrike/gold-melt-value"
)

func main() {
	spotPerOz := 3983.30 // USD per troy ounce
	grams := 10.0        // a 10 g chain
	karat := uint32(14)

	usd := goldmeltvalue.MeltValueGrams(grams, karat, spotPerOz)
	// 10 g of 14k contains ~5.83 g pure gold -> melt ~ $747.05
	fmt.Println(usd)

	fmt.Println(goldmeltvalue.Purity(karat)) // 14/24 ~= 0.5833
}
```

## API

| Function | Description |
| --- | --- |
| `Purity(karat uint32) float64` | Decimal purity (karat / 24). |
| `SpotPerGram(spotPerTroyOz float64) float64` | Spot price per gram of pure metal. |
| `MeltValueGrams(grams, karat, spotPerTroyOz)` | Melt value for a weight in grams. |
| `MeltValueDwt(dwt, karat, spotPerTroyOz)` | Melt value for a weight in pennyweight (1 dwt = 1/20 troy oz). |
| `MeltValueTroyOz(troyOz, karat, spotPerTroyOz)` | Melt value for a weight in troy ounces. |

Constants: `GramsPerTroyOz = 31.1034768`, and karat helpers `K24`, `K22`, `K18`,
`K14`, `K10`.

## License

MIT.

## Links

- Live calculator: <https://goldgramprice.com/>
- Package docs: <https://pkg.go.dev/github.com/theluckystrike/gold-melt-value>
