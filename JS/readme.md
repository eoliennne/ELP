# JARNAK
### Règles
Notre jeu de jarnak se joue à deux joueurs sur un même terminal, chacun son tour.
Une fois le premier tour passé, a chaque fois que c'est à vous de jouer, vous avez d'abord le choix entre piocher une lettre ou en échanger trois de votre jeu contre trois de la pioche (p ou e minuscule). Vous pouvez ensuite choisir une ligne et y entrer un mot de minimum 3 lettres et maximum 3 à l'aide des lettres éventuellement déjà présentes dessys et de celles de votre jeu, autant de fois que vous le souhaitez. Si vous ne souhaitez plus entrer de mot, choisissez une ligne au hasard puis entrez 'no' comme mot.
Pour faire un jarnak, répondez oui (o) lorsque l'on vous le propose. vous pouvez alors voler autant de mots que vous voulez au joueur adverse et choisir une ligne vide sur laquelle les placer dans votre grille. Une fois que toutes les lignes d'un joueur sont remplies et que le suivant ne peut pas faire de jarnac, la partie s'arrête et le score est calculé. 

### Lancement
Pour lancer notre jeu de jarnak une fois le dépot copié, effectuez un npm install dans le dossier du jeu puis lancez le fichier main avec la commande 'node main.js'.
Vous devez disposer d'un node js pour pouvoir jouer.

### Structure du programme
Notre programme est composé d'un fichier main qui lance la partie après avoir initialisé ses variables et d'un module 'joueur.js' qui contient une classe joueur avec toutes les fonctions qui affectent l'état de la grille d'un joueur et interagisse avec lui, ainsi que des fonctions d'initialisation de la partie (création de la pioche...).
