package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"net"
	"os"
)

func main() {
	// Attente de connexion
	network := "tcp"
	address := "localhost:8080"
	listener, err := net.Listen(network, address)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du serveur:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Serveur à l'écoute sur le port ", address)

	for {
		// Acceptation de la demande de connexion d'un client
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erreur lors de l'acceptation d'u client:", err)
			continue
		}
		fmt.Println("connexion établie avec le client: ", conn.RemoteAddr())
		// Création d'une goroutine pour le client
		//go session(conn)
		session(conn)

		fmt.Println("helloooo")
	}
}

func session(conn net.Conn) {
	defer conn.Close()

	fmt.Println("we are here")

	received := new(bytes.Buffer)
	_, err := io.Copy(received, conn)

	if err != nil {
		fmt.Println("Error copying data:", err)
		return

		fmt.Println("File received successfully.")

	}

	img, _, err := image.Decode(received)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	// Create a new file to save the decoded image
	outputFile, err := os.Create("decoded_image.png")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Save the decoded image to the file
	err = png.Encode(outputFile, img)
	if err != nil {
		fmt.Println("Error encoding image to file:", err)
		return
	}

	fmt.Println("Image successfully decoded and saved.")
}
