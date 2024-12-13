package backend

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
)

// ChooseWord est une fonction qui permet de choisir un mot aléatoire dans un fichier texte.
// Elle prend en paramètre la difficulté du mot à choisir.
// Elle retourne le mot choisi.
// Auteur : Enzo Giardinelli
func ChooseWord(difficulty int) string {
	path := "back/ressources/words.txt"

	switch difficulty {
	case 1:
		path = "back/ressources/words.txt"
	case 2:
		path = "back/ressources/words2.txt"
	case 3:
		path = "back/ressources/words3.txt"
	}

	toReturn := ""

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner1 := bufio.NewScanner(file)
	lineRead := 1
	lineCount := 0

	for scanner1.Scan() {
		lineCount++
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	scanner2 := bufio.NewScanner(file)
	lineToRead := rand.Intn(lineCount) + 1

	lineRead = 1

	for scanner2.Scan() {
		if lineRead == lineToRead {
			toReturn = scanner2.Text()
			break
		}
		lineRead++
	}

	return toReturn
}

func MaskWord (toFind string) string {
	var masked string      // Celle qui vas être return
	var notmask []string   //Liste de lettre à retirer
	n := len(toFind)/2 - 1 // Le nombre de lettre qui doit être reveal dans le mot
	if n == 0 {
		n = 1
	} //Si le mot est trop court, on met 1 lettre
	for i := 0; i < n; i++ { //La boucle s'arrête au nombre de
		j := rand.Intn(len(toFind) - 1) //Choisi un nombre aléatoire sur les mots
		notmask = append(notmask, string(toFind[j]))
	}
	for _, letter := range toFind { //Liste pour masquer le mot
		if HangmanContains(notmask, string(letter)) { //Check si la rune du mot est la lune de la lettre.
			masked += strings.ToUpper(string(letter))
		} else {
			masked += "_"
		}
	}
	return masked
}

func HangmanContains(letterNotMasked []string, letter string) bool {
	for _, l := range letterNotMasked {
		if l == letter {
			return true
		}
	}
	return false
}

func CheckEntry(input string, toFind string, masked string) string {
	for i, char := range toFind {
		if string(char) == input {
			maskedRunes := []rune(masked)
			maskedRunes[i] = []rune(strings.ToUpper(input))[0]
			masked = string(maskedRunes)
		}
	}
	return masked
}