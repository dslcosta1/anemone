package main

import "net/http"


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

	payload.Name = "test"
	// Call LLM model to get the classification
	var response = ResponseClassify{"valid"}


	if err := writeJSON(w, http.StatusCreated, response); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
