package api

import "github.com/labstack/echo/v4"

// Get Users
// (GET /users)
	func (i *IamServer) GetUsers(ctx echo.Context) error {

		return nil
	}
	// Delete User
	// (DELETE /users/{id})
	func (i *IamServer) DeleteUserById(ctx echo.Context, id string) error {

		return nil
	}
	// Get User
	// (GET /users/{id})
	func (i *IamServer) GetUserById(ctx echo.Context, id string) error {

		return nil
	}
	// Update User
	// (PUT /users/{id})
	func (i *IamServer) UpdateUserById(ctx echo.Context, id string) error {

		return nil
	}
