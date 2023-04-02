package main

import (
	"fmt"
	"strconv"
)

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

var joueur string = "P"
var mur string = "W"
var ennemi string = "E"

// joueurX représente le tableau dans lequel le joueur est
// joueurY représenta la case du tableau dans lequel le joueur est
var joueurX int
var joueurY int

// Définit la valeur initiale pour les cases adjacentes au joueur
var valeurInitCase int = 1

// ================================================================================================================================

var niveau int = 1

var tableauJoueur int = 4
var caseJoueur int = 0

// ================================================================================================================================

func main() {
	GenererCarte()

	Afficher()
}

func GenererCarte() {
	// Rempli les cases des tableau de carte
	for i := 0; i < len(carte); i++ {
		carte[i] = []interface{}{0, 0, 0, 0, 0}
	}

	// Placer le joueur
	carte[tableauJoueur].([]interface{})[caseJoueur] = joueur

	GenererMurs()

	SelectionNiveau()

	PathGeneration()

	fmt.Println("Map de base")
	Afficher()

	Move()
}

// Dans cette fonction, je dois faire bouger les ennemis et le joueur de façon alterné ( print le tableau à chaques changement )
func Move() {
	/* Mettre dans un tableau, les 8 cases qui sont autour d'un ennemi dans ce genre :

		 0 1 2
	   3 E 4 -> [0, 1, 2, 3, 4, 5, 6, 7]
		 5 6 7

	*/
	for t := 0; t < len(carte); t++ {
		for c := 0; c < len(carte[t].([]interface{})); c++ {
			if carte[t].([]interface{})[c] != ennemi {
				continue
			} else {
				// Le tableau dans lequel est l'ennemi
				var cooEnnemiX int = t
				// La case du tableau dans lequel est l'ennemi
				var cooEnnemiY int = c

				var casesAdjacentes []interface{} = make([]interface{}, 8)

				if t == 0 {
					if c == 0 {
						casesAdjacentes[0] = mur
						casesAdjacentes[1] = mur
						casesAdjacentes[2] = mur
						casesAdjacentes[3] = mur
						casesAdjacentes[4] = carte[t].([]interface{})[c+1]
						casesAdjacentes[5] = mur
						casesAdjacentes[6] = carte[t+1].([]interface{})[c]
						casesAdjacentes[7] = carte[t+1].([]interface{})[c+1]
					}
					if c == 4 {
						casesAdjacentes[0] = mur
						casesAdjacentes[1] = mur
						casesAdjacentes[2] = mur
						casesAdjacentes[3] = carte[t].([]interface{})[c-1]
						casesAdjacentes[4] = mur
						casesAdjacentes[5] = carte[t+1].([]interface{})[c-1]
						casesAdjacentes[6] = carte[t+1].([]interface{})[c]
						casesAdjacentes[7] = mur
					}
					if c == 1 || c == 2 || c == 3 {
						casesAdjacentes[0] = mur
						casesAdjacentes[1] = mur
						casesAdjacentes[2] = mur
						casesAdjacentes[3] = carte[t].([]interface{})[c-1]
						casesAdjacentes[4] = carte[t].([]interface{})[c+1]
						casesAdjacentes[5] = carte[t+1].([]interface{})[c-1]
						casesAdjacentes[6] = carte[t+1].([]interface{})[c]
						casesAdjacentes[7] = carte[t+1].([]interface{})[c+1]
					}
				}

				if t == 1 || t == 2 || t == 3 {
					if c == 0 {
						casesAdjacentes[0] = mur
						casesAdjacentes[1] = carte[t-1].([]interface{})[c]
						casesAdjacentes[2] = carte[t-1].([]interface{})[c+1]
						casesAdjacentes[3] = mur
						casesAdjacentes[4] = carte[t].([]interface{})[c+1]
						casesAdjacentes[5] = mur
						casesAdjacentes[6] = carte[t+1].([]interface{})[c]
						casesAdjacentes[7] = carte[t+1].([]interface{})[c+1]
					}
					if c == 4 {
						casesAdjacentes[0] = carte[t-1].([]interface{})[c-1]
						casesAdjacentes[1] = carte[t-1].([]interface{})[c]
						casesAdjacentes[2] = mur
						casesAdjacentes[3] = carte[t].([]interface{})[c-1]
						casesAdjacentes[4] = mur
						casesAdjacentes[5] = carte[t+1].([]interface{})[c-1]
						casesAdjacentes[6] = carte[t+1].([]interface{})[c]
						casesAdjacentes[7] = mur
					}
					if c == 1 || c == 2 || c == 3 {
						casesAdjacentes[0] = carte[t-1].([]interface{})[c-1]
						casesAdjacentes[1] = carte[t-1].([]interface{})[c]
						casesAdjacentes[2] = carte[t-1].([]interface{})[c+1]
						casesAdjacentes[3] = carte[t].([]interface{})[c-1]
						casesAdjacentes[4] = carte[t].([]interface{})[c+1]
						casesAdjacentes[5] = carte[t+1].([]interface{})[c-1]
						casesAdjacentes[6] = carte[t+1].([]interface{})[c]
						casesAdjacentes[7] = carte[t+1].([]interface{})[c+1]
					}
				}

				if t == 4 {
					if c == 0 {
						casesAdjacentes[0] = mur
						casesAdjacentes[1] = carte[t-1].([]interface{})[c]
						casesAdjacentes[2] = carte[t-1].([]interface{})[c+1]
						casesAdjacentes[3] = mur
						casesAdjacentes[4] = carte[t].([]interface{})[c+1]
						casesAdjacentes[5] = mur
						casesAdjacentes[6] = mur
						casesAdjacentes[7] = mur
					}
					if c == 4 {
						casesAdjacentes[0] = carte[t-1].([]interface{})[c-1]
						casesAdjacentes[1] = carte[t-1].([]interface{})[c]
						casesAdjacentes[2] = mur
						casesAdjacentes[3] = carte[t].([]interface{})[c-1]
						casesAdjacentes[4] = mur
						casesAdjacentes[5] = mur
						casesAdjacentes[6] = mur
						casesAdjacentes[7] = mur
					}
					if c == 1 || c == 2 || c == 3 {
						casesAdjacentes[0] = carte[t-1].([]interface{})[c-1]
						casesAdjacentes[1] = carte[t-1].([]interface{})[c]
						casesAdjacentes[2] = carte[t-1].([]interface{})[c+1]
						casesAdjacentes[3] = carte[t].([]interface{})[c-1]
						casesAdjacentes[4] = carte[t].([]interface{})[c+1]
						casesAdjacentes[5] = mur
						casesAdjacentes[6] = mur
						casesAdjacentes[7] = mur
					}
				}

				for i := 0; i < len(casesAdjacentes); i++ {
					if casesAdjacentes[i] == joueur {
						if casesAdjacentes[1] == joueur || casesAdjacentes[3] == joueur || casesAdjacentes[4] == joueur || casesAdjacentes[6] == joueur {
							fmt.Println("L'ennemi est arrivé")
						}

						if casesAdjacentes[0] == joueur || casesAdjacentes[2] == joueur || casesAdjacentes[5] == joueur || casesAdjacentes[7] == joueur {
							fmt.Println("L'ennemi est en case 2")
						}
					}
				}

				// Permet de trier l'interface en skipant tout ce qui n'est pas des int
				var casesAdjacentesTrie []interface{}

				for i := 0; i < len(casesAdjacentes); i++ {
					casesAdjacentesTrie = append(casesAdjacentesTrie, casesAdjacentes[i])
				}

				var intSlice []int

				for _, element := range casesAdjacentesTrie {
					if value, isInt := element.(int); isInt {
						intSlice = append(intSlice, value)
					}
				}

				triCroissant(casesAdjacentesTrie)

				fmt.Println(casesAdjacentes, "(", cooEnnemiX, cooEnnemiY, ")", casesAdjacentesTrie)
			}
		}
	}
}

