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

type ResponseClassify struct {
	Result string `json:"result" validate:"required,max=100"`
}

func (app *application) classifyHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateClassifyPayload
	if err := readJSON(w, r, &payload); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Call LLM model to get the classification
	result := controler.ClassifyName(payload.Name, payload.Language, payload.Country)
	var response = ResponseClassify{result}


	if err := writeJSON(w, http.StatusCreated, response); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
