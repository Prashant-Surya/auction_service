package main

import (
	"flag"
	"fmt"
	"github.com/Prashant-Surya/auction-service/auction/pkg/implementation"
	"github.com/Prashant-Surya/auction-service/auction/pkg/service"
	"github.com/Prashant-Surya/auction-service/auction/pkg/transport"
	httptransport "github.com/Prashant-Surya/auction-service/auction/pkg/transport/http"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8081", "HTTP listen address")
	)
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "auction",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	// Create Order Service
	var svc service.AuctionService
	{

		svc = implementation.NewService(logger)
	}

	var h http.Handler
	{
		endpoints := transport.MakeEndpoints(svc)
		h = httptransport.NewService(endpoints, logger)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		level.Info(logger).Log("transport", "HTTP", "addr", *httpAddr)
		server := &http.Server{
			Addr:         *httpAddr,
			Handler:      h,
			ReadTimeout:  time.Millisecond * 200,
			WriteTimeout: time.Millisecond * 200,
		}
		errs <- server.ListenAndServe()
	}()

	level.Error(logger).Log("exit", <-errs)
}
