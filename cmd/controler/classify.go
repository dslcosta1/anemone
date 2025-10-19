package controler

import (
	"fmt"
	"strings"
	"github.com/dslcosta1/anemone/cmd/gateway"
	"github.com/dslcosta1/anemone/cmd/lib"
)

func ClassifyName(nameInput string, languageInput string, countryInput string) string {
	name := strings.ToLower(nameInput)
	language := strings.ToLower(languageInput)
	country := strings.ToLower(countryInput)

	fmt.Println(name)

	prompt := lib.BuildClassificationPrompt(name, language, country)
	response := gateway.GetGenAIOutput(prompt)

	return response
}