package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
	"time"
)

var port = flag.String("port", "8080", "Define what TCP port to bind to")
var root = flag.String("root", ".", "Define the root filesystem path")

func main() {
	flag.Parse()
	fmt.Println("Server started on:")
	fmt.Println("\thttp://localhost:" + *port)
	srv := &http.Server{
		Addr:         ":" + *port,
		Handler:      handlers.CombinedLoggingHandler(os.Stdout, http.FileServer(http.Dir(*root))),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	srv.ListenAndServe()
}
