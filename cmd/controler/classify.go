package controler

import (
	"fmt"
	"github.com/dslcosta1/anemone/cmd/gateway"
)


var classifications = map[string]string{
	"Daniel":	"Valid",
	"Fátima": 	"Valid",
	"Giovana": 	"Valid",
	"André": 	"Valid",
	"Lixão":   "Offensive",
	"Maconheiro":   "Offensive",
	"kfrjsgbjrbhjr":   "Irracional",
	"erj4ri4uiujs":   "Irracional",
	"Maçã":   "Invalid",
	"Azul":   "Invalid",
}



func ClassifyName(name string, language string, country string) string {
	fmt.Println(name)

	prompt := "Explain how AI works in a few words"
	gateway.GetGenAIOutput(prompt)

	val, ok := classifications[name]
	if !ok {
		return "Inconclusive"	
	}

	return val
}