[![CI](https://github.com/philskat/self-service-password/actions/workflows/ci.yml/badge.svg)](https://github.com/philskat/self-service-password/actions/workflows/ci.yml)

# Self Service Password

A small web application to allow your users to change there own password of an LDAP-server.

![Example](doc/img/webui.png)

## Known Erros

- The application will not reconnect if the connection to the LDAP-server is lost

## Usage

The simplest way to run the app is to use docker.
You can also run the app on bare metal but for this you need to compile the app fist. See [Build](#build)

### Docker

To run the app with docker you can simply use this command:

```
docker run -d -p 8080:8080 \
    -e SSPASSWORD_LDAP_URL=ldap://localhost \
    -e SSPASSWORD_ADMIN_DN=cn=admin,dc=example,dc=org \
    -e SSPASSWORD_ADMIN_PASSWORD=password \
    -e SSPASSWORD_USER_BASE_DN=ou=users,dc=example,dc=org
    thedarkmen3000/self-service-password:latest
```

You can also put the configuration into a config file and mount the config to the container:

config.toml

```toml
[ldap]
ldap_url = 'ldap://localhost'
admin_dn = 'cn=admin,dc=example,dc=org'
admin_password = 'password'
user_base_dn = 'ou=users,dc=example,dc=org'
```

```
docker run -d -p 8080:8080 \
  -v ./config.toml:/etc/self-service-password/config.toml:ro \
  thedarkmen3000/self-service-password:latest
```

### Bare Metal

After you have [build](#build) the application you can put it in a folder on your system.
You also need to copy the `public` folder of the repositiory in that folder.
Additionally you can place a `config.toml` into the directory to set the configuration.

Layout of the folder

```
baseDir/
  publc/
    index.html
    ...
  config.toml
  self-service-password(.exe)
```

To run the server you just need to run the executable in the background.

## Documentation

The [documentation](docs/README.md) of the configuration is in the `/docs` directory.

## Development

If you want to contribute to the project I suggest you to create an issue first before implementing a new feature to avoid unnessesary work.
You can also freely work on issues posted.

### Requirements

You need to have go installed to compile the project.

You also need to have a running ldap server like [OpenLDAP](https://openldap.org).
To run the app for development you need to create a `config.toml` in the root of the directory to connect to the LDAP-Server.
See the [documentation](docs/config.md) to create the config file.

If you only want to change design the frontend of the application you can simply open the `public/index.html` in the browser.
Because it is only plain HTML, CSS and JS.

### Build

To build the app run the following command in the root directory of the project:

```
go build
```
