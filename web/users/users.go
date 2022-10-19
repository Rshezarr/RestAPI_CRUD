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

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		log.Printf("Read body - %v", err)
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

func UpdateUserHeader(w http.ResponseWriter, r *http.Request) {
	// id := mux.Vars(r)["id"]
	// if err != nil {
	// 	log.Printf("Get id - %v", err)
	// 	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	// 	return
	// }
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, "http://localhost:8080/user/25", r.Body)
	if err != nil {
		log.Printf("new req - %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("client do - %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write([]byte("pffff"))

	readall, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("new decode - %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var u2 user.User

	if err := json.Unmarshal(readall, u2); err != nil {
		log.Printf("unmarshal - %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := u2.UpdateUserByID(); err != nil {
		log.Printf("update - %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
