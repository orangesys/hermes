package main

import (
	"fmt"
	"log"
	"net/http"

	// "github.com/prometheus/common/version"

	"github.com/orangesys/janus/routers"
)

func registerServe() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: router,
	}
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
	go func() {
		// service Conections
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}
