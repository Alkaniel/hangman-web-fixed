package main

import (
	"fmt"
	backend "hangmanwebfixed/back"
	"html/template"
	"net/http"
	"os"
	"strings"
)

var game backend.Game

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/niveau", NiveauHandler)
	http.HandleFunc("/regle", RegleHandler)
	http.HandleFunc("/hangman", HangmanHandler)
	http.HandleFunc("/save", GameSaveHandler)
	http.HandleFunc("/win", WinHandler)
	http.HandleFunc("/loose", LooseHandler)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Roger boss, server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/index.html"))
	tmpl.Execute(w, nil)
}

func HangmanHandler(w http.ResponseWriter, r *http.Request) {
	filepath := "saves/" + game.Pseudo + ".json"
	if r.Method == "POST" {
		letter := strings.ToLower(strings.TrimSpace(r.FormValue("lettre")))
		if len(letter) > 1 {
			if letter == game.ToFind {
				os.Remove(filepath)
				http.Redirect(w, r, "/win", http.StatusSeeOther)
				return
			} else {
				game.AttemptsLeft -= 2
				fmt.Println("DEBUG : 2 vies perdues, il reste:", game.AttemptsLeft)
			}
		} else {
			if strings.Contains(game.ToFind, letter) {
				game.UpdateMasked(backend.CheckEntry(letter, game.ToFind, game.Masked))
				fmt.Println("DEBUG : Lettre trouvée, le mot est:", game.Masked)

				if !strings.Contains(game.Masked, "_") {
					os.Remove(filepath)
					http.Redirect(w, r, "/win", http.StatusSeeOther)
					return
				}
			} else {
				game.AttemptsLeft--
				fmt.Println("DEBUG : Vie perdue, il reste:", game.AttemptsLeft)
			}
		}

		if game.AttemptsLeft == 0 {
			os.Remove(filepath)
			http.Redirect(w, r, "/loose", http.StatusSeeOther)
			return
		}
	}

	tmpl := template.Must(template.ParseFiles("html/game.html"))
	tmpl.Execute(w, game)
}

func NiveauHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("pseudo")
		difficulty := r.FormValue("level")

		diffLevel := map[string]int{"facile": 1, "moyen": 2, "difficile": 3}[difficulty]

		game.LoadGame(username, diffLevel)
		fmt.Println("DEBUG : Le mot à trouver est: %s", game.ToFind)

		http.Redirect(w, r, "/hangman", http.StatusSeeOther)

	}
	tmpl := template.Must(template.ParseFiles("html/niveaux.html"))
	tmpl.Execute(w, nil)
}

func RegleHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/reglement.html"))
	tmpl.Execute(w, nil)
}

func GameSaveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		game.SaveGame(game.Pseudo)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func WinHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/victory.html"))
	tmpl.Execute(w, game)
}

func LooseHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/defeat.html"))
	tmpl.Execute(w, nil)
}
