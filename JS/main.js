const prompt = require('prompt');
prompt.start();

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

function pioche(sac,joueur){
    joueur.push(sac.pop());
};

class joueur {
	constructor(sac,num){
        this.num = num
        this.grille=[]
        for (let i = 0; i < 8; i++) {
            let lignevide = new Array(9).fill(null);
            this.grille.push(lignevide);
        }
		this.jeu=[]
        for (let i = 0; i < 6; i++) {
            this.jeu.push(sac.pop());
        };}

        nouveaumot(lignei){
            prompt.get(['mot'], function (err, result) {
                const ligne = this.grille[lignei]
                if (err) {
                  console.error(err);
                  return;
                }
                function check(lettre) {
                    return ligne.includes(lettre) || jeu.includes(lettre);}
                //const mot = result.mot;

                let valid = result.mot.split('').every(check);
              
            })
        };
        
        
        afficheChar(char){
            if (char!=null){
                process.stdout.write(char + '|');
            }else{
                process.stdout.write(' |');
            }
                    
        };

        afficheGrille(){
            console.log("\nGrille du joueur ",this.num);
            this.grille.forEach((ligne,i) => {process.stdout.write('\n'+i+'|');
                                        ligne.forEach(this.afficheChar);
                                        });

        };
        
        afficheDeck(){
            console.log("\nDeck du joueur ",this.num);
            this.jeu.forEach(l => process.stdout.write(" "+l));
        }


        //placeholders
        choixLigne(){
            return 0
        }

        tour(){  
            while (tour==this.num+1){
                let status = this.nouveaumot(this.choixLigne);
                if (status=="fini")
                {
                    tour = (tour+1) % 2;
                }else if (status=="OK"){
                    this.jeu.push(sac.pop()); //pioche 1 lettre
                }
            };

        }
           
};

            
            

// start 
const sac = init_sac();

joueur1 = new joueur(sac,0);
joueur2 = new joueur(sac,1);

// 1er tour joueur 1 (sans jarnak)
tour = 1
joueur1.afficheGrille();
joueur1.afficheDeck();
joueur2.afficheDeck();


while (tour==1){
    let status = joueur1.nouveaumot(joueur1.choixLigne);
    if (status=="fini")
    {
        tour = 2;
    }else if (status=="ok"){
        joueur1.jeu.push(sac.pop()); //pioche 1 lettre
    }
};

// alternance des tours
jeufini = false;
while(not(jeufini)){
    joueur2.tour()
    joueur1.tour()
}

//joueur 2 => jarnak
//            tire une lettre ou echange 3 de ses lettres
//            joue jusqu'à fini
//pioche 1 lettre si nouveau mot
//continue jusqu'à fini