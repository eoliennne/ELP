const joueur = require('./joueur.js');

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

joueur1 = new joueur(sac,0);
joueur2 = new joueur(sac,1);

// 1er tour joueur 1 (sans jarnak)
tour = 1
while (tour==1){
    joueur1.afficheGrille();
    joueur1.afficheDeck();
    let status = joueur1.nouveaumot(joueur1.choixLigne());
    if (status=="fini")
    {
        tour = 2;
    }else if (status=="ok"){
        joueur1.jeu.push(sac.pop()); //pioche 1 lettre
    }
};

// alternance des tours
jeufini = false;
while(!(jeufini)){
    tour = 2
    joueur1.afficheGrille();
    joueur2.afficheGrille();
    joueur1.afficheDeck();
    joueur2.afficheDeck();

    //choixPioche() //choisi de tirer une lettre ou d'en échanger 3
    //jarnak

    //tour normal
    while (tour==2){
        joueur2.afficheGrille();
        joueur2.afficheDeck();

        let status = joueur2.nouveaumot(joueur2.choixLigne());
        if (status=="fini")
        {
            tour = 1;
        }else if (status=="ok"){
            joueur1.jeu.push(sac.pop()); //pioche 1 lettre
        }
    };
    
    while (tour==1){
        joueur1.afficheGrille();
        joueur1.afficheDeck();
        let status = joueur1.nouveaumot(joueur1.choixLigne());
        if (status=="fini")
        {
            tour = 2;
        }else if (status=="ok"){
            joueur1.jeu.push(sac.pop()); //pioche 1 lettre
        }
    };
}

//joueur 2 => jarnak
//            tire une lettre ou echange 3 de ses lettres
//            joue jusqu'à fini
//pioche 1 lettre si nouveau mot
//continue jusqu'à fini