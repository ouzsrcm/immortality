package restapi

import (
	"fmt"
	"immortality/pkg/restapi/middleware"
	"immortality/pkg/restapi/routes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	PORT    = "8080"
	HOST    = "localhost"
	VERSION = "v1"
)

func Initialize() {
	fmt.Println("Initializing restapi...")

	r := mux.NewRouter()

	amw := middleware.AuthMiddleWare{}
	amw.Init()
	r.Use(amw.Middleware)

	routes.Routes(r, VERSION)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%s", HOST, PORT),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
