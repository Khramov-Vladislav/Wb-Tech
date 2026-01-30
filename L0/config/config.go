package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	// PostgreSQL
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	SSLMode    string

	// Kafka
	KafkaBrokers []string
	KafkaTopic   string
	KafkaGroupID string

	// HTTP
	HTTPPort string

	// Cache
	CacheTTL time.Duration

	// Graceful shutdown
	ShutdownTimeout time.Duration

	// Static / API
	StaticDir      string
	OrderURLPrefix string
}

// LoadConfig загружает конфигурацию из env
func LoadConfig() *Config {
	cacheTTLHours := getEnvAsInt("CACHE_TTL_HOURS", 1)
	shutdownTimeoutSec := getEnvAsInt("SHUTDOWN_TIMEOUT_SEC", 5)

	return &Config{
		// PostgreSQL
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "microservice_admin"),
		DBPassword: getEnv("DB_PASSWORD", "1234"),
		DBName:     getEnv("DB_NAME", "microservice"),
		SSLMode:    getEnv("DB_SSLMODE", "disable"),

		// Kafka
		KafkaBrokers: getEnvAsSlice("KAFKA_BROKERS", []string{"localhost:9092"}, ","),
		KafkaTopic:   getEnv("KAFKA_TOPIC", "orders"),
		KafkaGroupID: getEnv("KAFKA_GROUP_ID", "order-consumer-group"),

		// HTTP
		HTTPPort: getEnv("HTTP_PORT", "8081"),

		// Cache
		CacheTTL: time.Duration(cacheTTLHours) * time.Hour,

		// Graceful shutdown
		ShutdownTimeout: time.Duration(shutdownTimeoutSec) * time.Second,

		// Static / API
		StaticDir:      getEnv("STATIC_DIR", "./web"),
		OrderURLPrefix: getEnv("ORDER_URL_PREFIX", "/order/"),
	}
}

// GetDSN возвращает строку подключения к БД
func (c *Config) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost,
		c.DBPort,
		c.DBUser,
		c.DBPassword,
		c.DBName,
		c.SSLMode,
	)
}

// helpers
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if valueStr, exists := os.LookupEnv(key); exists {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
	}
	return defaultValue
}

func getEnvAsSlice(key string, defaultValue []string, sep string) []string {
	val := getEnv(key, "")
	if val == "" {
		return defaultValue
	}
	return splitAndTrim(val, sep)
}

func splitAndTrim(s, sep string) []string {
	parts := strings.Split(s, sep)
	result := make([]string, 0, len(parts))
	for _, p := range parts {
		if trimmed := strings.TrimSpace(p); trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}
