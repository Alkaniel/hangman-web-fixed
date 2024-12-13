package backend

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Game struct {
	ToFind       string `json:"toFind"`
	Masked       string `json:"masked"`
	AttemptsLeft int    `json:"attemptsLeft"`
	Status       string `json:"status"`
	Pseudo       string `json:"username"`
	Difficulty   int    `json:"difficulty"`
}

func (g *Game) newGame(username string, difficulty int) *Game {
	g.UpdateDifficulty(difficulty)
	g.UpdateToFind(ChooseWord(g.Difficulty))
	g.UpdateMasked(MaskWord(g.ToFind))
	g.UpdateAttemptsLeft(10)
	g.UpdateStatus("En cours")
	g.UpdatePseudo(username)
	return g
}

func (g *Game) LoadGame(username string, difficulty int) {
	username = strings.ToLower(strings.TrimSpace(username)) //Fonction qui permet de charger une partie sauvegardée
	filepath := "saves/" + username + ".json"
	if _, err := os.Stat(filepath); err == nil {
		// File exists, read the entire content
		data, err := os.ReadFile(filepath)
		if err != nil {
			fmt.Println("Failed to read save file. Starting a new game.")
			*g = *g.newGame(username, difficulty)
			return
		}

		// Unmarshal JSON data into the Game struct
		if err := json.Unmarshal(data, g); err != nil {
			fmt.Println("Failed to decode save file. Starting a new game.")
			*g = *g.newGame(username, difficulty)
			return
		}
		
		os.Remove(filepath)
		fmt.Println("Game loaded successfully:", g)
	} else {
		// File does not exist, initialize a new game
		fmt.Println("No save file found. Creating a new game.")
		*g = *g.newGame(username, difficulty)
	}
}

func (g *Game) UpdateDifficulty(difficulty int) { //Fonction qui permet de set la difficulté
	g.Difficulty = difficulty
}

func (g *Game) UpdateToFind(toFind string) { //Fonction qui permet de set le mot à trouver
	g.ToFind = toFind
}

func (g *Game) UpdateMasked(masked string) { //Fonction qui permet de set le mot masqué
	g.Masked = masked
}

func (g *Game) UpdateAttemptsLeft(attemptsLeft int) { //Fonction qui permet de set le nombre d'essais
	g.AttemptsLeft = attemptsLeft
}

func (g *Game) UpdateStatus(status string) { //Fonction qui permet de set le status de la partie
	g.Status = status
}

func (g *Game) UpdatePseudo(pseudo string) { //Fonction qui permet de set le pseudo du joueur
	g.Pseudo = pseudo
}

func (g *Game) SaveGame(pseudo string) { //Fonction qui permet de sauvegarder la partie
	data, err := json.Marshal(g)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("saves/"+strings.ToLower(strings.TrimSpace(pseudo))+".json", data, 0644)
	if err != nil {
		panic(err)
	}
}
