const moduleJoueur = require('./joueur.js');
const joueur = moduleJoueur.joueur;
const fs = require('fs');

// start 
const sac = moduleJoueur.initSac();
const logFile = fs.writeFileSync("jeu.log", "", 'utf-8');
const joueur1 = new joueur(sac,1);
const joueur2 = new joueur(sac,2);

// 1er tours 
console.log("\n-------------\nTOUR DU JOUEUR 1")
fs.appendFileSync("jeu.log","Tour du joueur 1\n",'utf-8')
let tour = 1;
while (tour==1){
    tour = moduleJoueur.jouer(joueur1,tour,false,joueur1,sac);
};

console.log("\n-------------\nTOUR DU JOUEUR 2");

fs.appendFileSync("jeu.log","Tour du joueur 2\n",'utf-8');
console.log("Jarnak ? o/n");
    const jarnak = moduleJoueur.tapeMot();
    if (jarnak == "o"){
        fs.appendFileSync("jeu.log","  Jarnak!\n",'utf-8');

        // Le tour se déroule comme si c'était celui de joueur1 
        // mais les mots sont ajoutés sur la grille du joueur2
        tour = 1;
        while (tour == 1){
            tour = moduleJoueur.jouer(joueur1,tour,true,joueur2,sac);
       }
    }
tour = 2;
while (tour==2){
    tour = moduleJoueur.jouer(joueur2,tour,false,joueur2,sac);
};

// alternance des tours
jeufini = false;
while(!(jeufini)){
    
    fs.appendFileSync("jeu.log","Tour du joueur 1\n",'utf-8');
    fs.appendFileSync("jeu.log",` Ses lettres sont : ${joueur1.jeu}\n et sa grille : ${joueur1.grille}\n`,'utf-8');
    jeufini = moduleJoueur.tourEntier(joueur1,joueur2,sac);
    fs.appendFileSync("jeu.log","Tour du joueur 2\n",'utf-8');
    fs.appendFileSync("jeu.log",` Ses lettres sont : ${joueur2.jeu}\n et sa grille : ${joueur2.grille}\n`,'utf-8');
    jeufini = jeufini || moduleJoueur.tourEntier(joueur2,joueur1,sac);
    
}

// Score

const score1 = joueur1.grille.reduce((somme,ligne) => {const carre = ligne.filter(val => val!== null).length **2;
    return somme+carre;},0)

const score2 = joueur2.grille.reduce((somme,ligne) => {const carre = ligne.filter(val => val!== null).length **2;
    return somme+carre;},0)

console.log("Joueur 1: ",score1);

console.log("Joueur 2: ",score2);

fs.appendFileSync("jeu.log",`Les scores finaux sont :\n Joueur 1 : ${score1}\n Joueur 2 : ${score2}\n`,'utf-8');

if (score1>score2){
    console.log("Le joueur 1 a gagné");
}else if (score2>score1){
    console.log("Le joueur 2 a gagné");
}else{
    console.log("Egalité")
}