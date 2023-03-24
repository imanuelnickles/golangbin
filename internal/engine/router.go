package engine

import (
	"golangbin/internal/handler"

	"github.com/gin-gonic/gin"
)

func initDepedencies() {

}

func initRouter(r *gin.Engine) {
	repo := initRepository()
	uc := initUsecase(repo)
	h := initHandler(uc)

	v1 := r.Group("/v1")
	registerUserHandler(v1.Group("/user"), h.User)

}

func registerUserHandler(r *gin.RouterGroup, h handler.User) {
	r.GET("/", h.Get)
}
