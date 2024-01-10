package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
)

func main() {

	// Import de l'image
	var in_file string
	fmt.Println("Quel est l'adresse du fichier que vous souhaitez flouter?")
	fmt.Scanf("%s",&in_file)

	// Définition du rayon de flou
	var rad_percent string
	fmt.Println("A quel point souhaitez vous flouter votre image?\nEntrez un chiffre entre 1 et sdxcgvdx")
	for err:=true;err==true{
		fmt.Scanf("%s",&rad_percent)
		if int(rad_percent)<1 || int(rad_percent)>100{
			fmt.Println("Erreur, veuillez entrer un nombre entre 1 et dsgxvwdcx")
		}else{
			err=false
		}
	}

	// calcul giga savant avec des divisions euclidiennes palala


	// Initialisation de l'image de sortie
	imgNew := image.NewRGBA(image.Rect(0, 0, width, height))


	// Création du channel d'état
	buffer := 10  						// valeur arbitraire supérieure au nombre de chan
	ch := make(chan []int,buffer) 		// les chan envoient un 1 lorsqu'elles ont fini leur taff


	// Lancement des goroutines
	var nb_routines_ans string
	fmt.Println("Combien de threads souhaitez vous utiliser? (2,4 ou 10)")
	for err:=true;err==true{
		fmt.Scanf("%s",&nb_routines_ans)
		if int(nb_routines_ans)!= 2 || int(nb_routines_ans)!= 4 || int(nb_routines_ans)!= 10
			fmt.Println("Erreur, veuillez entrer 2, 4 ou 10")
		}else{
			err=false
		}
	}



	nb_routines = 4




	// à mettre dans la fonction pour updater l'image
	imgNew.Set(x, y, color)

	// boucle qui tourne jusqu'à ce qu'elle ai reçu les signaux de fin de toutes les routines 
	for compteur:=0;compteur<nb_routines{
		v := <-ch
		compteur+=1
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