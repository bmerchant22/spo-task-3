package web

import (
	"github.com/bmerchant22/spo-task-3.git/pkg/store"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	r     *gin.Engine
	store *store.PostgresStore
}

func (srv *Server) UserSignup(c *gin.Context) {
	username := c.Query("username")
	hash := c.Query("hashDigest")

	if err := srv.store.UserSignup(username, hash); err != nil {
		zap.S().Errorf("Error while calling UserSignup method : %v", err)
		panic(err)
	}
	zap.S().Infof("user signed up with username: %v", username)
	c.String(http.StatusOK, "User signed up successfully !")
}

func (srv *Server) UserLogin(c *gin.Context) {
	username := c.Query("username")
	hash := c.Query("hashDigest")

	checkUser := srv.store.UserLogin(username, hash)
	if checkUser == true {
		c.String(http.StatusOK, "User logged in successfully !!")
	} else {
		c.String(http.StatusForbidden, "Wrong password or username !")
	}
}

func MW1(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
}
