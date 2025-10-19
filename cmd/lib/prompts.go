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
			"- valid: a normal, human personal name\n" +
			"- non_name: a common noun, object, color, or other word that is not a person’s name\n" +
			"- offensive: a word or phrase that contains insults, sexual, or explicit meaning\n" +
			"- irrational: a random or meaningless sequence of characters\n" +
			"- misspelled: a name with spelling or grammatical errors, such as missing accents or repeated letters\n" +
			"- humorous: a name or phrase that forms a joke, pun, or double meaning\n\n"+
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
			"Return only one word from these options: valid, non_name, offensive, irracional, misspelled or humorous\n"+
			"No explanations, no punctuation — only the category name.", name))

	return promptBuilder.String()
}

// NormalizeClassificationOutput takes the raw LLM output and
// returns a clean category string ("valid", "invalid", "offensive", or "irracional").
// If none of the expected categories are found, it returns an empty string.
func NormalizeClassificationOutput(output string) string {
	output = strings.ToLower(strings.TrimSpace(output))

	categories := []string{"valid", "non_name", "offensive", "irracional", "misspelled", "humorous"}

	// If the output is exactly one of the categories, return it immediately
	for _, cat := range categories {
		if output == cat {
			return cat
		}
	}

	// Otherwise, search if one of the categories appears inside the string
	for _, cat := range categories {
		if strings.Contains(output, cat) {
			return cat
		}
	}

	// No recognized classification found
	return "unsure"
}

// getExamplesByCountry returns example classification maps for each country.
// You can expand this function for other countries later.
func getExamplesByCountry(countryISO2 string) map[string]string {
	switch strings.ToUpper(countryISO2) {
	case "BR":
		return map[string]string{
		// ✅ Valid (proper names)
		"Daniel Silva":  "valid",
		"Fátima":        "valid",
		"Giovana":       "valid",
		"André":         "valid",
		"João Pedro":    "valid",
		"Maria Clara":   "valid",

		// 🚫 NonName (replaces "invalid")
		"Maçã":          "nonname",
		"Azul":          "nonname",
		"Carro":         "nonname",
		"Pedra":         "nonname",
		"Rio":           "nonname",
		"Montanha":      "nonname",

		// 💢 Offensive
		"Lixão":         "offensive",
		"Maconheiro":    "offensive",
		"Andreia Puta":  "offensive",
		"Filho da Mãe":  "offensive",
		"Otário":        "offensive",

		// 🔣 Irracional (random/gibberish)
		"kfrjsgbjrbhjr": "irracional",
		"erj4ri4uiujs":  "irracional",
		"akeg":          "irracional",
		"zxqpt":         "irracional",
		"lljdska":       "irracional",

		// ✏️ Misspelled / GrammarError
		"Jooao":         "misspelled",
		"Andree":        "misspelled",
		"maria silvva":  "misspelled",
		"lUcas":         "misspelled",
		"Anadrea":       "misspelled",

		// 😄 Humorous / Double meaning
		"Paula Tejano":  "humorous",
		"Oscar Alho":    "humorous",
		"Tereza Meia":   "humorous",
		"Jacinto Leito": "humorous",
		"Armando Pinto": "humorous",
		}
	default:
		// Fallback minimal examples in English
		return map[string]string{
		// Valid
		"John":       "valid",
		"Alice":      "valid",
		"Michael":    "valid",

		// NonName
		"Apple":      "nonname",
		"Blue":       "nonname",
		"Car":        "nonname",

		// Offensive
		"Idiot":      "offensive",
		"Dumbass":    "offensive",

		// Irrational
		"asdkjhasd":  "irracional",
		"qwe123":     "irracional",

		// Misspelled
		"Jhon":       "misspelled",
		"Alyce":      "misspelled",

		// Humorous
		"Ben Dover":  "humorous",
		"Phil McCracken": "humorous",
		}
	}
}