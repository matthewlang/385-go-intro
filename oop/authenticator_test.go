package main

import (
	"testing"
)

func TestAddUser(t *testing.T) {
	inmem := NewInMemoryDatabase()
	auth := NewAuthenticator(inmem)
	err := auth.AddUser("user1", "password")
	if err != nil {
		t.Fatalf("Adding a user once is ok: %v", err)
	}
	err = auth.AddUser("user1", "newpassword")
	if err == nil {
		t.Fatalf("Adding a user twice should be an error.")
	}
}
