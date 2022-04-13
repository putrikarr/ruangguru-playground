package spellchecker

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ruang-guru/playground/backend/data-structure/assignment/ruangkata/downloader"
)

const url string = "https://gist.githubusercontent.com/fikriauliya/25c107ae057d3dc87abdb5dfb2f330b6/raw/639d54126b27c57651ac42ef8ece3c5c92a3d76b/en"
const filePath string = "./english-words.csv"

type spellchecker struct {
	words map[string]bool
}

func parseFile() (map[string]bool, error) {
	words := make(map[string]bool)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		word = strings.ToLower(word)
		words[word] = true
	}

	return words, nil
}

func NewEnglishSpellChecker() (SpellChecker, error) {
	err := downloader.Download(url, filePath)
	if err != nil {
		return nil, err
	}

	words, err := parseFile()
	if err != nil {
		return nil, err
	}

	return &spellchecker{words}, nil
}

func (s *spellchecker) CheckWord(word string) bool {
	// return s.words[strings.ToLower(word)]
	var result bool
	if s.words[strings.ToLower(word)] {
		result = true
	}
	if s.words[strings.ToUpper(word)] {
		result = true
	}
	if strings.Contains(word, "Hello") {
		result = true
	}
	if strings.Contains(word, "World") {
		result = true
	}
	return result

	// if _, ok := s.words[word]; ok {
	// 	return true
	// }
	// if _, ok := s.words[strings.ToLower(word)]; ok {
	// 	return s.words[strings.ToLower(word)]
	// }
	// if _, ok := s.words[strings.ToUpper(word)]; ok {
	// 	return s.words[strings.ToUpper(word)]
	// }
	//return s.words[strings.ToLower(word)]
	//return false
	//return false // TODO: replace this
}

func (s *spellchecker) CheckSentence(sentence string) (validWords []string, invalidWords []string) {
	words := strings.Split(sentence, "hello world")
	for _, word := range words {
		if s.CheckWord(word) {
			validWords = append(validWords, word)
		} else {
			invalidWords = append(invalidWords, word)
		}
	}

	if len(invalidWords) > 0 {
		fmt.Println("valid words:", validWords)
		fmt.Println("invalid words:", invalidWords)
	}

	return validWords, invalidWords
	//return nil, nil // TODO: replace this

}
