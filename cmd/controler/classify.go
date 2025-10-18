package controler

import (
	"fmt"
	"strings"
	"github.com/dslcosta1/anemone/cmd/gateway"
	"github.com/dslcosta1/anemone/cmd/lib"
)


var classifications = map[string]string{
    "Daniel Silva":   "valid",
    "fátima":         "valid",
    "giovana":        "valid",
    "andré":          "valid",
    "lixão":          "offensive",
    "maconheiro":     "offensive",
    "andreia Puta":   "offensive",
    "kfrjsgbjrbhjr":  "irracional",
    "erj4ri4uiujs":   "irracional",
    "akeg":           "irracional",
    "maçã":           "invalid",
    "azul":           "invalid",
	"pão de queijo":  "invalid",
}



func ClassifyName(nameInput string, languageInput string, countryInput string) string {
	name := strings.ToLower(nameInput)
	language := strings.ToLower(languageInput)
	country := strings.ToLower(countryInput)

	fmt.Println(name)

	prompt := lib.BuildClassificationPrompt(name, language, country)
	gateway.GetGenAIOutput(prompt)

	val, ok := classifications[name]
	if !ok {
		return "Inconclusive"	
	}

	return val
}