package api

import (
	"net/http"

	"github.com/codeSum27/iam/pkg/global/db"
	"github.com/codeSum27/iam/pkg/global/util"
	"github.com/codeSum27/iam/pkg/model"
	"github.com/labstack/echo/v4"
)

// Login User
// (POST /auth/login)
	func (i *IamServer) PostAuthLogin(ctx echo.Context) error {

		return nil
	}
	// Logout User
	// (POST /auth/logout)
	func (i *IamServer) PostAuthLogout(ctx echo.Context) error {

		return nil
	}
	// SignUp User
	// (POST /auth/signup)
	func (i *IamServer) PostAuthSignup(ctx echo.Context) error {		
		user := model.User{}
		createUser(ctx, &user)

		// JWT Token logic
		

		return ctx.JSON(http.StatusCreated, user)
	}

	func createUser(ctx echo.Context, user *model.User) error {
		db := db.GetDBClient()

		if err := ctx.Bind(user); err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}
		
		if err := db.Find(&user, "username = ?", user.Username); err.Error != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{
				"message": "username is already exist!",
			})
		}
		if err := db.Find(&user, "email = ?", user.Email); err.Error != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{
				"message": "email is already exist!",
			})
		}

		hashPass, err := util.PasswordEncode(user.Password)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		user.Password = hashPass

		if err := db.Create(&user); err.Error != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error)
		}
		
		return nil
	}