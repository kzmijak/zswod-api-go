package database

type DatabaseConfig struct {
	DriverName string `json:"driverName"`
	DSN        string `json:"dsn"`
}