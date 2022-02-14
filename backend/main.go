package main

import (
	"database/sql"
	"embed"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/microcosm-cc/bluemonday"
	"log"
)

//go:embed sql/checkIfSchemaExists.sql
//go:embed sql/guestbookCreateSchema.sql
//go:embed sql/otherCreateSchema.sql
var f embed.FS
var db *sql.DB
var namePolicy = bluemonday.StrictPolicy()
var commentPolicy = bluemonday.UGCPolicy()

func main() {
	// Connect to database
	db = connectToDB()
	var schemas = []string{}
	CreateMissingSchemas(db, schemas)
	startHttpServer()
}

// CheckError Review and replace this function wherever possible
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func PrintError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func Panic(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetSQLFile(name string) string {
	data, _ := f.ReadFile(fmt.Sprintf("sql/%s.sql", name))
	return string(data)
}
