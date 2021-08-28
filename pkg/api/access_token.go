package api

import (
	apiError "github.com/codeSum27/iam/pkg/api/error"
	"github.com/codeSum27/iam/pkg/common"
)

type AccessTokenContoller struct {
	TokenString string
}

func (c *AccessTokenContoller)	VerifyToken(userId string) (*Token,error) {

	clinet := GetRedisClient()
	result := clinet.Get(c.TokenString)

	if result.Err() !=nil {
		return nil,apiError.NewExpiredAccessTokenError("Access token is Invalid")
	}
	err := verifyToken(common.Cnf.Token.AccessSecret, c.TokenString, userId)
	if err != nil{
		return nil, apiError.NewInvalidAccessTokenError("Access Token is invalid")
	}

	var tokenString = c.TokenString
	return &Token{
		AccessToken: &tokenString,
	},nil
}