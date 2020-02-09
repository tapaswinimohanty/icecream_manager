package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/lyquocnam/zalora_icecream/lib"
	"github.com/lyquocnam/zalora_icecream/services"
	"github.com/lyquocnam/zalora_icecream/view_models"
	"time"
)

func LoginHandler(c echo.Context) error {
	cc := c.(*lib.CustomContext)
	model := view_models.LoginModel{}
	err := cc.Bind(&model)
	if err != nil {
		return cc.BadRequest("Login information is not valid")
	}

	if err := model.Valid(); err != nil {
		return cc.BadRequest(err.Error())
	}

	user := services.FindUserByUsername(model.Username)
	if user == nil || user.ID == 0 {
		return cc.NotFound("user is not exist")
	}

	if !user.ComparePassword(model.Password) {
		return cc.BadRequest("password is not valid")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // 3 days

	t, err := token.SignedString([]byte(lib.Config.Secret))
	if err != nil {
		return err
	}

	return cc.OK(echo.Map{
		"token": t,
	})

}
