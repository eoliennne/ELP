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

	flou, _ := strconv.Atoi(os.Args[2])
	Choix_flou(conn, flou)

	img, _, err := decode_img.DecodeImage(os.Args[1])
	if err != nil {
		fmt.Println("Erreur lors du décodage de l'image:", err)
		return
	}
	serveurflou.EnvoiImage(conn, img)
	imgNew, _ := serveurflou.Reception_img(conn)
	crea_fichier(imgNew)

}

func Choix_flou(conn net.Conn, in_rad int) {
	// EnvoiFlou permet de gérer l'envoi du niveau de flou demandé au serveur
	//
	// Parameters:
	// 	 conn (net.conn): la connexion au serveur
	// 	 in_rad (int): le degré de flou demandé

	//writer := bufio.NewWriter(conn)
	_, err := io.WriteString(conn, fmt.Sprintf("%d\n", in_rad))
	if err != nil {
		fmt.Println("Erreur lors de l'envoi du flou:", err)
		return
	}
	fmt.Println("Envoi du flou terminé")
	return
}

func crea_fichier(imgNew image.Image) {
	// Création du png sortie
	out_file, err := os.Create("image_floue.png")
	if err != nil {
		log.Fatal(err)
	}
	defer out_file.Close()

	// Encodage de l'image dans un png
	err = png.Encode(out_file, imgNew)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Image traitée et disponible dans image_flou.png")

}
