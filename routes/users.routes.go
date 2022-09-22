package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matias0000/go-postgresql/db"
	"github.com/matias0000/go-postgresql/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user []models.User
	db.DB.Find(&user)
	json.NewEncoder(w).Encode(&user)
	// w.Write([]byte("Get Users Handler"))

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
	// fmt.Println(params["id"])
	json.NewEncoder(w).Encode(&user)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)

	w.Write([]byte("PostUsersHandler"))

}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
	}
	// remover de la db
	db.DB.Unscoped().Delete(&user, params["id"])
	w.WriteHeader(http.StatusOK)

}
