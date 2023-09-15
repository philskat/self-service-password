# Configuration

The configuration of the application can be done either by a configuration file or via the environment.
This allows an easy configuration when used as a container.

The configuration is loaded in the following order:

1. Environment
2. Different file location (see By File)
3. Default configuration options

Example config:

```toml
[ldap]
ldap_url = 'ldap://localhost'
admin_dn = 'cn=admin,dc=example,dc=org'
admin_password = 'password'
user_base_dn = 'ou=users,dc=example,dc=org'
```

## `ldap`
This section configures the ldap connection

### `ldap_url` (required)

Specify the URL where the LDAP-Server can be reached. Supports `ldap://` and `ldaps://`

### `admin_dn` (required)

The DN of the admin usere or the user used to change the passwords.

### `admin_password` (required)

The password of the user in `admin_dn`

### `user_base_dn` (required)

The base DN used to find the user.

## By Environment

All config options are available through the environment.
The following table shows the mapping of the names to the config file.

| Config           | Environment                 |
| ---------------- | --------------------------- |
| `ldap_url`       | `SSPASSWORD_LDAP_URL`       |
| `admin_dn`       | `SSPASSWORD_ADMIN_DN`       |
| `admin_password` | `SSPASSWORD_ADMIN_PASSWORD` |
| `user_base_dn`   | `SSPASSWORD_USER_BASE_DN`   |
