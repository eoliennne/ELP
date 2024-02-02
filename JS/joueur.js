const prompt = require('prompt-sync')();
const fs = require('fs');

function tapeMot() {
    return prompt('>');
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

        nouveauMot(lignei,jarnak,sac,j){
                       
            const mot = tapeMot();
            const MOT = mot.toUpperCase();
            const ligne = this.grille[lignei];

            // Vérifie si le joueur veut continuer d'entrer des mots
            if (MOT =="NO"){
                return "fini";
            };

            if (mot.length<3 || mot.length>9) {
                console.error("Veuillez entrez un mot d'au moins trois lettres et de maximum 9 lettres.");
                return this.nouveauMot(lignei);
            }

            const check = (lettre) => {
                return ligne.includes(lettre) || this.jeu.includes(lettre);
            };
            let M_O_T = MOT.split('')
            let valid = M_O_T.every(check);
            console.log(valid)

            if (!valid){
                console.log("Le mot ne correspond pas aux lettres disponibles!\nEssaie autre chose.");
                this.nouveauMot(lignei);

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
                    lignei = j.choixLigne();
                }
                let mot_grille = padding(M_O_T)
                j.grille[lignei] = mot_grille
                console.log(j.grille[lignei])
                fs.appendFileSync("jeu.log",`Le joueur ${j.num} ajoute le mot ${MOT} à sa grille. \n`,'utf-8')

                // retire les lettres jouées de la main du joueur
                for (const lettre of M_O_T){
                    if (!ligne.includes(lettre)){
                        this.remplaceLettre(lettre,sac)
                    }
                }
                console.log("Mot ajouté à la grille.")
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
                console.error("Entre un numéro de ligne entre 0 et 7");
                this.choixLigne()
            }
            console.log(num);
            return num;
        }

        tour(){  
            while (tour==this.num+1){
                let status = this.nouveauMot(this.choixLigne);
                if (status=="fini")
                {
                    tour = (tour+1) % 2;
                }else if (status=="OK"){
                    this.jeu.push(sac.pop()); //pioche 1 lettre
                }
            };

        }
        choixPioche(sac){
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
        choixEchange(sac){
            let liste_lettres = [];
            while (liste_lettres.length<3){
                console.log("Quelle lettre voulez vous échanger?");
                let temp = tapeMot();
                let lettre = temp.toUpperCase();
                if (!this.jeu.includes(lettre)){
                    console.log("Cette lettre n'est pas dans votre main.")
                } else {
                    console.log("Très bien! nous enlevons un ",lettre);
                    liste_lettres.push(lettre);
                }}
            
            for (const enleve of liste_lettres) {
                this.remplaceLettre(lettre,sac);
            };
                
        };
        remplaceLettre(lettre,sac){
            const i = this.jeu.indexOf(lettre);
            this.jeu.splice(i, 1);
            this.jeu.push(sac.pop())
        };
        
           
};

module.exports= {
    joueur : joueur,
    tapeMot : tapeMot
};
