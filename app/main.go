package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/cors"
	"github.com/yagrush/go-sample-restapi/app/infrastructure"
	"github.com/yagrush/go-sample-restapi/app/infrastructure/config"
	"golang.org/x/sync/errgroup"
	"gopkg.in/yaml.v2"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var c config.Config
	b, err := os.ReadFile("config.yaml")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	yaml.Unmarshal(b, &c)

	engine, err := infrastructure.NewEngine(c)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	hndlr := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodHead},
		AllowedHeaders: []string{"Origin", "Accept", "Content-Type", "X-Requested-With", "x-api-key"},
	}).Handler(engine)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", c.Port),
		Handler: hndlr,
	}

	errgrp, ctx := errgroup.WithContext(ctx)
	errgrp.Go(func() error {
		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			slog.Error("closed server (Fatal)")
			return err
		}
		return nil
	})

	<-ctx.Done()

	if err := server.Shutdown(context.Background()); err != nil {
		slog.Error("shutdown error")
		os.Exit(1)
	}

	err = errgrp.Wait()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
