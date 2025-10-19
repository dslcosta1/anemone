package main

import (
	"net/http"
	
	"github.com/dslcosta1/anemone/cmd/controler"
)

type CreateClassifyPayload struct {
	Name   string  	  `json:"name" validate:"required,max=100"`
	Language string   `json:"language" validate:"required,max=1000"`
	Country  string	  `json:"country" validate:"required,max=100"`
}

// ResponseClassify now wraps the full structured result
type ResponseClassify struct {
	Result controler.ClassificationResult `json:"result"`
}

func (app *application) classifyHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateClassifyPayload
	if err := readJSON(w, r, &payload); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Call LLM model to get the structured classification
	result := controler.ClassifyName(payload.Name, payload.Language, payload.Country)

	response := ResponseClassify{
		Result: result,
	}

	if err := writeJSON(w, http.StatusOK, response); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}