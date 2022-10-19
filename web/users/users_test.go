package users

import (
	"bytes"
	uu "crud/pkg/user"
	"encoding/json"
	"log"
	"net/http"
	"testing"
)

func Test_CreateUser(t *testing.T) {
	//Create
	u := uu.User{
		FirstName: "ezio",
		LastName:  "auditore",
	}

	res, err := json.Marshal(u)
	if err != nil {
		log.Fatalf("marshal - %v\n", err)
	}

	resp, err := http.Post("http://localhost:8080/user/25", "application/json", bytes.NewBuffer(res))
	if err != nil {
		log.Fatalf("post - %v\n", err)
	}

	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("status code - %v\n", err)
	}

	//Read
	var u2 uu.User

	resp2, err := http.Get("http://localhost:8080/user/25")
	if err != nil {
		log.Fatalf("get - %v\n", err)
	}

	if err = json.NewDecoder(resp2.Body).Decode(&u2); err != nil {
		log.Fatalf("unmarshal - %v\n", err)
	}

	if u.FirstName != u2.FirstName || u.LastName != u2.LastName {
		log.Fatalln("no match")
	}

	//Update

}
