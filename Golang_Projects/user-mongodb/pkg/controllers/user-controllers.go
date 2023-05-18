package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/parthin-baraiya/user-mongodb/pkg/models"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	user := models.GetUserByID(userId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(user)
	w.Write(res)

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetUsers()

	res, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)
	user.CreateUser()
	res, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]
	deletedUser := models.DeleteUser(ID)
	res, _ := json.Marshal(deletedUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]
	newUser := &models.User{}
	json.NewDecoder(r.Body).Decode(newUser)

	oldUser := models.GetUserByID(ID)
	if newUser.Name != "" {
		oldUser.Name = newUser.Name
	}
	if newUser.Gender != "" {
		oldUser.Gender = newUser.Gender
	}
	if newUser.Age != 0 {
		oldUser.Age = newUser.Age
	}

	oldUser.UpdateUser(ID)
	res, _ := json.Marshal(oldUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
