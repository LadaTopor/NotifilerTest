package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

type DB struct {
	Name     string `yaml:"POSTGRES_DB"`
	User     string `yaml:"POSTGRES_USER"`
	Password string `yaml:"POSTGRES_PASSWORD"`
	Port     string `yaml:"PORT"`
	Host     string `yaml:"HOST"`
	// Добавьте другие поля вашего конфига здесь
}

// PostgresConnection подключается к базе данных
func PostgresConnection() (*sql.DB, error) {
	config := getDBConfig()

	// строка присоединения к базе
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Name)

	// Подключаемся к базе
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// getDBConfig получает значения из конфига .yaml
func getDBConfig() DB {
	// Прочитать содержимое файла YAML
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Ошибка чтения файла YAML: %v", err)
	}

	// Создать экземпляр структуры для хранения данных из файла YAML
	var config DB

	// Распарсить YAML и сохранить данные в структуру
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Ошибка разбора YAML: %v", err)
	}

	return config
}
