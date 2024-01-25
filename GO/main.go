package main

import (
	_ "GO/decode_img"
	"fmt"
	"image"
	"image/png"
	_ "image/png"
	"log"
	"os"
)

func update(rad, xinf, xsup, yinf, ysup int, imgNew *image.RGBA, img image.Image) {
	//lance meanPixel sur tous les pixels d'une tranche
	for x := xinf; x <= xsup; x++ {
		for y := yinf; y <= ysup; y++ {

			color := meanPixel(x-rad, x+rad, y-rad, y+rad, img)
			imgNew.Set(x, y, color)

		}
	}
}

func main() {

	// Import de l'image
	var in_file string
	fmt.Println("Quel est l'adresse du fichier que vous souhaitez flouter?")
	fmt.Scanf("%s", &in_file)

	img, width, height := DecodeImage(in_file) //ou copier coller le code d'import_img direct ?

	// Définition du rayon de flou
	var rad_percent int
	fmt.Println("A quel point souhaitez vous flouter votre image?\nEntrez un chiffre entre 1 et 100")
	for err := true; err == true; {
		fmt.Scanf("%d", &rad_percent)
		if rad_percent < 1 || rad_percent > 100 {
			fmt.Println("Erreur, veuillez entrer un nombre entre 1 et 100")
		} else {
			err = false
		}
	}

	rad := 10 //conversion entre rad_percent et rad à faire

	// Initialisation de l'image de sortie
	imgNew := image.NewRGBA(image.Rect(0, 0, width, height))

	// Création du channel d'état
	buffer := 10                   // valeur arbitraire supérieure au nombre de chan
	ch := make(chan []int, buffer) // les chan envoient un 1 lorsqu'elles ont fini leur taff

	// Lancement des goroutines
	var nb_routines_ans int
	fmt.Println("Combien de threads souhaitez vous utiliser? (2,4 ou 10)")
	for err := true; err == true; {
		fmt.Scanf("%d", &nb_routines_ans)
		if int(nb_routines_ans) != 2 || int(nb_routines_ans) != 4 || int(nb_routines_ans) != 10 {
			fmt.Println("Erreur, veuillez entrer 2, 4 ou 10")
		} else {
			err = false
		}
	}

	nb_routines := 4 //nb_routines := nb_routines_ans

	// calcul giga savant avec des divisions euclidiennes palala
	tranches := Decoupage(nb_routines, width, height)

	for t := 0; t < nb_routines; t++ {
		var xinf, xsup, yinf, ysup int = tranches[t][0], tranches[t][1], tranches[t][2], tranches[t][3]

		go update(rad, xinf, xsup, yinf, ysup, imgNew, img)

	}

	// boucle qui tourne jusqu'à ce qu'elle ai reçu les signaux de fin de toutes les routines
	for compteur := 0; compteur < nb_routines; {
		v := <-ch
		compteur += 1
	}

	// Création du png sortie
	out_file, err := os.Create("output.png")
	if err != nil {
		log.Fatal(err)
	}
	defer out_file.Close()

	// Encodage de l'image dans le png
	err = png.Encode(out_file, imgNew)
	if err != nil {
		log.Fatal(err)
	}

}
