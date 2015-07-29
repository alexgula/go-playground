package render

import (
	"image"
	"image/color"
	"testing"
)

func BenchmarkLine(b *testing.B) {
	m := image.NewNRGBA(image.Rect(0, 0, 100, 100))
	for i := 0; i < b.N; i++ {
		for x := uint8(0); x < 100; x++ {
			for y := uint8(0); y < 100; y++ {
				Line(m, int(x), int(y), 50, 50, color.RGBA{x + y, x - y, x & y, 255})
			}
		}
	}
}
