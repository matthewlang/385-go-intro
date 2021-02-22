package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	DbPath string
)

type ComplexNumber struct {
	r float64
	i float64
}

func foo() {
	c := ComplexNumber{r: 3.6, i: 1.2}
	c.r = 3.6
	c.i = 1.2
}

// Type to authenticate users against a database of users.
type Authenticator struct {
	db KeyValueStore // Backing store
}

// Create new authenticator.
func NewAuthenticator(db KeyValueStore) Authenticator {
	return Authenticator{db}
}

// Add the given user. Returns non-nil if the user exists or there was an error.
func (a Authenticator) AddUser(user, pass string) (err error) {
	err = a.db.Put(user, pass)
	return
}

// Update the given user's password. Returns non-nil if the user did not exist.
func (a Authenticator) UpdateUser(user, pass string) (err error) {
	err = a.db.Put(user, pass)
	return
}

// Remove the given user. Returns non-nil if the user does not exist or there
// was an error removing the user.
func (a Authenticator) RemoveUser(user string) (err error) {
	err = a.db.Delete(user)
	return
}

// Validate a user/password pair. success will be true if the supplied password
// matches that of the user and false if not. Returns an error if there was
// an error retrieving information for the given user.
func (a Authenticator) Validate(user, pass string) (success bool, err error) {
	p, err := a.db.Get(user)
	if err != nil {
		return
	}
	success = (p == pass)
	return
}

func main() {
	flag.StringVar(&DbPath, "db", "sqlite", "path to database file")
	flag.Parse()

	db, err := NewKeyValueDatabase(DbPath)
	if err != nil {
		log.Fatalf("Error opening database %v", err)
	}

	auth := NewAuthenticator(db)

	in := bufio.NewReader(os.Stdin)

	fmt.Println("a) add user, u) update user, d) delete user, l) log in, e) exit")
	for {
		fmt.Print("-> ")
		text, _ := in.ReadString('\n')
		text = strings.TrimSpace(text)
		vals := strings.Split(text, " ")

		switch vals[0] {

		case "a":
			u, p := vals[1], vals[2]
			err := auth.AddUser(u, p)
			fmt.Printf("AddUser(%s, %s), Error: %v\n", u, p, err)

		case "u":
			u, p := vals[1], vals[2]
			err := auth.UpdateUser(u, p)
			fmt.Printf("UpdateUser(%s, %s), Error: %v\n", u, p, err)

		case "d":
			u := vals[1]
			err := auth.RemoveUser(vals[1])
			fmt.Printf("RemoveUser(%s), Error: %v\n", u, err)

		case "l":
			u, p := vals[1], vals[2]
			ok, err := auth.Validate(u, p)
			fmt.Printf("Validate(%s, %s): Ok: %v, Error: %v\n", u, p, ok, err)

		case "e":
			os.Exit(0)

		default:
			fmt.Println("a) add user, d) delete user, l) log in, e) exit")
		}
	}
}
