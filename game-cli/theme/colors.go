package theme

import "image/color"

var (
	// Base used in the game
	Base = makeSolidRGBA(153, 91, 0)

	// Sand used to paint sand tiles
	Sand = makeSolidRGBA(255, 200, 117)

	// Contour used to render countours
	Contour = makeSolidRGBA(7, 52, 99)
)

func makeSolidRGBA(r, g, b uint8) color.RGBA {
	return color.RGBA{
		R: r,
		G: g,
		B: b,
		A: 255,
	}
}
