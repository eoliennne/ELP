package decode_img

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
)


func DecodeImage(filename string) (image.Image, string, error) {
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

func main() {
	img, format, err := DecodeImage("cat.png")
	if err != nil {
		fmt.Println("Erreur lors du décodage de l'image:", err)
		return
	}

	fmt.Printf("L'image décodée est au format: %s\n", format)
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	fmt.Printf("L'image est de taille: %d x %d\n", width, height)
}
