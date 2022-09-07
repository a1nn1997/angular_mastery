package utils

import(
	"errors"
	"github.com/gin-gonic/gin"
)

// user correct access ACESS
func CheckUserType(c *gin.Context, role string) (err error){
	userType :=c.GetString("user_type")
	err = nil
	if userType != role {
		err=errors.New("Unautherized to acess this resource")
		return err
	}
	return err
}

func AllowedType(c *gin.Context, role1 string, role2 string) (err error){
	userType :=c.GetString("user_type")
	err = nil
	if userType != role1 {
		if userType != role2{
			err=errors.New("Unautherized to acess this resource")
			return err
		}
	}
	return err
}

 
func MathchUserTypeToUid(c *gin.Context,userId string) (err error){
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err=nil
	if userType =="USER" &&uid != userId{
		err=errors.New("Unautherized to acess this resource")
		return err
	}
	CheckUserType(c, userType)
	return err
}