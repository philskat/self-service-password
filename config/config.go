package config

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
  LdapUrlEnv       = "SSPASSWORD_LDAP_URL"
  AdminDnEnv       = "SSPASSWORD_ADMIN_DN"
  AdminPasswordEnv = "SSPASSWORD_ADMIN_PASSWORD"
  UserBaseDnEnv    = "SSPASSWORD_USER_BASE_DN"
)

type IConfig interface {
  isComplete() bool
  fillUnset(*IConfig)
  uncompleteError(*strings.Builder)
}


type Config struct {
  Ldap LdapConfig  `toml:"ldap"`
}

func LoadConfig() (*Config, error) {
  envCfg := LoadEnvConfig()

  cfg := envCfg

  if cfg.isComplete() {
    return cfg, nil
  }

  fileCfg := LoadFileConfig()

  cfg.fillUnset(fileCfg)

  if cfg.isComplete() {
    return cfg, nil
  }

  defaultCfg := LoadDefaultConfig()

  cfg.fillUnset(defaultCfg)

  if !cfg.isComplete() {
    var builder strings.Builder

    builder.WriteString("Missing configuration:\n")

    cfg.uncompleteError(&builder)

    return nil, errors.New(builder.String())
  }

  return cfg, nil 
}

func ConfigMiddleware(config *Config) gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Set("config", config)
  }
}

func (c *Config) isComplete() bool {
  if !c.Ldap.isComplete() {
    return false
  }

  return true
}

func (c *Config) fillUnset(otherCfg *Config) {
  c.Ldap.fillUnset(&otherCfg.Ldap)
}

func (c *Config) uncompleteError(builder *strings.Builder) {
  c.Ldap.uncompleteError(builder)
}
