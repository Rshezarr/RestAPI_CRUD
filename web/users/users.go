package users

import (
	"crud/pkg/user"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Printf("Get id - %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	u := user.User{
		ID: id,
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("new decode - %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(data, &u); err != nil {
		log.Printf("unmarshal - %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := u.CreateUser(); err != nil {
		log.Printf("Create User - %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Printf("Get id - %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	u := user.User{
		ID: id,
	}

	usr, err := u.GetUserByID()
	if err != nil {
		log.Printf("Get User - %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&usr); err != nil {
		log.Printf("Encode - %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Printf("Get id - %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	u2 := user.User{
		ID:        id,
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
	}

	if err := u2.UpdateUserByID(); err != nil {
		log.Printf("update - %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Printf("Get id - %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	u := user.User{
		ID: id,
	}

	if err := u.DeleteUserByID(); err != nil {
		log.Printf("delete - %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
