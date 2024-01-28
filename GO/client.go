package main

import (
	"GO/decode_img"
	"GO/serveurflou"
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	// Connexion au serveur
	network := "tcp"
	address := "localhost:8090"
	conn, err := net.Dial(network, address)
	if err != nil {
		fmt.Println("Erreur lors de la connexion au serveur:", err)
		return
	}
	defer conn.Close()

	// récupère l'argument correspondant au taux de flou et l'envoie au serveur
	flou, _ := strconv.Atoi(os.Args[2])
	if flou > 100 || flou < 1 {
		fmt.Println("Erreur, le flou demandé est invalide")
		return
	}
	EnvoiFlou(conn, flou)

	// envoie l'image à flouter sous format image.Image au serveur
	img, _, err := decode_img.DecodeImage(os.Args[1])
	if err != nil {
		fmt.Println("Erreur lors du décodage de l'image:", err)
		return
	}

	serveurflou.EnvoiImage(conn, img)
	imgNew, _ := serveurflou.Reception_img(conn)
	crea_fichier(imgNew)

}

func EnvoiFlou(conn net.Conn, in_rad int) {
	// EnvoiFlou permet de gérer l'envoi du niveau de flou demandé au serveur
	//
	// Parameters:
	// 	 conn (net.conn): la connexion au serveur
	// 	 in_rad (int): le degré de flou demandé

	_, err := io.WriteString(conn, fmt.Sprintf("%d\n", in_rad))
	if err != nil {
		fmt.Println("Erreur lors de l'envoi du flou:", err)
		return
	}
	fmt.Println("Envoi du flou terminé")
	return
}

func crea_fichier(imgNew image.Image) {
	// Cette fonction permet de créer un fichier contenant l'image floutée
	//
	// Parameters:
	// 	 imgNew (image.Image): image qui va être encodée dans le fichier créé
	//
	var filename string
	if len(os.Args) < 4 {
		filename = "image_floue.png"
	} else {
		filename = os.Args[3] + ".png"
	}

	out_file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out_file.Close()

	// Encodage de l'image dans un nouveau png
	err = png.Encode(out_file, imgNew)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Image traitée et disponible dans image_flou.png")

}
