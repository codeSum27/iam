// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package api

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// Token defines model for Token.
type Token struct {
	// Access Token for user
	AccessToken *string `json:"access_token,omitempty"`

	// Expire time
	ExpiresAt *int `json:"expires_at,omitempty"`

	// Expire time
	RefreshAt *int `json:"refresh_at,omitempty"`

	// Refresh Token for access token
	RefreshToken *string `json:"refresh_token,omitempty"`

	// Type of access_token.
	TokenType *string `json:"token_type,omitempty"`
}

// User defines model for User.
type User struct {
	// Email for user
	Email string `json:"email"`

	// UUID for user
	Id *string `json:"id,omitempty"`

	// Nickname of user
	NickName *string `json:"nickName,omitempty"`

	// Password for user
	Password string `json:"password"`
}

// CreateTokenJSONBody defines parameters for CreateToken.
type CreateTokenJSONBody User

// GetTokenParams defines parameters for GetToken.
type GetTokenParams struct {
	AccessToken  *string `json:"Access-token,omitempty"`
	RefreshToken *string `json:"Refresh-token,omitempty"`
}

// PostUsersJSONBody defines parameters for PostUsers.
type PostUsersJSONBody User

// DeleteUserByIdParams defines parameters for DeleteUserById.
type DeleteUserByIdParams struct {
	AccessToken string `json:"Access-token"`
}

// GetUserByIdParams defines parameters for GetUserById.
type GetUserByIdParams struct {
	AccessToken string `json:"Access-token"`
}

// UpdateUserJSONBody defines parameters for UpdateUser.
type UpdateUserJSONBody User

// UpdateUserParams defines parameters for UpdateUser.
type UpdateUserParams struct {
	AccessToken string `json:"Access-token"`
}

// CreateTokenJSONRequestBody defines body for CreateToken for application/json ContentType.
type CreateTokenJSONRequestBody CreateTokenJSONBody

// PostUsersJSONRequestBody defines body for PostUsers for application/json ContentType.
type PostUsersJSONRequestBody PostUsersJSONBody

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody UpdateUserJSONBody