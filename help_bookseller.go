package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func StockList(listArticlesStrings []string, listSearchArticleLetter []string) string {
	//word = strings.ToUpper(word)
	var articlesLettersCountsMap = make(map[string]int, len(listArticlesStrings))
	var article strings.Builder
	var isCountFounded = false
	var articleCountString strings.Builder
	for _, articleWithCountString := range listArticlesStrings {
		for _, symbol := range articleWithCountString {
			if isCountFounded == true {
				articleCountString.WriteString(string(symbol))
				continue
			}
			if unicode.IsSpace(symbol) {
				isCountFounded = true
				continue
			}
			article.WriteString(string(symbol))
		}
		if isCountFounded == false {
			panic("cannot find article count: " + articleWithCountString)
		}
		articleCount, err := strconv.Atoi(articleCountString.String())
		if err != nil {
			panic("cannot get int from article count: " + articleCountString.String())
		}
		for _, symbol := range article.String() {
			if _, ok := articlesLettersCountsMap[string(symbol)]; ok {
				articlesLettersCountsMap[string(symbol)] += articleCount
			} else {
				articlesLettersCountsMap[string(symbol)] = articleCount
			}
		}

		isCountFounded = false
		article.Reset()
		articleCountString.Reset()
	}
	fmt.Printf("result: %v\n", articlesLettersCountsMap)

	var articlesLettersCountString strings.Builder
	articleLetterCount := 0
	for _, articleLetter := range listSearchArticleLetter {
		if count, ok := articlesLettersCountsMap[articleLetter]; ok {
			articleLetterCount = count
		}
		articlesLettersCountString.WriteString("(" + articleLetter + " : " + strconv.Itoa(articleLetterCount) + ") -")
		articleLetterCount = 0
	}

	return articlesLettersCountString.String()
}

func doTestStockList(listArticlesStrings []string, listSearchArticleLetter []string, articlesLettersMatchExpected string) {
	if articlesLettersMatchResult := StockList(listArticlesStrings, listSearchArticleLetter); articlesLettersMatchResult != articlesLettersMatchExpected {
		fmt.Printf("TEST NOT PASSED. listArticlesStrings: %v, listSearchArticleLetter: %v, expected: %s, result: %s\n", listArticlesStrings, listSearchArticleLetter, articlesLettersMatchResult, articlesLettersMatchExpected)
	} else {
		fmt.Printf("test passed. listArticlesStrings: %v, listSearchArticleLetter: %v, expected: %s, result: %s\n", listArticlesStrings, listSearchArticleLetter, articlesLettersMatchResult, articlesLettersMatchExpected)
	}
}

func main() {
	doTestStockList(
		[]string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"},
		[]string{"A", "B", "C", "D"},
		"(A : 0) - (B : 1290) - (C : 515) - (D : 600)",
	)
}