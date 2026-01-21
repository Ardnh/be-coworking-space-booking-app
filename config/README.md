# CONFIG Directory

## 📖 Filosofi
**"Configuration should be external, not hardcoded"**. Filosofi dari folder config adalah memisahkan semua konfigurasi aplikasi ke satu tempat yang mudah dikelola. Ini mengikuti prinsip 12-Factor App dimana config harus disimpan di environment variables.

## 🎯 Definisi
Directory yang berisi file-file untuk mengelola konfigurasi aplikasi seperti database connection, API keys, server port, dan pengaturan lainnya. Biasanya menggunakan environment variables atau file konfigurasi (.env, .yaml, .json).

## 💡 Tanggung Jawab
- Load environment variables
- Validasi konfigurasi
- Setup database connection
- Manage external service credentials
- Environment-specific settings (dev, staging, prod)

## 📝 Contoh Struktur
```
config/
├── config.go         # Load & validate config
├── database.go       # Database connection setup
└── README.md         # Dokumentasi ini
```

## 🔧 Contoh Kode

### config.go
```go
package config

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

type Config struct {
    // Server
    Port        string
    Environment string

    // Database
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string

    // JWT
    JWTSecret string

    // External APIs
    PaymentAPIKey string
}

func LoadConfig() *Config {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Println("⚠️  No .env file found, using system environment variables")
    }

    config := &Config{
        Port:          getEnv("PORT", "3000"),
        Environment:   getEnv("ENVIRONMENT", "development"),
        DBHost:        getEnv("DB_HOST", "localhost"),
        DBPort:        getEnv("DB_PORT", "5432"),
        DBUser:        getEnv("DB_USER", "postgres"),
        DBPassword:    getEnv("DB_PASSWORD", ""),
        DBName:        getEnv("DB_NAME", "myapp"),
        JWTSecret:     getEnv("JWT_SECRET", "your-secret-key"),
        PaymentAPIKey: getEnv("PAYMENT_API_KEY", ""),
    }

    // Validate critical configs
    if config.DBPassword == "" {
        log.Fatal("❌ DB_PASSWORD is required")
    }

    if config.JWTSecret == "your-secret-key" {
        log.Println("⚠️  WARNING: Using default JWT secret. Change this in production!")
    }

    return config
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
```

### database.go
```go
package config

import (
    "fmt"
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

func ConnectDB(cfg *Config) *gorm.DB {
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
    )

    // Configure GORM logger
    logLevel := logger.Silent
    if cfg.Environment == "development" {
        logLevel = logger.Info
    }

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logLevel),
    })

    if err != nil {
        log.Fatal("❌ Failed to connect to database:", err)
    }

    log.Println("✅ Database connected successfully")
    return db
}
```

## ⚠️ Best Practices
- ✅ Gunakan environment variables untuk sensitive data
- ✅ Provide default values untuk non-critical configs
- ✅ Validasi konfigurasi saat startup
- ✅ Jangan commit file .env ke git
- ✅ Buat .env.example sebagai template
- ❌ Jangan hardcode credentials di code

## 🎓 Tips
- Gunakan library seperti `godotenv` untuk load .env files
- Untuk production, gunakan secret management (AWS Secrets Manager, Vault)
- Buat config berbeda per environment (dev, staging, prod)
