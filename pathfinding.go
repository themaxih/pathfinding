package main

import "fmt"

/*

INTRODUCTION


En gros je veux créer un tableau qui comporte 5 tableau de 5 cases :

( 5*5 cases )

[[0, 1, 2, 3, 4]    <- Tableau 0
 [0, 1, 2, 3, 4]    <- Tableau 1
 [0, 1, 2, 3, 4]    <- Tableau 2
 [0, 1, 2, 3, 4]    <- Tableau 3
 [0, 1, 2, 3, 4]]   <- Tableau 4

Il va donc y avoir un joueur placé sur une case random d'un des 5 tableaux sauf aux endroit où il yu aura des murs, et des ennemis ( qui serront définits selon le niveau )

Les ennemis vont apparaître aussi à un certain endroit selon le niveau ( il peut y en avoir plusieurs )

Chaque cases de chaque tableau, aura un chiffrer assigné:
  - 1 pour les cases les plus proches du joueur
  - Plus on s'éloigne du joueur plus le chiffre est élevé ( cela augmente de 1 en 1 )

Il y aura aussi des lettres pour caractériser les entités :
  - P pour le joueur
  - E pour les ennemis
  - W pour les murs


PLAN


1) Définir le niveau :
  - Il y a 12 niveau de 1 à 12 ( compris )

2) Placer les murs selon le niveau :
  - Niveau 1, 5 et 9 : murs en 8, 16 et 18
  - Niveau 2, 6 et 10 : murs en 11, 14 et 16
  - Niveau 3, 7 et 11 : murs en 6, 8 et 16
  - Niveau 4, 8 et 12 : murs en 8, 10 et 13

3) Placer le joueur :
  - Toujours sur la case 20 au départ

4) Placer les ennemis, toujours à la même place selon le niveau :
  - Niveau 1 : cases 3 et 14
  - Niveau 2 : cases 2 et 9
  - Niveau 3 : case 9
  - Niveau 4 : cases 1
  - Niveau 5 : cases 1, 9 et 24
  - Niveau 6 : cases 3, 8 et 24
  - Niveau 7 : cases 7 et 19
  - Niveau 8 : cases 3 et 23
  - Niveau 9 : cases 0, 12, 19 et 22
  - Niveau 10 : cases 0, 7, 13 et 19
  - Niveau 11 : cases 1, 3 et 14
  - Niveau 12 : cases 7, 9 et 19

5) Effectuer la numérotation des cases ( qui devra toujours s'adapter selon la position du joueur, même si elle change durant le run ) :
  - Si le joueur est dans une case d’un tableau du tableau :
    - Si le joueur est dans la case 0 d’un tableau :
      - Les cases d’après seront dans l’ordre croissant à partir de 1 :
        - P, 1, 2, 3, 4
    - Si le joueur est dans la case 1,2 ou 3 d’un tableau :
      - Les cases d’après seront dans l’ordre croissant à partir de 1
		  - Les cases d’avant seront dans l’ordre décroissant jusqu’au joueur et la case d’avant le joueur sera 1 :
			  - 1, P, 1, 2, 3
        - 2, 1, P, 1, 2
        - 3, 2, 1, P, 1
    - Si le joueur est dans la case 4 d’un tableau :
		  - Les cases d’avant seront dans l’ordre décroissant jusqu’au joueur et la case d’avant le joueur sera 1 :
			  - 4, 3, 2, 1, P
  -  Si le joueur est dans un tableau :
    - Si le joueur est dans le tableau 0 :
	    - Les tableaux suivants seront dans l’ordre croissant, mais en commençant avec un chiffre de plus à chaques tableau de plus :
		    - Tableau 1 → 1, 2, 3, 4, 5
		    - Tableau 2 → 2, 3, 4, 5, 6
		    - Tableau 3 → 3, 4, 5, 6, 7
        - Tableau 4 → 4, 5, 6, 7, 8
    - Si le joueur est dans le tableau 1, 2 ou 3 :
	    - Les tableaux suivants seront dans l’ordre croissant, mais en commençant avec un chiffre de plus à chaques tableau de plus
	    - Les tableaux précédents seront dans l’ordre croissant, mais en commençant avec un chiffre de plus à chaques tableau de plus :
        - Tableau 0 → 3, 4, 5, 6, 7
		    - Tableau 1 → 2, 3, 4, 5, 6
		    - Tableau 2 → 1, 2, 3, 4, 5
	    	- Tableau 4 → 1, 2, 3, 4, 5
    - Si le joueur est dans le tableau 4 :
	    - Les tableaux précédents seront dans l’ordre croissant, mais en commençant avec un chiffre de plus à chaques tableau de plus :
		    - Tableau 0 → 4, 5, 6, 7, 8
        - Tableau 1 → 3, 4, 5, 6, 7
		    - Tableau 2 → 2, 3, 4, 5, 6
		    - Tableau 3 → 1, 2, 3, 4, 5

6) Déplacement des ennemis :
  - Un ennemi ne peut pas se déplacer sur un case où il y a soit un mur, le joueur ou un autre ennmi

*/

