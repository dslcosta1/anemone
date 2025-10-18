package lib

import (
	"fmt"
	"strings"
)

// BuildClassificationPrompt builds a classification prompt for an LLM
// that classifies a given name into one of four categories:
// "Valid", "Invalid", "Offensive", or "Irracional".
// The model must return exactly one of these words, nothing else.
func BuildClassificationPrompt(name, language, countryISO2 string) string {
	examples := getExamplesByCountry(countryISO2)

	promptBuilder := strings.Builder{}
	promptBuilder.WriteString(fmt.Sprintf(
		"You are a name classifier. Your task is to classify a given name according to the following categories:\n\n"+
			"- Valid: a normal, human personal name\n"+
			"- Invalid: a common noun, object, color, or word that is not a name\n"+
			"- Offensive: a name or phrase that contains insults, sexual, or offensive words\n"+
			"- Irracional: a sequence of random or meaningless characters\n\n"+
			"Language: %s\n"+
			"Country: %s\n\n"+
			"Here are examples from this country:\n\n",
		strings.Title(language),
		strings.ToUpper(countryISO2),
	))

	for example, category := range examples {
		promptBuilder.WriteString(fmt.Sprintf("%q → %s\n", example, category))
	}

	promptBuilder.WriteString(fmt.Sprintf(
		"\nNow classify the following name:\n\n%q\n\n"+
			"Return only one word from these options: Valid, Invalid, Offensive, or Irracional.\n"+
			"No explanations, no punctuation — only the category name.", name))

	return promptBuilder.String()
}

// getExamplesByCountry returns example classification maps for each country.
// You can expand this function for other countries later.
func getExamplesByCountry(countryISO2 string) map[string]string {
	switch strings.ToUpper(countryISO2) {
	case "BR":
		return map[string]string{
			"Daniel Silva":  "Valid",
			"Fátima":        "Valid",
			"Giovana":       "Valid",
			"André":         "Valid",
			"Lixão":         "Offensive",
			"Maconheiro":    "Offensive",
			"Andreia Puta":  "Offensive",
			"kfrjsgbjrbhjr": "Irracional",
			"erj4ri4uiujs":  "Irracional",
			"akeg":          "Irracional",
			"Maçã":          "Invalid",
			"Azul":          "Invalid",
		}
	default:
		// Fallback minimal examples in English
		return map[string]string{
			"John":       "Valid",
			"Alice":      "Valid",
			"Apple":      "Invalid",
			"Blue":       "Invalid",
			"Idiot":      "Offensive",
			"asdkjhasd":  "Irracional",
			"qwe123":     "Irracional",
		}
	}
}