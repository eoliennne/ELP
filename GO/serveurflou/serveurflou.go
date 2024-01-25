package serveurflou

import (
	_ "GO/decode_img"
	"image"
	"image/color"
	_ "image/png"
)

func Update(rad, xinf, xsup, yinf, ysup int, imgNew *image.RGBA, img image.Image) {
	//lance meanPixel sur tous les pixels d'une tranche
	for x := xinf; x <= xsup; x++ {
		for y := yinf; y <= ysup; y++ {

			color := meanPixel(x-rad, x+rad, y-rad, y+rad, img)
			imgNew.Set(x, y, color)

		}
	}
}

func Decoupage(n, width, height int) [][]int {
	// Decoupage est une fonction qui permet de découper l'image en différentes zones
	// pour que chacune puisse être traiter par une goroutine
	//
	// Parameters:
	//   n (int): le nombre de tranches qui vont être crées
	//   width (int): la largeur (pixel) de l'image à découper
	//	 height (int): la hauteur (pixel) de l'image à découper
	//
	// Returns:
	//   [][]int : tableau d'entier correspondant aux coordonnées de jsplusquoi
	//

	var list [][]int //ligne = tranche, list[] = [xinf, xsup, yinf, ysup]

	if width >= height {
		//découpage en colonnes

		bande := width/n - 1 //faire division euclidienne
		reste := width % n
		var min, max int = 0, bande

		for i := 0; i < n-1; i++ {
			var arr []int
			arr = append(arr, min, max, 0, height-1)
			list = append(list, arr)
			min, max = max+1, max+bande+1
			//print(list[i])
		}
		var arr []int
		arr = append(arr, min, max+reste)
		list = append(list, arr)
	}
	if height > width {
		//découpage en lignes
		bande := height/n - 1 //faire division euclidienne
		reste := height % n
		var min, max int = 0, bande

		for i := 0; i < n-1; i++ {
			var arr []int
			arr = append(arr, 0, width-1, min, max)
			list = append(list, arr)
			min, max = max+1, max+bande+1
			print(list[i])
		}
		var arr []int
		arr = append(arr, min, max+reste)
		list = append(list, arr)
	}

	return list

}

func meanPixel(xmin, xmax, ymin, ymax int, img image.Image) (colour color.RGBA) {

	// meanPixel est la fonction qui permet de calculer pour un pixel de l'image
	// ses composantes une fois le flou appliqué
	//
	// Parameters:
	//   xmin (int): définition du rayon max considéré pour le calcul du flou
	//   xmax (int): -
	//	 ymin (int): -
	//   ymax (int): -
	//
	//
	// Returns:
	//   color.RGBA: les composantes couleurs et alpha du pixel une fois le flou calculé

	var r, g, b, a uint32 = 0, 0, 0, 0

	var NB_PIXELS uint32 = 0

	for xtemp := xmin; xtemp <= xmax; xtemp++ {
		for ytemp := ymin; ytemp <= ymax; ytemp++ {
			colortemp := img.At(xtemp, ytemp)
			rtemp, gtemp, btemp, atemp := colortemp.RGBA()
			r += rtemp
			g += gtemp
			b += btemp
			a += atemp

			NB_PIXELS += 1
		}
	}
	r = r / NB_PIXELS
	g = g / NB_PIXELS
	b = b / NB_PIXELS
	a = a / NB_PIXELS

	colour = color.RGBA{convuint8(r), convuint8(g), convuint8(b), convuint8(a)}
	return
}

func convuint8(v uint32) uint8 {
	res := float64(v) * 255.0 / 0xffff
	return uint8(res)

}

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
