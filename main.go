package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/soerjadi/exam/middleware"
	"github.com/soerjadi/exam/utils"
)

const (
	isDebugDefault                  = true
	isDebugUsage                    = "Enable debug mode"
	enablePrintRecoveryStackDefault = true
	enablePrintRecoveryUsage        = "Used to print all stack traces when panic happened"
	addressDefault                  = "127.0.0.1:8080"
	addressUsage                    = "Setup your running ip & port"
	timeoutDefault                  = 15
	timeoutUsage                    = "Set your write and timeout limit"
	envFileDefault                  = ""
	envFileUsage                    = "Set your dot env filpath"
)

var (
	logger              *utils.Logger
	isDebug             bool
	enablePrintRecovery bool
	address             string
	timeout             int
	envFile             string
	wait                time.Duration
)

func init() {
	flag.StringVar(&address, "address", addressDefault, addressUsage)
	flag.StringVar(&envFile, "env", envFileDefault, envFileUsage)
	flag.BoolVar(&isDebug, "debug", isDebugDefault, isDebugUsage)
	flag.BoolVar(&enablePrintRecovery, "enablePrintStack", enablePrintRecoveryStackDefault, enablePrintRecoveryUsage)
	flag.IntVar(&timeout, "timeout", timeoutDefault, timeoutUsage)
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	logger = utils.LogBuilder(isDebug)
	logger.Debug(fmt.Sprintf("Enable Print Recovery? %t", enablePrintRecovery))
	logger.Debug(fmt.Sprintf("Timeout %v", timeout))
	logger.Info(fmt.Sprintf("Running in %v", address))
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	midl := middleware.InitMiddleware()
	r.Use(midl.LoggingMiddleware)
	r.Use(handlers.RecoveryHandler(handlers.PrintRecoveryStack(enablePrintRecovery)))

	routers := RegisterRouter(r)

	allowedOrigin := handlers.AllowedOrigins([]string{"*"})
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Access-Token", "Content-Type"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	server := &http.Server{
		Handler:      handlers.CORS(allowedOrigin, allowedHeaders, allowedMethods)(routers),
		Addr:         address,
		WriteTimeout: time.Duration(15) * time.Second,
		ReadTimeout:  time.Duration(15) * time.Second,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	server.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
