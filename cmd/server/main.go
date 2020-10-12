package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/iam-solutions/api-gateway/internal/routes"
)

var (
	name, addr string
)

func configureEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err.Error())
	}

	name = os.Getenv("SERVICE_NAME")
	if name == "" {
		name = "API GATEWAY"
	}

	addr = os.Getenv("SERVICE_PORT")
	if addr == "" {
		addr = ":8081"
	} else {
		addr = ":" + addr
	}
}

func waitForShutdown(srv http.Server, l *log.Logger) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Graceful shutdown:", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}

func main() {
	configureEnv()

	l := log.New(os.Stdout, strings.ToUpper(name)+" ", log.LstdFlags)

	r := mux.NewRouter()

	routes.MapURLPathsToHandlers(r, l)

	srv := http.Server{
		Addr:         addr,
		Handler:      r,
		ErrorLog:     l,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			l.Fatalln(err)
		}
	}()
	l.Printf("%s is running on %s\n", name, addr)

	waitForShutdown(srv, l)
}
