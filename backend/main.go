package main

import (
	"database/sql"
	"embed"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

//go:embed sql/checkIfSchemaExists.sql
//go:embed sql/gameCreateSchema.sql
//go:embed sql/createDummyData.sql
var f embed.FS
var db *sql.DB

func main() {
	// Connect to database
	db = connectToDB()
	var schemas = []string{
		"game",
	}
	DropAllSchemas(db, schemas) //Be careful with this lmao, only for testing
	CreateMissingSchemas(db, schemas)
	query := GetSQLFile("createDummyData")
	_, err := db.Exec(query)
	Panic(err)

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
