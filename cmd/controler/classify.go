package controler

import (
	"fmt"
	"strings"
	"github.com/dslcosta1/anemone/cmd/gateway"
	"github.com/dslcosta1/anemone/cmd/lib"
)

// ClassificationFeature represents a single category and its score
type ClassificationFeature struct {
	Category string `json:"category"`
	Score    int    `json:"score"`
}

// ClassificationResult is the structured response
type ClassificationResult struct {
	Name     string                  `json:"name"`
	Country  string                  `json:"country"`
	Language string                  `json:"language"`
	Features []ClassificationFeature `json:"features"`
}

func ClassifyName(nameInput string, languageInput string, countryInput string) ClassificationResult {
	name := strings.ToLower(nameInput)
	language := strings.ToLower(languageInput)
	country := strings.ToLower(countryInput)

	fmt.Println(name)

	prompt := lib.BuildClassificationPrompt(name, language, country)
	rawResponse := gateway.GetGenAIOutput(prompt)
	classification := lib.NormalizeClassificationOutput(rawResponse)

	categories := []string{"valid", "invalid", "offensive", "irracional"}
	features := make([]ClassificationFeature, 0, len(categories))

	// Assign 100 to the predicted category, 0 to others
	for _, cat := range categories {
		score := 0
		if classification == cat {
			score = 100
		}
		features = append(features, ClassificationFeature{
			Category: cat,
			Score:    score,
		})
	}


	result := ClassificationResult{
		Name:     name,
		Country:  country,
		Language: language,
		Features: features,
	}

	return result
}