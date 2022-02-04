package api

import (
	"fmt"
	apiError "github.com/codeSum27/iam/pkg/api/error"
	"github.com/codeSum27/iam/pkg/common"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type TokenContoller interface {
	VerifyToken(userId string) (*Token,error)
}

func (i *IamServer) CreateToken(ctx echo.Context) error {
	i.Lock.Lock()
	defer i.Lock.Unlock()

	var user User
	err := ctx.Bind(&user)
	if err != nil {
		return sendCreateTokenError(ctx, http.StatusBadRequest, "Invalid format for Issue Token")
	}

	err = verifyUser(&user)
	if err != nil {
		return sendCreateTokenError(ctx, http.StatusForbidden, err.Error())
	}

	ctx.Logger().Print("Successfully verify user.")

	token, err := generateJWT(&user)
	if err!=nil{
		return sendCreateTokenError(ctx, http.StatusForbidden, err.Error())
	}

	err = ctx.JSON(http.StatusCreated, token)
	if err != nil {
		// Something really bad happened, tell Echo that our handler failed
		return err
	}
	return nil
}

// (GET /tokens/{id})
func (i *IamServer) GetToken(ctx echo.Context, id string, params GetTokenParams) error {

	tokenContoller := checkTokenHeader(params)

	if tokenContoller == nil{
		return sendCreateTokenError(ctx, http.StatusBadRequest, "Invalid Header. Access token or Refresh Token Must be setted")
	}

	token, err := tokenContoller.VerifyToken(id)
	if err != nil {
		return handleTokenError(ctx,err)
	}

	err = ctx.JSON(http.StatusOK, token)
	if err != nil {
		// Something really bad happened, tell Echo that our handler failed
		return err
	}
	return nil
}

func generateJWT(user *User) (*Token, error) {

	var mySigningKey = []byte(common.Cnf.Token.AccessSecret)
	tokenExpiredAt := time.Now().Add(time.Minute * 30).Unix()
	refreshTokenExpiredAt := time.Now().Add(time.Minute * 30).Unix()

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = user.Id
	claims["email"] = user.Email
	claims["exp"] = tokenExpiredAt

  at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  tokenString, err := at.SignedString(mySigningKey)
	tokenType := "jwt"

	var mySigningRefreshKey = []byte(common.Cnf.Token.RefreshSecret)
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["id"] = user.Id
	rtClaims["exp"] = refreshTokenExpiredAt

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refreshTokenString, err := rt.SignedString([]byte(mySigningRefreshKey))
	if err != nil {
		return nil, err
	}

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return nil, err
	}

	var expiredAt = int(tokenExpiredAt)
	var refreshAt = int(refreshTokenExpiredAt)

	tokenInfo := &Token{
		AccessToken: &tokenString,
		RefreshToken: &refreshTokenString,
		ExpiresAt: &expiredAt,
		RefreshAt: &refreshAt,
		TokenType: &tokenType,
	}

	err = saveTokenInfo(tokenInfo)
	if err != nil{
		return nil, err
	}

	return tokenInfo, nil
}

func saveTokenInfo(token *Token) error  {
	redisClient := GetRedisClient()
	expiredAt := int64(*token.ExpiresAt)
	refreshAt := int64(*token.ExpiresAt)
	at := time.Unix(expiredAt, 0) //converting Unix to UTC
	rt := time.Unix(refreshAt, 0)
	now := time.Now()

	err := redisClient.Set(*token.AccessToken, true, at.Sub(now)).Err()
	if err != nil {
		return apiError.NewRedisClientError("Cannot save Access token")
	}
	err = redisClient.Set(*token.RefreshToken, true, rt.Sub(now)).Err()
	if err != nil {
		return apiError.NewRedisClientError("Cannot save Refresh token")
	}
	return nil
}

func checkTokenHeader(params GetTokenParams)  TokenContoller {
	if params.AccessToken != nil {
		return  &AccessTokenContoller{
			TokenString: *params.AccessToken,
		}
	}

	if params.RefreshToken != nil {
		return  &RefreshTokenContoller{
			TokenString: *params.RefreshToken,
		}
	}
	return nil
}

func verifyToken(secret string,tokenString string, userId string )  error{
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		return []byte(secret), nil
	})
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) || claims["id"] != userId {
		return apiError.NewTokenVerifyFailError("Token is invalid")
	}
	return nil
}