package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"resu-me/config"
	"resu-me/handler"
	"resu-me/middleware"
	"resu-me/provider"
	"resu-me/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	llmClient := provider.NewOpenAIClient(cfg.LLM.APIKey, cfg.LLM.BaseURL, cfg.LLM.Model)
	analyzeSvc := service.NewAnalyzeService(llmClient, cfg.LLM.Timeout)

	filterClient := provider.NewOpenAIClient(cfg.LLM.APIKey, cfg.LLM.BaseURL, cfg.LLM.FilterModel)
	filterSvc := service.NewFilterService(filterClient, cfg.LLM.FilterTimeout)

	analyzeHandler := handler.NewAnalyzeHandler(analyzeSvc, filterSvc)

	rl := service.NewRateLimiter(cfg.RateLimit.Rate, cfg.RateLimit.Burst, cfg.RateLimit.TTL)
	defer rl.Stop()

	var trustedNetworks []*net.IPNet
	for _, cidr := range cfg.Server.TrustedProxies {
		if _, p, err := net.ParseCIDR(cidr); err == nil {
			trustedNetworks = append(trustedNetworks, p)
		}
	}

	r := mux.NewRouter()
	r.Use(middleware.BodyLimit(100 * 1024))
	r.Use(middleware.CORS(cfg.Server.CORSOrigin))
	r.Use(middleware.RateLimit(rl, trustedNetworks))

	r.Handle("/api/analyze", analyzeHandler)
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}).Methods("GET")

	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      r,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	go func() {
		log.Printf("Server starting on port %s", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}
	log.Println("Server stopped")
}
