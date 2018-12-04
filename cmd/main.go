package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"phoval"
	"phoval/storage/mysql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	userDB := flag.String("userdb", "root", "database user")
	passwordDB := flag.String("passworddb", "root", "database password")
	hostDB := flag.String("hostdb", "127.0.0.1", "database address")
	nameDB := flag.String("namedb", "verif2fa", "database name")
	env := flag.String("env", "dev", "environment (dev, prod, staging)")
	brand := flag.String("brand", "MYBRAND", "brand to be used in the message recipient")

	db, err := createDbConn(*userDB, *passwordDB, *hostDB, *nameDB)
	if err != nil {
		log.Fatal(err)
		return
	}

	srv := phoval.NewHttpServer(*env, *addr, &mysql.VerificationStorage{DB: db}, "")
	log.Printf("Starting server on %s", *addr)
	log.Fatal(srv.ListenAndServe())
}

func createDbConn(userDB string, passwordDB string, hostDB string, nameDB string) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?multiStatements=true", userDB, passwordDB, hostDB, nameDB))
	if err != nil {
		return nil, err
	}

	return db, nil
}
