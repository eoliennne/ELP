# ELM

## Description
Jeu de Guess It, où il faut deviner un mot basé sur sa définition

## Exécuter le programme

### Pré-requis
Installer http-server : `npm install -g http-server`
Lancer `http-server static -a localhost -p 5018 --cors`

### Lancement
Dans un nouveau terminal, exécuter `elm reactor`, puis ouvrir le fichier /src/Main.elm à l'adresse localhost donnée.

## Déroulement du jeu
Entrer le mot dans le champ de saisie, appuyer sur le bouton Solution pour valider le mot et afficher sa solution. Le bouton Refresh charge un nouveau mot.
Lors du lancement, le premier mot est par défaut "word". 

