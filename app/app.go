package main

import (
	"log"
	"net/http"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/julienschmidt/httprouter"
	"github.com/codehell/rolGame/models"
)

var db *mgo.Database

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	db = session.DB("rolGame")
	router := httprouter.New()
	router.GET("/characters", GetAllCharacters)
	router.POST("/characters", StoreCharacter)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetAllCharacters(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	collection := db.C("characters")
	w.Header().Set("Content-Type", "application/json")
	var characters []models.Character
	if err := collection.Find(bson.M{}).All(&characters); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": {"code": 500, "message":"Ey! Boy"}}`))
		return
	}

	response, _ := json.Marshal(characters)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func StoreCharacter(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// TODO: The character name must be unique
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	character := models.Character{}
	if err := json.NewDecoder(r.Body).Decode(&character); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": {"code": 400, "message":"Ey! Boy"}}`))
		return
	}
	collection := db.C("characters")
	if err := collection.Insert(&character); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": {"code": 500, "message":"Ey! Boy"}}`))
		return
	}
	response, _ := json.Marshal(character)

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
