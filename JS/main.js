const moduleJoueur = require('./joueur.js');
const joueur = moduleJoueur.joueur;
const fs = require('fs');

// start 
const sac = moduleJoueur.initSac();
const logFile = fs.writeFileSync("jeu.log", "", 'utf-8');
const joueur1 = new joueur(sac,1);
const joueur2 = new joueur(sac,2);

// 1er tours 
fs.appendFileSync("jeu.log","Tour du joueur 1\n",'utf-8')
let tour = 1;
while (tour==1){
    tour = moduleJoueur.jouer(joueur1,tour,false,joueur1,sac);
};

fs.appendFileSync("jeu.log","Tour du joueur 2\n",'utf-8');
console.log("Jarnak ? o/n");
    const jarnak = tapeMot();
    if (jarnak == "o"){
        fs.appendFileSync("jeu.log","  Jarnak!\n",'utf-8');

        // Le tour se déroule comme si c'était celui de joueur1 
        // mais les mots sont ajoutés sur la grille du joueur2
        tour = 1;
        while (tour == 1){
            tour = jouer(joueur1,tour,true,joueur2,sac);
       }
    }
tour = 2;
while (tour==2){
    tour = moduleJoueur.jouer(joueur2,tour,false,joueur2,sac);
};

// alternance des tours
jeufini = false;
while(!(jeufini)){
    
    fs.appendFileSync("jeu.log","Tour du joueur 2\n",'utf-8');
    moduleJoueur.tourEntier(joueur2,joueur1,sac);
    moduleJoueur.tourEntier(joueur1,joueur2,sac);

    if (joueur1.grillePleine()) {
        jeufini = true;
        console.log("Le joueur 1 a gagné");
    }
    else if (joueur2.grillePleine()) {
        jeufini = true;
        console.log("Le joueur 2 a gagné");
    }
    
    
}
