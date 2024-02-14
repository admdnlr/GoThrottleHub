package config

import (
	"log"
	"os"
)

// Config struct'ı, yapılandırma parametrelerini saklar.
type Config struct {
	APIKey             string
	DatabaseConnection string
	// Diğer yapılandırma parametrelerini de buraya ekleyebilirsiniz.
}

// NewConfig, yeni bir yapılandırma örneği oluşturur.
func NewConfig() *Config {
	return &Config{
		APIKey:             getEnv("API_KEY", ""),
		DatabaseConnection: getEnv("DATABASE_URL", "user:password@/dbname"),
		// Diğer parametreler için de çevre değişkenlerini okuyabilirsiniz.
	}
}

// getEnv, bir çevre değişkenini okur ve eğer tanımlı değilse varsayılan değeri döndürür.
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValue
	}
	return value
}

// MustGetEnv, bir çevre değişkenini okur ve eğer tanımlı değilse loglar ve programı durdurur.
func MustGetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Çevre değişkeni %s tanımlı olmalıdır", key)
	}
	return value
}
