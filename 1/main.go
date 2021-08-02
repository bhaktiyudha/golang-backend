package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int        `json:"id"`
	Username string     `json:"username"`
	Parent   NullString `json:"parent"`
}

type NullString struct {
	sql.NullString
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

func main() {
	fmt.Println("Running SQL")
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/stockbit_mission")
	defer db.Close()

	checkErr(err)

	var version string

	errVer := db.QueryRow("SELECT VERSION()").Scan(&version)

	checkErr(errVer)

	fmt.Println(version)

	// Execute the query
	results, err := db.Query(`
			SELECT 
				user.id as id, 
				user.username as username, 
				parent.username as parent 
			FROM user_ex user
			LEFT JOIN user_ex parent ON user.parent_id = parent.id`)
	checkErr(err)

	for results.Next() {
		var user User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.ID, &user.Username, &user.Parent)
		checkErr(err)

		log.Printf("user instance := %#v\n", user)
		userJSON, err := json.Marshal(&user)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("json := %s\n\n", userJSON)
		}
	}

	err = results.Err()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
