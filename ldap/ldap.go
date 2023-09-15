package ldap

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	ldap "github.com/go-ldap/ldap/v3"
)

// Middleware to access Connection
func Ldap(conn *Connection) gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Set("conn", conn)

    c.Next()
  }
}

// Wrapper of LDAP connection
type Connection struct {
  conn ldap.Conn
  url string
  adminDn string
  adminPassword string
  userBaseDn string
  isConnected bool
}


func NewConnection(url string, userBasedn string, admin string, password string) (*Connection, error) {
  result := &Connection{
    url: url,
    adminDn: admin,
    adminPassword: password,
    userBaseDn: userBasedn,
    isConnected: true,
  } 

  err := result.Connect()
  if err != nil {
    return nil, err
  }

  err = result.Login()
  if err != nil {
    return nil, err
  }

  return result, nil
}

func (c *Connection) Connect() error {
  con, err := ldap.DialURL(c.url)
  if err != nil {
    return err
  }

  c.conn = *con

  c.isConnected = true
  
  return nil
}

func (c *Connection) Disconnect() {
  c.conn.Unbind()

  c.conn.Close()
}

func (c *Connection) Login() error {
  res := c.conn.Bind(c.adminDn, c.adminPassword)

  return res
}

func (c *Connection) ChangePassword(dn string, oldPassword string, newPassword string) error {
  request := ldap.NewPasswordModifyRequest(
    dn,
    oldPassword,
    newPassword,
  )

  _, err := c.conn.PasswordModify(request)

  return err
}

func (c *Connection) SearchUser(username string) (string, error) {
  searchRequest := ldap.NewSearchRequest(
    c.userBaseDn,
    ldap.ScopeWholeSubtree,
    ldap.NeverDerefAliases,
    1,
    2,
    false,
    fmt.Sprintf("(uid=%s)", username),
    []string{"dn",},
    nil,
  )
  
  res, err := c.conn.Search(searchRequest)
  if err != nil {
    return "", err
  }

  if len(res.Entries) == 0 {
    return "", errors.New("No user found")
  }

  return res.Entries[0].DN, nil
}
