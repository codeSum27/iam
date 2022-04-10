package api

import (
	"net/http"

	"github.com/codeSum27/iam/pkg/global/db"
	"github.com/codeSum27/iam/pkg/model"
	"github.com/labstack/echo/v4"
)

// Get Users
// (GET /users)
	func (i *IamServer) GetUsers(ctx echo.Context) error {
		db := db.GetDBClient()

		users := []model.User{}
		err := db.Debug().Model(&model.User{}).Find(&users).Error
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}

		return ctx.JSON(http.StatusOK, users)
	}
	// Get User
	// (GET /users/{id})
	func (i *IamServer) GetUserById(ctx echo.Context, id string) error {
		db := db.GetDBClient()

		user := model.User{}
		err := db.Debug().Model(&model.User{}).Where("id = ?", id).Take(&user).Error
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}

		return ctx.JSON(http.StatusOK, user)
	}
	// Update User
	// (PUT /users/{id})
	func (i *IamServer) UpdateUserById(ctx echo.Context, id string) error {
		
		return nil
	}
	// Delete User
	// (DELETE /users/{id})
	func (i *IamServer) DeleteUserById(ctx echo.Context, id string) error {
		db := db.GetDBClient()

		user := model.User{}
		err := db.Debug().Model(&model.User{}).Where("id = ?", id).Take(&user).Delete(&user).Error
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}

		return nil
	}