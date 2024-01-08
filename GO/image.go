package main

import (
	"fmt"
	"image"
	_ "image/jpg"
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

func main() {
	img, format, err := decodeImage("chat.jpg")
	if err != nil {
		fmt.Println("Erreur lors du décodage de l'image:", err)
		return
	}

	fmt.Printf("L'image décodée est au format: %s\n", format)
	fmt.Println(img)
}
