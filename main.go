package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/username/projectname/database"
	"github.com/username/projectname/helpers"
	"github.com/username/projectname/router"
)

func main() {
	db := database.InitDB("username:password@tcp(localhost:3306)/dbname")
	defer db.Close()

	database.RunMigrations()

	r := router.NewRouter()

	http.Handle("/", helpers.JwtMiddleware(r))

	fmt.Println("Server started on :8000")
	http.ListenAndServe(":8000", nil)
}
