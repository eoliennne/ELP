# GO_ELP

## Description de l'application
Notre application permet d'appliquer un flou gaussien à une image au format png, jpeg ou jpg.
Pour la lancer, il faut lancer le serveur puis lancer le client en renseignant l'image que l'on souhaite affecter et le rayon de flou (maximum 50 pixels).

## Commandes à utiliser pour tester l'application
##### Pour lancer le serveur de traitement des images
go run server.go

##### Pour traiter une image 
go run client.go nom_du_fichier.extension pixels *nom_du_fichier_de_sortie

#### Exemples : 
##### Pour appliquer un flou de rayon 20 pixels à l'image cat.png :
##### go run client.go cat.png 20

##### Pour appliquer un flou de rayon 5 pixels à l'image chat.jpg et le récupérer dans chat_flou.png : 
##### go run client.go cat.png 20 chat_flou


## Informations pour les tests
Vous pouvez lancer plusieurs clients en même dans différents terminaux.

Plusieurs images de tests sont disponibles dans la racine du projet :

l'image cat.png qui fait 775x607
l'image chat.jpg qui fait 622x311
l'image bigcat.jpeg qui fait 1200x800
l'image cut_cat.png qui fait 360x360 et a un fond transparent

##### Attention !
L'image créée par le serveur est au format png et a un nom par défaut qui est image_floue.png
Si vous lancez plusieurs clients en même temps ou à la suite sans préciser de nom, elle sera écrasée.

## Structure du code
Le code de l'application est composé d'un package main contenant le programme serveur et le programme client.
Le package decode_img contient une fonction permettant de récupérer un objet traitable depuis un jpg ou un png.
Le package serveurflou contient les fonctions de traitement de l'image ainsi que des fonctions permettant d'envoyer et recevoir une image au format png sous forme de bytes.
