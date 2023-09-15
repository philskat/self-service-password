package config

import (
	"fmt"
	"strings"
)

type LdapConfig struct {
	LdapUrl       string `toml:"ldap_url"`
	AdminDn       string `toml:"admin_dn"`
	AdminPassword string `toml:"admin_password"`
	UserBaseDn    string `toml:"user_base_dn"`
}

func (c *LdapConfig) isComplete() bool {
	if c.LdapUrl == "" {
		return false
	}

	if c.AdminDn == "" {
		return false
	}

	if c.AdminPassword == "" {
		return false
	}

	if c.UserBaseDn == "" {
		return false
	}

	return true
}

func (c *LdapConfig) fillUnset(otherCfg *LdapConfig) {
	if c.LdapUrl == "" && otherCfg.LdapUrl != "" {
		c.LdapUrl = otherCfg.LdapUrl
	}

	if c.AdminDn == "" && otherCfg.AdminDn != "" {
		c.AdminDn = otherCfg.AdminDn
	}

	if c.AdminPassword == "" && otherCfg.AdminPassword != "" {
		c.AdminPassword = otherCfg.AdminPassword
	}

	if c.UserBaseDn == "" && otherCfg.UserBaseDn != "" {
		c.UserBaseDn = otherCfg.UserBaseDn
	}
}

func (c *LdapConfig) uncompleteError(builder *strings.Builder) {
  if c.LdapUrl == "" {
    builder.WriteString(fmt.Sprintf("- LdapUrl (%s)\n", LdapUrlEnv))
  }

  if c.AdminDn == "" {
    builder.WriteString(fmt.Sprintf("- AdminDn (%s)\n", AdminDnEnv))
  }

  if c.AdminPassword == "" {
    builder.WriteString(fmt.Sprintf("- Password (%s)\n", AdminPasswordEnv))
  }

  if c.UserBaseDn == "" {
    builder.WriteString(fmt.Sprintf("- UserBaseDn (%s)\n", UserBaseDnEnv))
  }
}
