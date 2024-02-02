const moduleJoueur = require('./joueur.js');
const joueur = moduleJoueur.joueur;

// Functions 

function shuffleArray(array) {
    for (let i = array.length - 1; i > 0; i--) {
      const j = Math.floor(Math.random() * (i + 1));
      // Swap array[i] and array[j]
      [array[i], array[j]] = [array[j], array[i]];
    }
}

function init_sac(){
    let sac = []

    lettres = [{l:"A",n:14},{l:"B",n:4},{l:"C",n:7},{l:"D",n:5}
        ,{l:"E",n:19},{l:"F",n:2},{l:"G",n:4},{l:"H",n:2},{l:"I",n:11}
        ,{l:"J",n:1},{l:"K",n:1},{l:"L",n:6},{l:"M",n:5},{l:"N",n:9}
        ,{l:"O",n:8},{l:"P",n:4},{l:"Q",n:1},{l:"R",n:10},{l:"S",n:7}
        ,{l:"T",n:9},{l:"U",n:8},{l:"V",n:2},{l:"W",n:1},{l:"X",n:1}
        ,{l:"Y",n:1},{l:"Z",n:2}];

    for (let i=0;i<lettres.length;i++){
        let{l,n} = lettres[i];
        for (j=0;j<n;j++){
            sac.push(l);
        }
    }
    shuffleArray(sac);
    return sac

};
    

// start 
const sac = init_sac();

const joueur1 = new joueur(sac,0);
const joueur2 = new joueur(sac,1);

// 1er tour joueur 1 (sans jarnak)
let tour = 1;
while (tour==1){
    tour = jouer(joueur1,tour);
};

// alternance des tours
jeufini = false;
while(!(jeufini)){
    // TOUR JOUEUR 2
    tour = 2
    console.log("/n-------------/nTOUR DU JOUEUR 2")
    joueur1.afficheGrille();
    joueur1.afficheDeck();
    joueur2.afficheGrille();
    joueur2.afficheDeck();

        //joueur2.choixPioche(sac)

        //jarnak 
    console.log("Jarnak ? o/n")
    const jarnak = moduleJoueur.tapeMot()
    if (jarnak == "o"){
        
    }
    
    // TOUR JOUEUR 1

    console.log("/n-------------/nTOUR DU JOUEUR 1")
    joueur1.afficheGrille();
    joueur1.afficheDeck();
    joueur2.afficheGrille();
    joueur2.afficheDeck();

        //joueur1.choixPioche(sac)

        //jarnak

    while (tour==1){
        tour = jouer(joueur1,tour);
    };
}

function jouer(joueur,tour){
    joueur.afficheGrille();
    joueur.afficheDeck();

    let status = joueur.nouveauMot(joueur.choixLigne());
    if (status=="fini")
    {
        tour = (tour+1)%2 ;
    }else if (status=="ok"){
        joueur.jeu.push(sac.pop()); //pioche 1 lettre
    }

    return tour
}