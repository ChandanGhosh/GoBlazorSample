package main

import (
	"encoding/json"
	"fmt"
	"github.com/chandanghosh/go-blazor-app/go-server/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func staticFilesHandler(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, filepath.Join("app", "blazor-client", "dist", r.URL.Path[1:]))
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	jsonfile, err := os.Open(filepath.Join("app", "blazor-client", "dist", "sample-data", "weather.json"))
	if err != nil {
		log.Fatalln("Unable to open the file", err.Error())
	}
	defer jsonfile.Close()

	cont, _ := ioutil.ReadAll(jsonfile)

	weathers := models.Weathers{}
	err = json.Unmarshal(cont, &weathers)
	if err != nil {
		fmt.Println("JSON unmarshaling error", err.Error())
	}

	json.NewEncoder(w).Encode(weathers.Weathers)
}

func main() {

	http.HandleFunc("/", staticFilesHandler)
	http.HandleFunc("/api/weather", weatherHandler)

	fmt.Println("Server started at http://localhost:8080/")
	log.Panicln(http.ListenAndServe(":8080", nil))
}
