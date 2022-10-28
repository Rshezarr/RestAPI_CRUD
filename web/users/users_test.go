package users

import (
	"bytes"
	model "crud/pkg/user"
	"encoding/json"
	"log"
	"net/http"
	"testing"
)

func Test_CreateUser(t *testing.T) {
	//Create
	u := model.User{
		FirstName: "ezio",
		LastName:  "auditore",
	}

	postRes, err := json.Marshal(u)
	if err != nil {
		log.Fatalf("create: marshal - %v\n", err)
	}

	resp, err := http.Post("http://localhost:8080/user/1", "application/json", bytes.NewBuffer(postRes))
	if err != nil {
		log.Fatalf("create: post - %v\n", err)
	}

	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("create: create status code - %v\n", err)
	}

	//Read
	var u2 model.User

	resp2, err := http.Get("http://localhost:8080/user/1")
	if err != nil {
		log.Fatalf("read: get - %v\n", err)
	}

	if err = json.NewDecoder(resp2.Body).Decode(&u2); err != nil {
		log.Fatalf("read: unmarshal - %v\n", err)
	}

	if u.FirstName != u2.FirstName || u.LastName != u2.LastName {
		log.Fatalln("read: fields doesn't match")
	}

	//Update
	u3 := model.User{
		FirstName: "ezio2",
		LastName:  "auditore2",
	}

	updRes, err := json.Marshal(u3)
	if err != nil {
		log.Fatalf("update: marshal - %v\n", err)
	}

	updReq, err := http.NewRequest(http.MethodPut, "http://localhost:8080/user/1", bytes.NewBuffer(updRes))
	if err != nil {
		log.Fatalf("update: new request - %v\n", err)
	}

	updResp, err := http.DefaultClient.Do(updReq)
	if err != nil {
		log.Fatalf("update: client do - %v\n", err)
	}

	if updResp.StatusCode != http.StatusNoContent {
		log.Fatalln("update: status code doesn't match")
	}

	//Delete
	delReq, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/user/1", nil)
	if err != nil {
		log.Fatalf("delete: new request - %v\n", err)
	}

	delResp, err := http.DefaultClient.Do(delReq)
	if err != nil {
		log.Fatalf("delete: client do - %v\n", err)
	}

	if delResp.StatusCode != http.StatusAccepted {
		log.Fatalln("delete: status code doesn't match")
	}
}
