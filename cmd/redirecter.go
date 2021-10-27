package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"smartShortLink/database"
	"time"
)

var dbInstance = database.Instance

func Redirect(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	keyObj := database.Keys{
		UpdatedAt: time.Now().UTC(),
		UserKey:   vars["id"],
	}
	dbInstance.IncreaseKey(keyObj)

	url, err := dbInstance.GetDomainByID(getUrlIdByTime())
	if err != nil {
		return
	}

	http.Redirect(writer, request, fmt.Sprintf("http://%s", url), 302)
}

func getUrlIdByTime() int {
	now := time.Now().UTC() // current local tim
	key := now.Hour() / (24 / amountOfTimeDivisions)
	if key <= 0 {
		key = 1
	}
	return key
}
