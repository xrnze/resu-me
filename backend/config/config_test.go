package config

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoad_Defaults(t *testing.T) {
	os.Setenv("OPENROUTER_API_KEY", "test-key")
	defer os.Unsetenv("OPENROUTER_API_KEY")

	cfg, err := Load()
	assert.NoError(t, err)
	assert.Equal(t, "8080", cfg.Server.Port)
	assert.Equal(t, "https://openrouter.ai/api/v1", cfg.LLM.BaseURL)
	assert.Equal(t, "openai/gpt-oss-120b", cfg.LLM.Model)
	assert.Equal(t, 60*time.Second, cfg.LLM.Timeout)
	assert.Equal(t, "meta-llama/llama-3.1-8b-instruct", cfg.LLM.FilterModel)
	assert.Equal(t, 10*time.Second, cfg.LLM.FilterTimeout)
	assert.Equal(t, 10.0, cfg.RateLimit.Rate)
	assert.Equal(t, 3, cfg.RateLimit.Burst)
	assert.Equal(t, 10*time.Minute, cfg.RateLimit.TTL)
	assert.Equal(t, 30*time.Second, cfg.Server.ReadTimeout)
	assert.Equal(t, 60*time.Second, cfg.Server.WriteTimeout)
	assert.Equal(t, 120*time.Second, cfg.Server.IdleTimeout)
}

func TestLoad_CustomEnv(t *testing.T) {
	os.Setenv("OPENROUTER_API_KEY", "custom-key")
	os.Setenv("PORT", "9090")
	os.Setenv("OPENROUTER_BASE_URL", "https://custom.ai/v1")
	os.Setenv("OPENROUTER_MODEL", "custom/model")
	os.Setenv("RATE_LIMIT_RATE", "20")
	os.Setenv("RATE_LIMIT_BURST", "5")
	os.Setenv("RATE_LIMIT_TTL", "5m")
	os.Setenv("SERVER_READ_TIMEOUT", "15s")
	os.Setenv("SERVER_WRITE_TIMEOUT", "30s")
	os.Setenv("LLM_TIMEOUT", "45s")
	os.Setenv("FILTER_MODEL", "custom/filter-model")
	os.Setenv("FILTER_TIMEOUT", "5s")
	os.Setenv("SERVER_IDLE_TIMEOUT", "60s")
	defer func() {
		os.Unsetenv("OPENROUTER_API_KEY")
		os.Unsetenv("PORT")
		os.Unsetenv("OPENROUTER_BASE_URL")
		os.Unsetenv("OPENROUTER_MODEL")
		os.Unsetenv("LLM_TIMEOUT")
		os.Unsetenv("FILTER_MODEL")
		os.Unsetenv("FILTER_TIMEOUT")
		os.Unsetenv("RATE_LIMIT_RATE")
		os.Unsetenv("RATE_LIMIT_BURST")
		os.Unsetenv("RATE_LIMIT_TTL")
		os.Unsetenv("SERVER_READ_TIMEOUT")
		os.Unsetenv("SERVER_WRITE_TIMEOUT")
		os.Unsetenv("SERVER_IDLE_TIMEOUT")
	}()

	cfg, err := Load()
	assert.NoError(t, err)
	assert.Equal(t, "custom-key", cfg.LLM.APIKey)
	assert.Equal(t, "9090", cfg.Server.Port)
	assert.Equal(t, "https://custom.ai/v1", cfg.LLM.BaseURL)
	assert.Equal(t, "custom/model", cfg.LLM.Model)
	assert.Equal(t, 45*time.Second, cfg.LLM.Timeout)
	assert.Equal(t, "custom/filter-model", cfg.LLM.FilterModel)
	assert.Equal(t, 5*time.Second, cfg.LLM.FilterTimeout)
	assert.Equal(t, 20.0, cfg.RateLimit.Rate)
	assert.Equal(t, 5, cfg.RateLimit.Burst)
	assert.Equal(t, 5*time.Minute, cfg.RateLimit.TTL)
	assert.Equal(t, 15*time.Second, cfg.Server.ReadTimeout)
	assert.Equal(t, 30*time.Second, cfg.Server.WriteTimeout)
	assert.Equal(t, 60*time.Second, cfg.Server.IdleTimeout)
}

func TestLoad_MissingAPIKey(t *testing.T) {
	os.Unsetenv("OPENROUTER_API_KEY")
	_, err := Load()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "OPENROUTER_API_KEY")
}

func TestLoad_InvalidDuration(t *testing.T) {
	os.Setenv("OPENROUTER_API_KEY", "test-key")
	os.Setenv("RATE_LIMIT_TTL", "not-a-duration")
	defer func() {
		os.Unsetenv("OPENROUTER_API_KEY")
		os.Unsetenv("RATE_LIMIT_TTL")
	}()

	_, err := Load()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "RATE_LIMIT_TTL")
}

func TestLoad_InvalidInt(t *testing.T) {
	os.Setenv("OPENROUTER_API_KEY", "test-key")
	os.Setenv("RATE_LIMIT_BURST", "not-an-int")
	defer func() {
		os.Unsetenv("OPENROUTER_API_KEY")
		os.Unsetenv("RATE_LIMIT_BURST")
	}()

	_, err := Load()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "RATE_LIMIT_BURST")
}

func TestLoad_InvalidFloat(t *testing.T) {
	os.Setenv("OPENROUTER_API_KEY", "test-key")
	os.Setenv("RATE_LIMIT_RATE", "not-a-float")
	defer func() {
		os.Unsetenv("OPENROUTER_API_KEY")
		os.Unsetenv("RATE_LIMIT_RATE")
	}()

	_, err := Load()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "RATE_LIMIT_RATE")
}
