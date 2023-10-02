package config

import "fmt"

type DatabaseConfig interface {
	Uri() string
}

func (pc PostgresConfig) Uri() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		pc.User, pc.Password, pc.Host, pc.Port, pc.Dbname)
}

type PostgresConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Dbname   string
}

type ServerConfig struct {
	Address string
}

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}
