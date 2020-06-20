package theme

import "image/color"

var (
	// Base used in the game
	Base = makeSolidRGBA(153, 91, 0)

	// Sand used to paint sand tiles
	Sand = makeSolidRGBA(255, 200, 117)

	// Contour used to render countours
	Contour = makeSolidRGBA(7, 52, 99)

	// PlayerBody contains the collor used to fill the player's body
	PlayerBody = makeSolidRGBA(0, 118, 15)

	// PlayerHighlight contains the color used to highlight the players body
	PlayerHighlight = makeSolidRGBA(7, 150, 25)
)

func makeSolidRGBA(r, g, b uint8) color.RGBA {
	return color.RGBA{
		R: r,
		G: g,
		B: b,
		A: 255,
	}
}
