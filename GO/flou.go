package flou

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func meanPixel(xmin, xmax, ymin, ymax int, img image.Image) (colour color.RGBA) {
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

func main() {

	// Decode the cat => image imgCat
	imgCat, format, err := decodeImage("cat.png")
	if err != nil {
		fmt.Println("Erreur lors du décodage de l'image:", err)
		return
	}

	fmt.Printf("L'image décodée est au format: %s\n", format)
	width := imgCat.Bounds().Dx()
	height := imgCat.Bounds().Dy()
	fmt.Printf("L'image est de taille: %d x %d\n", width, height)

	// Low quality cat

	// create new image
	imgLQ := image.NewRGBA(image.Rect(0, 0, width, height))

	RAD := 10

	for y := 10; y < height-10; y++ {
		for x := 10; x <= width-10; x++ {
			colour := meanPixel(x-RAD, x+RAD, y-RAD, y+RAD, imgCat)
			imgLQ.Set(x, y, colour)
		}
	}

	// Create a new PNG file
	fileLQ, err := os.Create("bad_cat.png")
	if err != nil {
		log.Fatal(err)
	}
	defer fileLQ.Close()

	// Encode the image to the PNG file
	err = png.Encode(fileLQ, imgLQ)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Image created and saved as bad_cat.png")

}
