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