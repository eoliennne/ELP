package main

import (
	"GO/serveurflou"
	"bufio"
	"fmt"
	"image"
	"net"
)

func main() {
	// Attente de connexion
	network := "tcp"
	address := "localhost:8090"
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
		go Session(conn)

	}
}

func Reception_flou(conn net.Conn) int {
	// Reception_flou permet d'obtenir l'information sur à quel point le client veut flouter son image
	//
	// Parameters:
	//   conn (net.Conn): la connexion définie avec le client courant
	//
	// Returns:
	//   int: l'indicateur sur le niveau de flou
	reader := bufio.NewReader(conn)
	flou, err := reader.ReadString('\n')
	fmt.Println("flou:", flou)
	if err != nil {
		fmt.Println("Erreur lors de la réception du flou:", err)
		return 0
	}

	fmt.Println("Flou demandé par le client:", flou)

	return 2
}

func Session(conn net.Conn) {
	// Session permet de gérer l'usage de l'application pour chacun des clients
	//
	// Parameters:
	//   conn (net.Conn): la connexion définie avec le client courant
	//
	defer conn.Close()

	fmt.Println("Session ouverte pour le client:", conn.RemoteAddr())
	rad := Reception_flou(conn)

	img, _ := serveurflou.Reception_img(conn)
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	imgNew := image.NewRGBA(image.Rect(0, 0, width, height))

	// Lancement des goroutines et création du channel vérifiant leur état
	// les routines envoient un au channel 1 lorsqu'elles ont fini leur travail
	nb_routines := 2          // runtime.NumCPU()
	buffer := nb_routines + 2 // au cas où
	ch := make(chan int, buffer)

	// défintion des parties de l'image qui vont chacune être traitées dans une goroutine
	tranches := serveurflou.Decoupage(nb_routines, width, height)

	fmt.Println("Lancement des goroutines")
	// Lancement des goroutines pour chaque tranche
	for t := 0; t < nb_routines; t++ {
		var xinf, xsup, yinf, ysup int = tranches[t][0], tranches[t][1], tranches[t][2], tranches[t][3]
		go serveurflou.Update(rad, xinf, xsup, yinf, ysup, imgNew, img, ch)

	}
	fmt.Println("attente des goroutines")
	// boucle qui tourne jusqu'à ce qu'elle ait reçu les signaux de fin de toutes les routines
	for compteur := 0; compteur < nb_routines; {
		_ = <-ch
		compteur += 1
	}
	fmt.Println("Fin des goroutines")
	// Envoi de l'image floue au client
	serveurflou.EnvoiImage(conn, imgNew)

}
