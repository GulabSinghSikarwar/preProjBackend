package userRoutes

import (
	userController "github.com/GulabSinghSikarwar/preProjBackend/controllers/user"
	middleware "github.com/GulabSinghSikarwar/preProjBackend/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(usergroup *gin.RouterGroup) {
	usergroup.Use(middleware.Authenticate)
	usergroup.GET("/:id", userController.GetUserByID)
	

}
