package main

import (
	"flag"
	//_ "github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/docs"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"os"
	"time"
	"timtube/config"
	"timtube/controller"
)

var sessionManager *scs.SessionManager

func Environment() *config.Env {
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.IdleTimeout = 20 * time.Minute
	env := &config.Env{
		ErrorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime),
		InfoLog:  log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile),
		Path:     "./view/html/",
		Session:  sessionManager,
	}
	return env
}

// @title Timtube web
// @version 1.0
// @description Timtube client interface.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email espoir@ditkay.com

// @license.name MIT
// @license.url http://timtube.org/LICENCE

// @BasePath /api/v1
func main() {
	addr := flag.String("addr", ":2221", "HTTP network address")
	flag.Parse()
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: Environment().ErrorLog,
		Handler:  controller.Controllers(Environment()),
	}

	Environment().InfoLog.Printf("Starting server on %s", *addr)
	// Call the ListenAndServe() method on our new http.Server struct.
	error := srv.ListenAndServe()
	Environment().ErrorLog.Fatal(error)

}
