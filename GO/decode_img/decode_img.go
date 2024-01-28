package decode_img

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func DecodeImage(filename string) (image.Image, string, error) {
	// Cette fonction permet de récupérer un objet de type image.Image à partir de l'adresse d'un fichier image
	// Elle a surtout servi pour les premiers essais du code mais n'est pas tellement utile dans l'application,
	// si ce n'est pour convertir les images au format jpeg/jpg en png avant de les envoyer.
	//
	// Parameters:
	//  adresse du fichier
	// Returns:
	//	(image.Image): l'objet décodé
	// 	(string): le format de l'image
	//	(error): le rapport d'erreur éventuel
	//
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
	img, format, err := DecodeImage("cat.png")
	if err != nil {
		fmt.Println("Erreur lors du décodage de l'image:", err)
		return
	}

	fmt.Printf("L'image décodée est au format: %s\n", format)
	// au final notre application ne traite que les png...
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	fmt.Printf("L'image est de taille: %d x %d\n", width, height)
}
