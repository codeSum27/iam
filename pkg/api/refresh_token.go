package api

import (
	apiError "github.com/codeSum27/iam/pkg/api/error"
	"github.com/codeSum27/iam/pkg/common"
)

type RefreshTokenContoller struct {
	TokenString string
}

func (c *RefreshTokenContoller)	VerifyToken(userId string) (*Token,error) {
	clinet := GetRedisClient()
	result := clinet.Get(c.TokenString)

	if result.Err() !=nil {
		return nil,apiError.NewExpiredRefreshTokenError("Refresh token is Invalid")
	}

	err := verifyToken(common.Cnf.Token.RefreshSecret, c.TokenString, userId)
	if err != nil{
		return nil, apiError.NewInvalidRefreshTokenError("Refresh Token is invalid")
	}
	user, err := getUser(userId)
	if err != nil{
		return nil, apiError.NewInvalidRefreshTokenError("Cannot get user for creating Token for refresh")
	}
	token, err := generateJWT(user)
	if err != nil{
		return nil, apiError.NewInvalidRefreshTokenError("Cannot create Token for refresh")
	}
	err =  saveTokenInfo(token)
	if err != nil{
		return nil, apiError.NewInvalidRefreshTokenError("Cannot save Token for refresh")
	}
	return token, nil
}
