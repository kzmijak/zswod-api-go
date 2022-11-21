package config

type Database struct {
	DriverName string `json:"driverName"`
	DSN        string `json:"dsn"`
}

type Config struct {
	Database Database `json:"database"`
}