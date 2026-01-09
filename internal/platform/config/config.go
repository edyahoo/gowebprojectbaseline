package config

import (
	"bufio"
	"flag"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	ServerAddr             string
	DBHost                 string
	DBPort                 int
	DBUser                 string
	DBPassword             string
	DBName                 string
	LogLevel               string
	Env                    string
	SessionDurationMinutes int
}

func (c *Config) IsProduction() bool {
	return c.Env == "production"
}

/*
	func init() {
		// Bootstrap logger: always DEBUG, simple, early
		h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		})

		slog.SetDefault(slog.New(h))

		slog.Debug("Bootstrap logger initialized (DEBUG)")
	}
*/
func applyCLIOverrides() {
	logLevel := flag.String("log-level", "", "log level (superdebug, debug, info, warn, error)")
	flag.Parse()

	if *logLevel != "" {
		_ = os.Setenv("LOG_LEVEL", *logLevel)
	}
}

// LoadDotEnv reads a .env file and sets environment variables for keys that
// are not already set in the process environment.
// If the file does not exist, it returns nil (no error).
func LoadDotEnv(path string) error {
	slog.Debug("Looking for config file:", "path", path)
	f, err := os.Open(path)
	if err != nil {
		slog.Warn("Config File Not Found:", "path", path)
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// skip blanks and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// split KEY=VALUE
		i := strings.IndexByte(line, '=')
		if i <= 0 {
			continue
		}

		key := strings.TrimSpace(line[:i])
		key = strings.TrimPrefix(key, "\ufeff")
		val := strings.TrimSpace(line[i+1:])

		// strip surrounding quotes
		val = strings.Trim(val, `"'`)
		slog.Info("Key, Value Pairs", "key", key, "val", val)

		// don't overwrite existing env vars
		if _, exists := os.LookupEnv(key); exists {
			slog.Info("Key exists already in OS.  Not updating", "key", key)
			continue
		}

		_ = os.Setenv(key, val)
	}

	return scanner.Err()
}

func Load() *Config {
	applyCLIOverrides() // ← one line, no pollution
	cwd, _ := os.Getwd()
	slog.Info("Current working directory", "cwd", cwd)
	_ = LoadDotEnv(".env") // ignore missing file; returns nil if not present
	return &Config{
		ServerAddr:             getEnv("SERVER_ADDR", "8080"),
		DBHost:                 getEnv("DB_HOST", "localhost"),
		DBPort:                 getEnvAsInt("DB_PORT", 5432),
		DBUser:                 getEnv("DB_USER", "postgres"),
		DBPassword:             getEnv("DB_PASSWORD", ""),
		DBName:                 getEnv("DB_NAME", "tailwindtest"),
		LogLevel:               getEnv("LOG_LEVEL", "debug"),
		Env:                    getEnv("ENV", "development"),
		SessionDurationMinutes: getEnvAsInt("SESSION_DURATION_MINUTES", 60),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	slog.Warn("Invalid string for environment variable, using fallback",
		"key", key, "fallback", fallback)
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
		slog.Warn("Invalid integer for environment variable, using fallback",
			"key", key,
			"value", value,
			"fallback", fallback,
		)
	}
	return fallback
}