// Trier une interface par ordre croissant
func triCroissant(interfaceSlice []interface{}) {
	for i := 1; i < len(interfaceSlice); i++ {
		j := i
		for j > 0 && interfaceSlice[j-1].(string) > interfaceSlice[j].(string) {
			interfaceSlice[j], interfaceSlice[j-1] = interfaceSlice[j-1], interfaceSlice[j]
			j--
		}
	}
}

func PathGeneration() {
	// Vérifie si la position du joueur est un mur ou un ennemi
	if carte[tableauJoueur].([]interface{})[caseJoueur] == mur || carte[tableauJoueur].([]interface{})[caseJoueur] == ennemi {
		carte[4].([]interface{})[0] = joueur
	}

	// Détermine la position du joueur dans la carte
	for i := 0; i < len(carte); i++ {
		for j := 0; j < len(carte[i].([]interface{})); j++ {
			if carte[i].([]interface{})[j] == joueur {
				joueurX = j
				joueurY = i

				break
			}
		}
	}

	// Parcours la carte pour définir les valeurs des autres cases
	for distance := 1; distance < len(carte)*len(carte[0].([]interface{})); distance++ {
		for i := 0; i < len(carte); i++ {
			for j := 0; j < len(carte[i].([]interface{})); j++ {
				// Ignore les cases qui sont des murs ou des ennemis
				if carte[i].([]interface{})[j] == mur || carte[i].([]interface{})[j] == ennemi {
					continue
				}

				// Calcule la distance entre la case actuelle et le joueur
				dx := joueurX - j
				dy := joueurY - i
				if dx < 0 {
					dx = -dx
				}
				if dy < 0 {
					dy = -dy
				}
				dist := dx + dy

				// Ignore les cases qui sont plus éloignées que la distance actuelle
				if dist != distance {
					continue
				}

				// Définit la valeur de la case actuelle
				carte[i].([]interface{})[j] = strconv.Itoa(valeurInitCase) // strconv.Itoa() permet de convertir un int en string pour les assigner aux cases de la carte
			}
		}

		// Incrémente la valeur pour les cases adjacentes à la distance actuelle
		valeurInitCase++
	}
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

	if niveau == 2 {
		carte[0].([]interface{})[2] = ennemi
		carte[1].([]interface{})[4] = ennemi
	}

	if niveau == 3 {
		carte[1].([]interface{})[4] = ennemi
	}

	if niveau == 4 {
		carte[0].([]interface{})[1] = ennemi
	}

	if niveau == 5 {
		carte[0].([]interface{})[1] = ennemi
		carte[1].([]interface{})[4] = ennemi
		carte[4].([]interface{})[4] = ennemi
	}

	if niveau == 6 {
		carte[0].([]interface{})[3] = ennemi
		carte[1].([]interface{})[3] = ennemi
		carte[4].([]interface{})[4] = ennemi
	}

	if niveau == 7 {
		carte[1].([]interface{})[2] = ennemi
		carte[3].([]interface{})[4] = ennemi
	}

	if niveau == 8 {
		carte[0].([]interface{})[3] = ennemi
		carte[4].([]interface{})[3] = ennemi
	}

	if niveau == 9 {
		carte[0].([]interface{})[0] = ennemi
		carte[2].([]interface{})[2] = ennemi
		carte[3].([]interface{})[4] = ennemi
		carte[4].([]interface{})[2] = ennemi
	}

	if niveau == 10 {
		carte[0].([]interface{})[0] = ennemi
		carte[1].([]interface{})[2] = ennemi
		carte[2].([]interface{})[3] = ennemi
		carte[3].([]interface{})[4] = ennemi
	}

	if niveau == 11 {
		carte[0].([]interface{})[1] = ennemi
		carte[0].([]interface{})[3] = ennemi
		carte[2].([]interface{})[4] = ennemi
	}

	if niveau == 12 {
		carte[1].([]interface{})[2] = ennemi
		carte[1].([]interface{})[4] = ennemi
		carte[3].([]interface{})[4] = ennemi
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
	fmt.Println("---------------------------------------")
}
