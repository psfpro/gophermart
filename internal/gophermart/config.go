package gophermart

import (
	"flag"
	"os"
)

type Config struct {
	serverAddress  string
	dsn            string
	accrualAddress string
}

func NewConfig() *Config {
	serverAddress := flag.String("a", ":8080", "Server run address")
	dsn := flag.String("d", "postgres://app:pass@localhost:5432/app", "DSN")
	accrualAddress := flag.String("r", "http://localhost:8081", "Accrual system address")
	flag.Parse()

	if envRunAddr := os.Getenv("RUN_ADDRESS"); envRunAddr != "" {
		serverAddress = &envRunAddr
	}
	if envDatabaseDsn := os.Getenv("DATABASE_URI"); envDatabaseDsn != "" {
		dsn = &envDatabaseDsn
	}
	if envAccrualAddr := os.Getenv("ACCRUAL_SYSTEM_ADDRESS"); envAccrualAddr != "" {
		accrualAddress = &envAccrualAddr
	}

	return &Config{
		serverAddress:  *serverAddress,
		dsn:            *dsn,
		accrualAddress: *accrualAddress,
	}
}
