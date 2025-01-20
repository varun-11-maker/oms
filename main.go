// for starting with psql open pgAdmin4 enter password 1234 for the connection
// in command prompt change user to postgres by command psql -U postgres and enter 1234 as the password

// Commands to enter

// CREATE DATABASE <dbname>
// \c <dbname> - switch to that database
// then create table

package main

import (
	"log"
	"net/http"
	"usingPostgres/router"
)

func main() {
	r := router.Router()
	log.Fatal(http.ListenAndServe(":8080", r))

}
