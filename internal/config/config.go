package config

import (
	"encoding/json"
	"net"
	"os"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	dbPassEscSeg = "{password}"
	password     = "note-service-password"
)

type Config struct {
	HTTP HTTP `json:"http"`
	GRPC GRPC `json:"grpc"`
	DB   DB   `json:"db"`
}

type DB struct {
	DSN                string `json:"dsn"`
	MaxOpenConnections int32  `json:"max-open-connections"`
}

type GRPC struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type HTTP struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func NewConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return config, err
}

func (c *Config) GetDBConfig() (*pgxpool.Config, error) {
	dbDsn := strings.ReplaceAll(c.DB.DSN, dbPassEscSeg, password)

	poolConfig, err := pgxpool.ParseConfig(dbDsn)
	if err != nil {
		return nil, err
	}

	poolConfig.ConnConfig.BuildStatementCache = nil
	poolConfig.ConnConfig.PreferSimpleProtocol = true
	poolConfig.MaxConns = c.DB.MaxOpenConnections

	return poolConfig, nil
}

func (g *GRPC) GetAddressGRPC() string {
	return net.JoinHostPort(g.Host, g.Port)
}

func (h *HTTP) GetAddressHTTP() string {
	return net.JoinHostPort(h.Host, h.Port)
}
