package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github/robotxt/propertyzag/app/logging"
	"github/robotxt/propertyzag/app/rest/apiv1"

	"github.com/gorilla/mux"
)

type AppHandler struct {
	Router *mux.Router
}

var env string = strings.ToUpper(os.Getenv("ENVIRONMENT"))
var router *mux.Router = mux.NewRouter().StrictSlash(false)

func (a *AppHandler) Initialize() {

	port := os.Getenv("PORT")
	if port == "" {
		// set default port if empty
		port = "8080"
	}

	logging.Info("running on port: ", port)

	// setup apiv1
	apiv1 := apiv1.ApiV1{}
	apiv1.Router = router
	apiv1.Ctx = context.Background()
	apiv1.SetRouters()

	a.Router = router

	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		ReadTimeout:  10 * time.Second, // timeout in data read
		WriteTimeout: 10 * time.Second, // timeout for response
	}

	go func() {
		logging.Info("Starting Server")
		if err := srv.ListenAndServe(); err != nil {
			logging.Fatal(err)
		}
	}()
	waitForShutdown(srv)

}

func checkErr(err error) {
	if err != nil {
		logging.Fatal(err)
	}
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	logging.Info("Shutting down")
	os.Exit(0)
}
