package bord

import "image"

func Bord(x, y, r int, img image.Image) (xmin, xmax, ymin, ymax int) {

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	if a := x - r; a < 0 {
		xmin = 0
	} else {
		xmin = a
	}

	if a := x + r; a > width {
		xmax = width
	} else {
		xmax = a
	}

	if a := y - r; a < 0 {
		ymin = 0
	} else {
		ymin = a
	}

	if a := y + r; a > height {
		ymax = height
	} else {
		ymax = a
	}
	return
}
