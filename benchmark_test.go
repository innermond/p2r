package pange

import "testing"

var seltxt = "1-1000"

func BenchmarkSelection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sel := Selection(seltxt)
		ss, _ := sel.Split()

		for _, ee := range ss {
			for i := ee.A; i <= ee.Z; i++ {
			}
		}
	}
}
