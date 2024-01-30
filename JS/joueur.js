const prompt = require('prompt-sync')();

function tapeMot() {
    return prompt('.');
}

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
            const mot = tapeMot();
            const MOT = mot.toUpperCase();
            const ligne = this.grille[lignei];

            // Vérifie si le joueur veut continuer d'entrer des mots
            if (MOT =="NO"){
                return "fini";
            };

            if (mot.length<3 || mot.length>9) {
                console.error("Veuillez entrez un mot d'au moins trois lettres et de maximum 9 lettres.");
                return this.nouveaumot(lignei);
            }

            const check = (lettre) => {
                return ligne.includes(lettre) || this.jeu.includes(lettre);
            };

            let valid = MOT.split('').every(check);
            console.log(valid)

            if (!valid){
                console.log("Le mot ne correspond pas aux lettres disponibles!\nEssaie autre chose.");
                this.nouveaumot(lignei);

            } else {
                console.log("mot ajouté à la grille")
                this.grille[lignei] = MOT.split('')
                console.log(this.grille[lignei])
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
        }


        //placeholders
        choixLigne(){
            console.error("Entre un numéro de ligne entre 0 et 7");
            const num = tapeMot();
            if (isNaN(num)){
                console.error("Entre un numéro de ligne entre 0 et 7");
                this.choixLigne()
            }
            console.log(num);
            return num;
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
        choix_pioche(sac){
            console.log("Souhaitez vous piocher une nouvelle carte ou en échanger trois de votre jeu contre trois nouvelles? (p/e)")
            choix = tapeMot();
            if (choix =="p"){
                this.jeu.push(sac.pop());
                this.afficheDeck()
                return;
            } else if (choix=="e"){
                // à faire choix_echange(){}
                this.jeu.push(sac.pop());

            } else {
                this.choix_pioche(sac)
            }
        }
        choix_echange(){}
           
};

module.exports= joueur;