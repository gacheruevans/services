package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/TwiN/go-color"
	"github.com/egacheru/services/pkg/routers"
	"github.com/egacheru/services/pkg/utils"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	get := utils.GetEnvWithKey
	SITEURL := get("SITE_URL")
	//PORT := get("SERVER_PORT")

	Port := 5000
	portstring := strconv.Itoa(Port)

	r := mux.NewRouter()
	routers.RegisterStoreRoutes(r)
	http.Handle("/", r)

	message := fmt.Sprintf("Listening on >> %s:%s", SITEURL, portstring)
	log.Print(color.Ize(color.Blue, message))
	err := http.ListenAndServe(":"+portstring, r)
	if err != nil {
		log.Fatal(color.Ize(color.Red, "ListenAndServe error: "), err)
	}
}
