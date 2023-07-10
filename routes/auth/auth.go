package authRoutes

import (
	authController "github.com/GulabSinghSikarwar/preProjBackend/controllers/auth"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(authGroup *gin.RouterGroup) {
	authGroup.POST("/signup", authController.Sign_up)
	authGroup.POST("/login", authController.Login)

}
