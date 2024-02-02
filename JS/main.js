const moduleJoueur = require('./joueur.js');
const joueur = moduleJoueur.joueur;
const fs = require('fs');

// Fonctions 


    

// start 
const sac = moduleJoueur.init_sac();
const logFile = fs.writeFileSync("jeu.log", "", 'utf-8');
const joueur1 = new joueur(sac,1);
const joueur2 = new joueur(sac,2);

// 1er tour joueur 1 (sans jarnak)
fs.appendFileSync("jeu.log","Tour du joueur 1\n",'utf-8')
let tour = 1;
while (tour==1){
    tour = moduleJoueur.jouer(joueur1,tour,false,joueur1,sac);
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
