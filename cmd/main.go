package main

import (
	lbConfig "github.com/violetaplum/go-balancer/config"
	"github.com/violetaplum/go-balancer/internal/proxy"
	"github.com/violetaplum/go-balancer/internal/server"
	"log"
	"net/http"
)

func main() {
	cfg, err := lbConfig.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	lb := proxy.NewLoadBalancer(cfg.Nodes)
	handler := server.NewHandler(lb)

	log.Printf("Starting server on.. :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, handler); err != nil {
		log.Fatalf("Server failed.. : %v", err)
	}
}
