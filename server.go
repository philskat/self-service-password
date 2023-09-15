package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/philskat/self-service-password/config"
	"github.com/philskat/self-service-password/ldap"
)

type ChangePasswordRequest struct {
	User        string `form:"user"`
	Password    string `form:"password"`
	NewPassword string `form:"newPassword"`
}

func main() {
	r := gin.Default()

	// Frontend
	r.Static("/", "./public")

	conf, err := config.LoadConfig()
	if err != nil {
		return
	}

	conn, err := ldap.NewConnection(conf.Ldap.LdapUrl, conf.Ldap.UserBaseDn, conf.Ldap.AdminDn, conf.Ldap.AdminPassword)
	if err != nil {
		os.Exit(-1)
	}

	defer conn.Disconnect()

	api := r.Group("/api")
	api.Use(ldap.Ldap(conn))
	api.Use(config.ConfigMiddleware(conf))
	{
		api.POST("/changePassword", func(c *gin.Context) {
			var data ChangePasswordRequest
			conn := c.MustGet("conn").(*ldap.Connection)

			err := c.ShouldBindJSON(&data)
			if err != nil {
				c.JSON(200, gin.H{
					"error":   true,
					"message": "You need to provide user, password and newPassword",
				})
				return
			}

			dn, err := conn.SearchUser(data.User)
			if err != nil {
				c.JSON(200, gin.H{
					"error":   true,
					"message": "User not found",
				})

				return
			}

			err = conn.ChangePassword(dn, data.Password, data.NewPassword)
			if err != nil {
				c.JSON(200, gin.H{
					"error":   true,
					"message": "Old password is not correct",
				})
				return
			}

			c.JSON(200, gin.H{
				"error":   false,
				"message": "Password changed",
			})
		})
	}

	// Handle graceful Shutdown
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		if err := server.Close(); err != nil {
			log.Fatal("Server Close:", err)
		}
	}()

	server.ListenAndServe()
}
