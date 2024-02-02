const prompt = require('prompt-sync')();
const fs = require('fs');

function tapeMot() {
    return prompt('>');
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

        nouveauMot(lignei,jarnak,j){
            console.log("Quel mot souhaitez vous ajouter à la grille?");
            console.log("Tapez no pour terminer votre tour.");
            const mot = tapeMot();
            const MOT = mot.toUpperCase();
            const ligne = this.grille[lignei];
            // Vérifie si le joueur veut continuer d'entrer des mots
            if (MOT =="NO"){
                return "fini";
            };

            if (mot.length<3 || mot.length>9) {
                console.error("Veuillez entrez un mot d'au moins trois lettres et de maximum 9 lettres.");
                return this.nouveauMot(lignei,jarnak,j);
            }
            console.log(MOT)
            let M_O_T = MOT.split('');

            const check = (lettre) => {
                return ligne.includes(lettre) || this.jeu.includes(lettre);
            };
            // fonction pour vérifier que le nouveau mot utilise tous les mots de la ligne
            const recycle = (lettre) => {
                if (!lettre){
                    return true;
                } else {
                return M_O_T.includes(lettre);
            }
            };
            
            let valid = M_O_T.every(check) && ligne.every(recycle);

            if (!valid){
                console.log("Le mot ne correspond pas aux lettres disponibles!\nEssaie autre chose.");
                this.nouveauMot(lignei,jarnak,j);

            } else {
                function padding(liste){
                    if (liste.length < 9){
                        liste.push(null);
                        return padding(liste);
                    } else {
                        return liste;
                    };
                }

                if (jarnak){
                    this.grille[lignei]= new Array(9).fill(null);
                    console.log("Où voulez-vous placer votre mot volé?");
                    lignei = this.choixLigne();
                    while (!this.ligneVide(j.grille[lignei])){
                        lignei= j.choixLigne();
                    };
                };
                let mot_grille = padding(M_O_T);
                j.grille[lignei] = mot_grille;
                fs.appendFileSync("jeu.log",`  Le joueur ${j.num} ajoute le mot ${MOT} à sa grille. \n`,'utf-8')

                // retire les lettres jouées de la main du joueur
                for (const lettre of M_O_T){
                    if (!ligne.includes(lettre)){
                        this.enleveLettre(lettre)
                    }
                }
                console.log("Mot ajouté à la grille.");
                return "ok";
            };
                
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
            console.log("");
        }

        choixLigne(){
            console.error("Entre un numéro de ligne entre 0 et 7");
            const num = tapeMot();
            if (isNaN(num)){
                return this.choixLigne();
            }
            return num;
        }

        tour(){  
            while (tour==this.num+1){
                let status = this.nouveauMot(this.choixLigne,jarnak,j);
                if (status=="fini")
                {
                    tour = (tour+1) % 2;
                }else if (status=="OK"){
                    this.jeu.push(sac.pop()); //pioche 1 lettre
                }
            };

        }
        choixPioche(sac){
            let choix = tapeMot();
            let echange_possible = (this.jeu.length>2);
            if (echange_possible){
                console.log("Souhaitez vous piocher une nouvelle carte ou en échanger trois de votre jeu contre trois nouvelles? (p/e)");
            } else {
                console.log("Vous n'avez pas assez de cartes pour réaliser un échange, vous piochez une nouvelle carte.");
            };

            if (!echange_possible || choix =="p"){
                this.jeu.push(sac.pop());
                this.afficheDeck()
                fs.appendFileSync("jeu.log",`  Il pioche une nouvelle lettre.\nVoici son jeu : ${this.jeu}\n`,'utf-8')
                return;
            } else if (choix=="e"){
                this.choixEchange(sac);
                fs.appendFileSync("jeu.log",`  Il échange trois lettres de son jeu contre des lettres de la pioche.\nVoici son jeu : ${this.jeu}.\n`,'utf-8')
                return;
            }
            else {
                // cas d'erreur
                this.choixPioche(sac)
            }
        }
        choixEchange(sac){
            let liste_lettres = [];
            while (liste_lettres.length<3){
                console.log("Quelle lettre voulez vous échanger?");
                let temp = tapeMot();
                let lettre = temp.toUpperCase();
                if (!this.jeu.includes(lettre)){
                    console.log("Cette lettre n'est pas dans votre main.");
                } else {
                    console.log("Très bien! nous enlevons un ",lettre);
                    liste_lettres.push(lettre);
                }};
            
            for (const enleve of liste_lettres) {
                this.enleveLettre(enleve);
                this.jeu.push(sac.pop());
            };
                
        };
        enleveLettre(lettre){
            const i = this.jeu.indexOf(lettre);
            this.jeu.splice(i, 1);
        };
        ligneVide(ligne){
            return ligne.every(element => element === null);
        };

        grillePleine()
        {
            if (this.grille.every(ligne => this.ligneVide(ligne))){
                return false;
            }else{
                return true;
            };
        };
           
};

function shuffleArray(array) {
    for (let i = array.length - 1; i > 0; i--) {
      const j = Math.floor(Math.random() * (i + 1));
      [array[i], array[j]] = [array[j], array[i]];
    }
}

function initSac(){
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
    return sac;

};

function jouer(joueur,tour,jarnak,joueurB,sac){
    // joueur est le joueur dont le deck et les mots sont utilisés
    // joueurB est le joueur qui joue et place les mots sur sa grille
    joueur.afficheGrille();
    joueur.afficheDeck();
    let status = joueur.nouveauMot(joueur.choixLigne(),jarnak,joueurB);
    if (status=="fini")
    {
        tour = (tour+1)%2 ;
    }else if (status=="ok"){
        joueur.jeu.push(sac.pop()); //pioche 1 lettre
    }
    return tour
}

function tourEntier(joueur,adversaire,sac){
    console.log("\n-------------\nTOUR DU JOUEUR",joueur.num);
    joueur.afficheGrille();
    joueur.afficheDeck();
    adversaire.afficheGrille();
    adversaire.afficheDeck();

    joueur.choixPioche(sac);
    
    console.log("Jarnak ? o/n");
    const jarnak = tapeMot();
    if (jarnak == "o"){
        fs.appendFileSync("jeu.log","  Jarnak!\n",'utf-8')

        // Le tour se déroule comme si c'était celui de joueur1 
        // mais les mots sont ajoutés sur la grille du joueur2
        tour = adversaire.num;
        while (tour == adversaire.num){
            tour = jouer(adversaire,tour,true,joueur,sac);
       };
    }

    tour = joueur.num;
    while (tour==joueur.num){
        tour = jouer(joueur,tour,false,joueur,sac);    
    };
    if (joueur.grillePleine()){
        return true;
    } else {
        return false;
    };
    
}



module.exports= {
    joueur : joueur,
    tapeMot : tapeMot,
    initSac : initSac,
    jouer : jouer,
    tourEntier : tourEntier,
};
