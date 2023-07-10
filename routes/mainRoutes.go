package routes

import (
	authRoutes "github.com/GulabSinghSikarwar/preProjBackend/routes/auth"
	userRoutes "github.com/GulabSinghSikarwar/preProjBackend/routes/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine) {

	api := r.Group("/api")
	authRoutes.RegisterAuthRoutes(api.Group("/auth"))
	userRoutes.RegisterUserRoutes(api.Group("/user"))

}
