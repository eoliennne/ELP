const prompt = require('prompt');
prompt.start();

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

function pioche(sac,joueur){
    joueur.push(sac.pop());
};



// Deroulé
const sac = init_sac();
let joueur1 = [], joueur2 = [];ligne=[];

tour = 1;
while (joueur1.length < 6) {pioche(sac,joueur1)};
tour = 2;
while (joueur2.length < 6) {pioche(sac,joueur2)};

console.log("joueur 1", joueur1,"joueur 2",joueur2);

function jouer_lettre(joueur, ligne,f){
    prompt.get(['mot'], function (err, result) {
   
    if (err) {
      console.error(err);
      return;
    }

    const mot = result.mot;
    //parcourir le mot lettre par lettre

    if (joueur.includes(lettre) || ligne.includes(lettre)) {
        f(lettre,ligne);
    }else{console.log("Tu n'as pas cette lettre");
        jouer_lettre(joueur,ligne,callback)}

  });}

function enter(mot,ligne){
    ligne.push(mot) 
    //enregistrer le mot dans fichier

}




    

