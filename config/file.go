package config

import (
	"os"

	toml "github.com/pelletier/go-toml/v2"
)

func LoadFileConfig() *Config {
  var cfg Config

  file, err := readFiles()
  if err != nil {
    return new(Config)
  }

  err = toml.Unmarshal([]byte(file), &cfg) 
  if err != nil {
    return new(Config)
  }

  return &cfg
}

func readFiles() (string, error) {
  var result string

  file, err := os.Open("./config.toml")
  if err != nil {
    file, err = os.Open("/etc/self-service-password/config.toml")
  }

  if err != nil {
    return "", err
  }

  stats, err := file.Stat()
  if err != nil {
    return "", err
  }

  data := make([]byte, stats.Size())

  _, err = file.Read(data)
  if err != nil {
    return "", err
  }

  result = string(data)

  return result, nil
}
