package configs

import (
	"fmt"
	project_root_directory "github.com/golang-infrastructure/go-project-root-directory"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var SourceCodeRootDirectory, _ = GetRootDirectory()

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func GetConfigs(cfg *Config) {

	file, err := os.Open(SourceCodeRootDirectory + "/configs/config.yaml")
	if err != nil {
		log.Fatal("Error while reading configs", err.Error())
	}

	err = yaml.NewDecoder(file).Decode(cfg)
	if err != nil {
		log.Fatal("Error while decoding configs", err.Error())
	}
	defer file.Close()
}

func GetRootDirectory() (string, error) {
	directory, err := project_root_directory.GetSourceCodeRootDirectory()
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return directory, nil
}
