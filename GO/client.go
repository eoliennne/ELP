package main

import (
	"GO/decode_img"
	"bytes"
	"fmt"
	"image/png"
	"io"
	"net"
)

func main() {
	// Connexion au serveur
	network := "tcp"
	address := "localhost:8080"
	conn, err := net.Dial(network, address)
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	defer conn.Close()

	img, format, err := decode_img.DecodeImage("cat.png")

	if err != nil {
		fmt.Println("Erreur lors du décodage de l'image:", err)
		return
	}

	var imgBytes []byte
	buffer := new(bytes.Buffer)
	err := png.Encode(buffer, img)
	if err != nil {
		fmt.Println("Error encoding image:", err)
		return
	}

	// Copy data from the image to the connection
	_, err = io.Copy(conn, file)
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}

	fmt.Println("File sent successfully.")

}
