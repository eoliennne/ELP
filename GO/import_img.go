package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
)

// fonction qui ouvre le fichier image et retourne à l'aide de la fonction decode une image
// du type image.Image du package ainsi que son type
func decodeImage(filename string) (image.Image, string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()
	img, format, err := image.Decode(file)
	if err != nil {
		return nil, "", err
	}
	return img, format, nil
}

// fonction qui permet de savoir à quel voisin d'un pixel on
// applique le flou selon le rayon souhaité
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

func main() {
	img, format, err := decodeImage("cat.png")
	if err != nil {
		fmt.Println("Erreur lors du décodage de l'image:", err)
		return
	}

	fmt.Printf("L'image décodée est au format: %s\n", format)
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	fmt.Printf("L'image est de taille: %d x %d\n", width, height)
}
