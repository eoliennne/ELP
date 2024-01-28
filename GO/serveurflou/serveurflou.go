package serveurflou

import (
	_ "GO/decode_img"
	"bytes"
	"encoding/gob"
	"fmt"
	"image"
	"image/color"
	"image/png"
	_ "image/png"
	"net"
)

func EnvoiImage(conn net.Conn, imag image.Image) {
	// EnvoiImage permet de gérer l'envoi d'une image
	//
	// Parameters:
	//   conn (net.Conn): la connexion client/serveur

	var buf bytes.Buffer
	err := png.Encode(&buf, imag)
	if err != nil {
		fmt.Println("Erreur lors de l'encodage de l'image:", err)
		return
	}

	encoder := gob.NewEncoder(conn)
	if err := encoder.Encode(buf.Bytes()); err != nil {
		fmt.Println("Erreur lors de l'envoi de l'image:", err)
		return
	}

	fmt.Println("Envoi de l'image terminé.")
}

func Reception_img(conn net.Conn) (image.Image, error) {
	// Reception permet de gérer la réception de l'image envoyée par le client
	//
	// Parameters:
	//   conn (net.Conn): la connexion définie avec le client courant
	//
	// Returns:
	//   image.Image : l'image reçue  au format image.Image
	//
	var buf []byte
	decoder := gob.NewDecoder(conn)
	if err := decoder.Decode(&buf); err != nil {
		fmt.Println("Erreur lors de la réception de l'image:", err)
		return nil, err
	}

	fmt.Println("Image Reçue.")

	img, _, err := image.Decode(bytes.NewReader(buf))
	if err != nil {
		fmt.Println("Erreur lors du décodage de l'image:", err)
		return nil, err
	}

	return img, nil
}

func Update(rad, xinf, xsup, yinf, ysup int, imgNew *image.RGBA, img image.Image, ch chan int) {
	//Update est une fonction qui lance la fonction meanPixel (pixel prend la moyenne de couleurs des pixels environnant)
	// sur tous les pixels d'une tranche de l'image
	//
	// Parameters:
	//   rad (int): le rayon de flou souhaité
	//   xinf (int): la borne x inférieure de la tranche
	//	 xsup (int): la borne x supérieure de la tranche
	//	 yinf (int): la borne y inférieure de la tranche
	//	 ysup (int): la borne y supérieure de la tranche
	//	 imgNew (*image.RBA) : l'image contenant les pixels traités
	//	 img (image.Image) : l'image que l'on souhaite flouter
	//	 ch (chan int) : channel utilisée pour évaluer l'avancement des différentes routines
	//
	for x := xinf; x <= xsup; x++ {
		for y := yinf; y <= ysup; y++ {
			var xmin, xmax, ymin, ymax = Bord(x, y, rad, img)

			color := meanPixel(xmin, xmax, ymin, ymax, img)
			imgNew.Set(x, y, color)

		}
	}
	ch <- 1
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
	//   [][]int : tableau d'entier contenant xinf, xsup, yinf, ysup les bornes de la tranche
	//

	var list [][]int

	if width >= height {

		//découpage en colonnes
		bande := width/n - 1
		reste := width % n
		var min, max int = 0, bande

		for i := 0; i < n-1; i++ {
			var arr []int
			arr = append(arr, min, max, 0, height-1)
			list = append(list, arr)
			min, max = max+1, max+bande+1
		}
		var arr []int
		arr = append(arr, min, max+reste, 0, height-1)
		list = append(list, arr)
	}
	if height > width {

		//découpage en lignes
		bande := height/n - 1
		reste := height % n
		var min, max int = 0, bande

		for i := 0; i < n-1; i++ {
			var arr []int
			arr = append(arr, 0, width-1, min, max)
			list = append(list, arr)
			min, max = max+1, max+bande+1
		}
		var arr []int
		arr = append(arr, 0, width-1, min, max+reste)
		list = append(list, arr)
	}

	return list

}

func meanPixel(xmin, xmax, ymin, ymax int, img image.Image) (colour color.RGBA) {

	// meanPixel est la fonction qui permet pour un pixel de l'image de
	// faire la moyenne des composantes des pixels autour dans un rayon défini
	//
	// Parameters:
	//   xmin (int): abscisse minimal des pixels utilisés dans la moyenne
	//   xmax (int): abscisse maximal des pixels utilisés dans la moyenne
	//	 ymin (int): ordonnée minimale des pixels utilisés dans la moyenne
	//   ymax (int): ordonnée maximale des pixels utilisés dans la moyenne
	//	 img (image.Image): l'image à laquelle appartient le pixel
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
	// convuint8 fait la conversion d'un uint32 en un uint8
	//
	// Parameters:
	//  v (uint32): l'uint32 que l'on souhaite convertir
	//
	// Returns:
	//   (uint8) converti

	res := float64(v) * 255.0 / 0xffff
	return uint8(res)

}

func Bord(x, y, r int, img image.Image) (xmin, xmax, ymin, ymax int) {
	// Bord est la fonction qui permet de délimiter un carré englobant des pixels autour d'un pixel centre
	// selon un rayon défini, et s'adaptant aux bords de l'image
	//
	// Parameters:
	//   x,y (int): coordonnées du pixel centre
	//	 r (int): distance entre le pixel centre et ses côtés
	//	 img (image.Image): l'image à laquelle appartient le pixel
	//
	//
	// Returns:
	//   color.RGBA: les composantes couleurs et alpha du pixel une fois le flou calculé

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
