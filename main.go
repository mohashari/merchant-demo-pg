package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/joho/godotenv"
	"github.com/mohashari/merchant-demo/merchant"

	"github.com/go-kit/kit/log/level"
)

func main() {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "merchant",
			"time", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller)
	}
	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	loadErr := godotenv.Load()
	if loadErr != nil {
		level.Error(logger).Log("msg", "failed to load env")
	}

	ports := os.Getenv("PORT")
	var httpAddr = flag.String("http", ":"+ports, "http listen address")
	flag.Parse()

	ctx := context.Background()
	srv := merchant.NewService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	endpoints := merchant.MakeEndpoint(srv)

	go func() {
		fmt.Println("listen on port", *httpAddr)
		handler := merchant.NewHttpServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("Exit", <-errChan)

}
