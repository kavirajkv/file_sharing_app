package auth_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"fileshare/middleware/auth"
)

func TestSignup(t *testing.T) {
	reqbody:=auth.User{
		Username: "testuser",
		Email: "test@gmail.com",
		Password: "testpassword",
	}

	reqbodyjson, err:=json.Marshal(reqbody)

	if err!=nil {
		t.Fatal(err)
	}

	resp, err:=http.Post("http://127.0.0.1:8080/signup", "application/json", bytes.NewBuffer(reqbodyjson))

	if err!=nil {
		t.Fatal(err)
	}


	defer resp.Body.Close()

	if resp.StatusCode!=http.StatusOK {
		t.Fatalf("expected status OK; got %v", resp.Status)
	}

	var res auth.Response
	err=json.NewDecoder(resp.Body).Decode(&res)

	if err!=nil {
		t.Fatal(err)
	}

	if res.Message!="User created successfully" {
		t.Fatalf("expected message User created successfully; got %v", res.Message)
	}

}