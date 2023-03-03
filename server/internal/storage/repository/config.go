package repository

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"gopkg.in/yaml.v3"
	"net/url"
	"os"
	"strconv"
)

// настройки соединения
type Config struct {
	Storage struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DbName   string `yaml:"dbname"`
		Timeout  string `yaml:"timeout"`
	} `yaml:"storage"`
}

// функция для создания строки подключения
func NewPoolConfig() (*pgxpool.Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open("./server/internal/storage/config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decoder
	decoder := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	timeout, _ := strconv.Atoi(config.Storage.Timeout)

	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=%s",
		"postgres",
		url.QueryEscape(config.Storage.Username),
		url.QueryEscape(config.Storage.Password),
		os.Getenv("DB_HOST"),
		config.Storage.Port,
		config.Storage.DbName,
		timeout)

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}
	return poolConfig, nil
}
