package web

import (
	"github.com/bmerchant22/spo-task-3.git/pkg/store"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateWebServer(store *store.PostgresStore) *Server {
	srv := new(Server)
	srv.store = store
	srv.r = gin.Default()

	srv.r.POST(kUserSignup, MW1, srv.UserSignup)
	srv.r.GET(kUserLogin, MW1, srv.UserLogin)
	if err := srv.r.Run("localhost:8080"); err != nil {
		zap.S().Errorf("Error while running the server !")
	}

	zap.S().Infof("Web server created successfully !!")

	return srv
}
