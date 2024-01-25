package main

import (
	"GO/decode_img"
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"net"
	"os"
	"strconv"
)

func main() {
	// Connexion au serveur
	network := "tcp"
	address := "localhost:8080"
	conn, err := net.Dial(network, address)
	if err != nil {
		fmt.Println("Erreur lors de la connexion au serveur:", err)
		return
	}
	defer conn.Close()
	flou, _ := strconv.Atoi(os.Args[2])
	fmt.Println("flou:", flou)
	Choix_flou(conn, flou)
	EnvoiImage(conn, os.Args[1])

}

func EnvoiImage(conn net.Conn, filename string) {
	// EnvoiImage permet de gérer l'envoi de l'image que le client veut flouter au serveur application
	//
	// Parameters:
	// 	 conn (net.conn): la connexion au serveur
	// Import de l'image
	var imag image.Image
	var err error
	for valid := true; valid == true; {
		imag, _, err = decode_img.DecodeImage(filename)
		if err != nil {
			fmt.Println("Erreur lors du chargement de l'image.", err)
		} else {
			valid = false
		}
	}

	buffer := new(bytes.Buffer)
	err = png.Encode(buffer, imag)
	if err != nil {
		fmt.Println("Erreur lors de l'encodage de l'image:", err)
		return
	}

	reader := bytes.NewReader(buffer.Bytes())
	_, err = io.Copy(conn, reader)
	if err != nil {
		fmt.Println("Erreur lors de l'envoi du fichier:", err)
		return
	}

	fmt.Println("Envoi du fichier terminé.")
}

func Choix_flou(conn net.Conn, in_rad int) {
	// EnvoiFlou permet de gère l'envoi du niveau de flou demandé au serveur
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
