package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"selltech/internal"
	"selltech/internal/handlers/get_names"
	"selltech/internal/handlers/state"
	"selltech/internal/handlers/update"
	"selltech/internal/repository"
)

func main() {
	appConfig, err := internal.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Fatal error loading config: %s", err)
	}

	database, err := sql.Open("postgres", appConfig.DBURL)
	if err != nil {
		log.Fatalf("Fatal error connecting to database: %s", err)
	}
	defer database.Close()
	repository := repository.New(database)
	updateHandler := update.New(appConfig, repository)
	getStateHandler := state.New(appConfig, repository)
	getNamesHandler := get_names.New(appConfig, repository)

	http.HandleFunc("/update", updateHandler.UpdateHandler)
	http.HandleFunc("/state", getStateHandler.GetState)
	http.HandleFunc("/get_names", getNamesHandler.GetNames)
	port := 8080
	fmt.Printf("Server is running on :%d\n", port)
	//Обепрнуть в ошибку

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		return
	}
}
