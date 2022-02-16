package main

import (
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	port     = 5432
	user     = "dawid"
	password = "&H2FEZ5+0X!y\"G8?!beWlV:j5"
	dbname   = "dawid"
)

type Session struct {
	SessionId string
	PostTime  time.Time
	Ip        string
}

func connectToDB() *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, port, user, password, dbname)
	for {
		conn, err := sql.Open("postgres", psqlconn)
		if err != nil || conn.Ping() != nil {
			log.Println("Cannot connect to database. Trying again...")
			time.Sleep(time.Second * 2)
		} else {
			return conn
		}
	}
}

// DropAllSchemas This function must succeed, so it can panic all it wants.
func DropAllSchemas(db *sql.DB, schemas []string) {
	for _, x := range schemas {
		// TODO: Make this work without Sprintf
		_, err := db.Exec(fmt.Sprintf("drop schema if exists %s cascade;", x))
		Panic(err)
	}
}

// CreateMissingSchemas This function must succeed, so it can panic all it wants.
func CreateMissingSchemas(db *sql.DB, schemas []string) {
	exists, err := CheckIfSchemasExists(db, schemas)
	Panic(err)
	for i, x := range exists {
		if x {
			fmt.Printf("%s : YES\n", schemas[i])
		} else {
			fmt.Printf("Does %s exist? : NO\n", schemas[i])
			fmt.Printf("Creating %s Schema...\n", schemas[i])
			query := GetSQLFile(fmt.Sprintf("%sCreateSchema", schemas[i]))
			_, err := db.Exec(query)
			Panic(err)
		}
	}
}

func CheckIfSchemasExists(db *sql.DB, schemas []string) ([]bool, error) {
	data := GetSQLFile("checkIfSchemaExists")

	var hasSchema = make([]bool, len(schemas))

	for i, x := range schemas {
		rows, err := db.Query(data, x)
		if rows != nil {
			for rows.Next() {
				var exists bool
				err = rows.Scan(&exists)
				if err != nil {
					return []bool{false}, err
				}
				hasSchema[i] = exists
			}
		}
	}

	return hasSchema, nil
}

func GetActiveSessions(db *sql.DB) []Session {
	query := "SELECT session_id, post_time, ip FROM game.active_sessions"

	rows, err := db.Query(query)
	if err != nil {
		PrintError(err)
	}
	var output []Session
	for rows.Next() {
		var sessionId string
		var postTime time.Time
		var ip string
		err = rows.Scan(&sessionId, &postTime, &ip)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		output = append(output, Session{SessionId: sessionId, PostTime: postTime, Ip: ip})
	}
	return output
}

/*
func QueryDBExampleFunc(db *sql.DB) []Post {
	query := "SELECT name, comment, post_time, visible FROM dawid.guestbook.posts WHERE visible = true;"

	rows, err := db.Query(query)
	if err != nil {
		PrintError(err)
	}
	var output []Post
	for rows.Next() {
		var name string
		var comment string
		var postTime time.Time
		var visible bool
		err = rows.Scan(&name, &comment, &postTime, &visible)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		output = append(output, Post{Name: name, Comment: comment, Time: postTime})
	}
	return output
}
*/
