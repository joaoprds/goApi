package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repository"
	"api/src/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	if err = user.Prepare(); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	repository := repository.NewRepositoryUsers(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusCreated, user)

	defer db.Close()
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Users"))
}

func GetOneUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get One  User"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting User"))
}
