package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server    ServerConfig
	LLM       LLMConfig
	RateLimit RateLimitConfig
}

type ServerConfig struct {
	Port           string
	CORSOrigin     string
	TrustedProxies []string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	IdleTimeout    time.Duration
}

type LLMConfig struct {
	APIKey        string
	BaseURL       string
	Model         string
	Timeout       time.Duration
	FilterModel   string
	FilterTimeout time.Duration
}

type RateLimitConfig struct {
	Rate  float64
	Burst int
	TTL   time.Duration
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("no .env file loaded (expected in Docker): %v", err)
	}

	apiKey := os.Getenv("LLM_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("LLM_API_KEY environment variable is required")
	}

	rateLimitRate, err := parseFloatEnv("RATE_LIMIT_RATE", 10)
	if err != nil {
		return nil, fmt.Errorf("invalid RATE_LIMIT_RATE: %w", err)
	}

	rateLimitBurst, err := parseIntEnv("RATE_LIMIT_BURST", 3)
	if err != nil {
		return nil, fmt.Errorf("invalid RATE_LIMIT_BURST: %w", err)
	}

	rateLimitTTL, err := parseDurationEnv("RATE_LIMIT_TTL", 10*time.Minute)
	if err != nil {
		return nil, fmt.Errorf("invalid RATE_LIMIT_TTL: %w", err)
	}

	llmTimeout, err := parseDurationEnv("LLM_TIMEOUT", 60*time.Second)
	if err != nil {
		return nil, fmt.Errorf("invalid LLM_TIMEOUT: %w", err)
	}

	filterTimeout, err := parseDurationEnv("FILTER_TIMEOUT", 10*time.Second)
	if err != nil {
		return nil, fmt.Errorf("invalid FILTER_TIMEOUT: %w", err)
	}

	corsOrigin := getEnv("CORS_ORIGIN", "*")

	trustedProxiesStr := os.Getenv("TRUSTED_PROXIES")
	var trustedProxies []string
	if trustedProxiesStr != "" {
		for _, p := range strings.Split(trustedProxiesStr, ",") {
			trustedProxies = append(trustedProxies, strings.TrimSpace(p))
		}
	}

	readTimeout, err := parseDurationEnv("SERVER_READ_TIMEOUT", 30*time.Second)
	if err != nil {
		return nil, fmt.Errorf("invalid SERVER_READ_TIMEOUT: %w", err)
	}

	writeTimeout, err := parseDurationEnv("SERVER_WRITE_TIMEOUT", 60*time.Second)
	if err != nil {
		return nil, fmt.Errorf("invalid SERVER_WRITE_TIMEOUT: %w", err)
	}

	idleTimeout, err := parseDurationEnv("SERVER_IDLE_TIMEOUT", 120*time.Second)
	if err != nil {
		return nil, fmt.Errorf("invalid SERVER_IDLE_TIMEOUT: %w", err)
	}

	return &Config{
		Server: ServerConfig{
			Port:           getEnv("PORT", "8080"),
			CORSOrigin:     corsOrigin,
			TrustedProxies: trustedProxies,
			ReadTimeout:    readTimeout,
			WriteTimeout:   writeTimeout,
			IdleTimeout:    idleTimeout,
		},
		LLM: LLMConfig{
			APIKey:        apiKey,
			BaseURL:       getEnv("LLM_BASE_URL", "https://openrouter.ai/api/v1"),
			Model:         getEnv("LLM_MODEL", "openai/gpt-oss-120b"),
			Timeout:       llmTimeout,
			FilterModel:   getEnv("FILTER_MODEL", "meta-llama/llama-3.1-8b-instruct"),
			FilterTimeout: filterTimeout,
		},
		RateLimit: RateLimitConfig{
			Rate:  rateLimitRate,
			Burst: rateLimitBurst,
			TTL:   rateLimitTTL,
		},
	}, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func parseIntEnv(key string, fallback int) (int, error) {
	v := os.Getenv(key)
	if v == "" {
		return fallback, nil
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func parseFloatEnv(key string, fallback float64) (float64, error) {
	v := os.Getenv(key)
	if v == "" {
		return fallback, nil
	}
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

func parseDurationEnv(key string, fallback time.Duration) (time.Duration, error) {
	v := os.Getenv(key)
	if v == "" {
		return fallback, nil
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return 0, err
	}
	return d, nil
}
