package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	model "crud/pkg/user"
)

var id string = "1"

// Create
func Test_CreateUser(t *testing.T) {
	u := model.User{
		Data: model.Data{
			FirstName: "test_first_name_01",
			LastName:  "test_first_name_01",
			Interests: "test_interest1_01,test_interest2_01",
		},
	}

	postRes, err := json.Marshal(u)
	if err != nil {
		log.Fatalf("create: marshal - %v\n", err)
	}

	resp, err := http.Post("http://localhost:8080/user/"+id, "application/json", bytes.NewBuffer(postRes))
	if err != nil {
		log.Fatalf("create: repsonse - %v\n", err)
	}

	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("create: create status code - %v\n", err)
	}
}

// Read
func Test_ReadUser(t *testing.T) {
	u := model.User{
		Data: model.Data{
			FirstName: "test_first_name_01",
			LastName:  "test_first_name_01",
			Interests: "test_interest1_01,test_interest2_01",
		},
	}
	var u2 model.User

	resp2, err := http.Get("http://localhost:8080/user/" + id)
	if err != nil {
		log.Fatalf("read: get - %v\n", err)
	}

	data, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		log.Fatalf("read: read ioutil - %v\n", err)
	}

	fmt.Printf("DATA: %T\n", data)

	if err := json.Unmarshal(data, &u2); err != nil {
		log.Fatalf("read: unmarshal - %v\n", err)
	}

	if u.Data.FirstName != u2.Data.FirstName || u.Data.LastName != u2.Data.LastName || u.Data.Interests != u2.Data.Interests {
		log.Fatalln("read: fields doesn't match")
	}
}

// Update
func Test_UpdateUser(t *testing.T) {
	u3 := model.User{
		Data: model.Data{
			FirstName: "test_first_name_02",
			LastName:  "test_first_name_02",
			Interests: "test_interest1_02,test_interest2_02",
		},
	}

	updRes, err := json.Marshal(u3)
	if err != nil {
		log.Fatalf("update: marshal - %v\n", err)
	}

	testUpdURL := fmt.Sprintf("%s?first_name=%s&last_name=%s&interests=%s", id, u3.FirstName, u3.LastName, u3.Interests)

	updReq, err := http.NewRequest(http.MethodPut, "http://localhost:8080/user/"+testUpdURL, bytes.NewBuffer(updRes))
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
}

// Delete
func Test_DeleteUser(t *testing.T) {
	delReq, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/user/"+id, nil)
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
