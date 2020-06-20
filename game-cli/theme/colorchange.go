package theme

import (
	"image/color"

	colorful "github.com/lucasb-eyer/go-colorful"
)

// Darken the input color by the given ammount
func Darken(srcColor color.Color, ammount float64) color.Color {
	auxColor, ok := colorful.MakeColor(srcColor)
	if !ok {
		return srcColor
	}
	h, c, l := auxColor.Hcl()
	l *= ammount
	return colorful.Hcl(h, c, l)
}
