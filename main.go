package main

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	cfg "github.com/node-a-team/irisnet-validator_exporter/config"
	"github.com/node-a-team/irisnet-validator_exporter/exporter"
	rpc "github.com/node-a-team/irisnet-validator_exporter/getData/rpc"
)

func main() {

	port := "26661"

	log, _ := zap.NewDevelopment()
	defer log.Sync()

	cfg.ConfigPath = os.Args[1]

	cfg.Init()
	rpc.OpenSocket(log)

	http.Handle("/metrics", promhttp.Handler())
	go exporter.Start(log)

	err := http.ListenAndServe(":"+port, nil)

	// log
	if err != nil {
		// handle error
		log.Fatal("HTTP Handle", zap.Bool("Success", false), zap.String("err", fmt.Sprint(err)))
	} else {
		log.Info("HTTP Handle", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Listen&Serve", "Prometheus Handler(Port: "+port+")"))
	}

}
