package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	_ "modernc.org/sqlite"
)

type KeyValueStore interface {
	Put(key string, value string) (err error)
	Get(key string) (value string, err error)
	Delete(key string) (err error)
}

type InMemoryDatabase struct {
	db map[string]string
}

func NewInMemoryDatabase() *InMemoryDatabase {
	return &InMemoryDatabase{db: make(map[string]string)}
}

func (i *InMemoryDatabase) Put(key string, value string) (err error) {
	i.db[key] = value
	return
}

func (i *InMemoryDatabase) Get(key string) (value string, err error) {
	value, ok := i.db[key]
	if !ok {
		err = errors.New("key not in database")
		return
	}
	return
}

func (i *InMemoryDatabase) Delete(key string) (err error) {
	delete(i.db, key)
	return
}

// A database of key/value pairs, backed by disk.
type KeyValueDatabase struct {
	dbPath string  // path to the backing store
	db     *sql.DB // database connection
}

// Create a new KeyValueDatabase with the given database location.
func NewKeyValueDatabase(path string) (kv *KeyValueDatabase, err error) {
	kv = &KeyValueDatabase{dbPath: path}
	err = kv.connect()
	return
}

// Puts the given key/value into the database. Returns an error if there was an
// error updating the key.
func (db *KeyValueDatabase) Put(key string, value string) (err error) {
	query := fmt.Sprintf("INSERT OR REPLACE INTO kv VALUES (\"%s\", \"%s\");", key, value)
	_, err = db.db.Exec(query)
	return
}

// Retrieve the value associated with the given key.
func (db *KeyValueDatabase) Get(key string) (value string, err error) {
	query := fmt.Sprintf("SELECT value FROM kv WHERE key == \"%s\";", key)
	row := db.db.QueryRow(query)
	err = row.Scan(&value)
	return
}

// Delete the given key from the database, returning an error if the key
// could not be removed.
func (db *KeyValueDatabase) Delete(key string) (err error) {
	query := fmt.Sprintf("DELETE FROM kv WHERE key == \"%s\";", key)
	_, err = db.db.Exec(query)
	return
}

// Connect to the database.
func (db *KeyValueDatabase) connect() (err error) {
	db.db, err = sql.Open("sqlite", db.dbPath)
	return
}

// Module initializer. Will create a new instance of the database if -init is
// supplied with -initPath being the path to the database to initialize.
func init() {
	var initPath string
	var initDb bool
	flag.StringVar(&initPath, "initPath", "./sqlite", "database filename")
	flag.BoolVar(&initDb, "init", false, "whether to initialize the database")
	//flag.Parse()

	if !initDb {
		return
	}

	if initPath == "" {
		log.Fatalf("If -init is supplied, -initPath cannot be empty.")
	}

	fmt.Println(sql.Drivers())
	db, err := NewKeyValueDatabase(initPath)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	createDb(db)
}

// Creates the initial database state.
func createDb(kv *KeyValueDatabase) {
	if kv.db == nil {
		log.Fatalf("Nil db")
	}
	if _, err := kv.db.Exec(
		`
		CREATE TABLE IF NOT EXISTS kv (
			key 	 VARCHAR(20)  NOT NULL,
			value   VARCHAR(20)	NOT NULL,
			PRIMARY KEY (key)
		);
		INSERT OR REPLACE INTO kv
		VALUES
			("matt", "abc123"),
			("elise", "password"),
			("zelda", "chicken<3");
		`); err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
}
