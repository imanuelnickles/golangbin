package engine

import (
	"github.com/gin-gonic/gin"
)

func InitApp() {
	initEngine()
}

func initEngine() {
	r := gin.Default()
	initRouter(r)
	r.Run() /// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