// Déclarer "carte" en tant qu'interface permet de poour mettre plusieurs type de valeurs à l'interieur ( dans ce cas, des string et des int )
var carte []interface{} = make([]interface{}, 5)

var niveau int = 3

var joueur string = "P"
var mur string = "W"
var ennemi string = "E"

func main() {
	GenererCarte()

	Afficher()
}

func GenererCarte() {
	// Rempli les cases des tableau de carte
	for i := 0; i < len(carte); i++ {
		carte[i] = []interface{}{0, 0, 0, 0, 0}
	}

	// Faire en sorte de pouvoir choisir avec une variable la position du joueur sur la carte

	// Ajouter le joueur au coordonnés (0,4) case 0 du tableau 4
	carte[4].([]interface{})[0] = joueur

	GenererMurs()

	SelectionNiveau()
}

func SelectionNiveau() {
	// Carte vide
	if niveau <= 0 || niveau > 12 {
		for i := 0; i < len(carte); i++ {
			carte[i] = []interface{}{0, 0, 0, 0, 0}
		}
	}

	if niveau == 1 {
		carte[0].([]interface{})[3] = ennemi
		carte[2].([]interface{})[4] = ennemi
	}

	// Changer tout les setup de spawn d'ennemi en dessous

	if niveau == 2 {
		carte[0].([]interface{})[3] = ennemi
		carte[2].([]interface{})[4] = ennemi
	}

	if niveau == 3 {
		carte[0].([]interface{})[3] = ennemi
		carte[2].([]interface{})[4] = ennemi
	}

	if niveau == 4 {
		carte[0].([]interface{})[3] = ennemi
		carte[2].([]interface{})[4] = ennemi
	}

	if niveau == 5 {
		carte[0].([]interface{})[3] = ennemi
		carte[2].([]interface{})[4] = ennemi
	}

	if niveau == 6 {
		carte[0].([]interface{})[3] = ennemi
		carte[2].([]interface{})[4] = ennemi
	}

	if niveau == 7 {
		carte[0].([]interface{})[3] = ennemi
		carte[2].([]interface{})[4] = ennemi
	}

	if niveau == 8 {
		carte[0].([]interface{})[3] = ennemi
		carte[2].([]interface{})[4] = ennemi
	}

	if niveau == 9 {
		carte[0].([]interface{})[3] = ennemi
		carte[2].([]interface{})[4] = ennemi
	}

	if niveau == 10 {
		carte[0].([]interface{})[3] = ennemi
		carte[2].([]interface{})[4] = ennemi
	}

	if niveau == 11 {
		carte[0].([]interface{})[3] = ennemi
		carte[2].([]interface{})[4] = ennemi
	}

	if niveau == 12 {
		carte[0].([]interface{})[3] = ennemi
		carte[2].([]interface{})[4] = ennemi
	}
}

func GenererMurs() {
	if niveau == 1 || niveau == 5 || niveau == 9 {
		carte[1].([]interface{})[3] = mur
		carte[3].([]interface{})[1] = mur
		carte[3].([]interface{})[3] = mur
	}

	if niveau == 2 || niveau == 6 || niveau == 10 {
		carte[2].([]interface{})[1] = mur
		carte[2].([]interface{})[4] = mur
		carte[3].([]interface{})[1] = mur
	}

	if niveau == 3 || niveau == 7 || niveau == 11 {
		carte[1].([]interface{})[3] = mur
		carte[1].([]interface{})[1] = mur
		carte[3].([]interface{})[1] = mur
	}

	if niveau == 4 || niveau == 8 || niveau == 12 {
		carte[1].([]interface{})[3] = mur
		carte[2].([]interface{})[0] = mur
		carte[2].([]interface{})[3] = mur
	}
}

func Afficher() {
	// Permet l'affichage en colonne
	for i := 0; i < len(carte); i++ {
		fmt.Printf("%v\n", carte[i])
	}
}
