package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if role != userType {
		return errors.New("Unauthorized Access")
	}

	return err
}
func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {

	userType := c.GetString("user_type")
	uid := c.GetString("uid")

	err = nil
	if userType != "USER" && uid != userId {
		return errors.New("UnAuthorized Access ")

	}
	err = CheckUserType(c, userType)

	return nil
}
