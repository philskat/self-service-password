package config

import "os"

func LoadEnvConfig() *Config {
  return &Config{
    Ldap: LdapConfig{
      LdapUrl: os.Getenv(LdapUrlEnv),
      AdminDn: os.Getenv(AdminDnEnv),
      AdminPassword: os.Getenv(AdminPasswordEnv),
      UserBaseDn: os.Getenv(UserBaseDnEnv),
    },
  } 
}
