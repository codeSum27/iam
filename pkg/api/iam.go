//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=api --generate server -o server.gen.go iam-spec.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=api --generate types -o type.gen.go iam-spec.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=api --generate spec -o spec.gen.go iam-spec.yaml


package api

import (
	apiError "github.com/codeSum27/iam/pkg/api/error"
	"github.com/labstack/echo/v4"
	"reflect"
	"sync"
)

type IamServer struct {
	Lock   sync.Mutex
}


func NewIamServer()  *IamServer {
	return &IamServer{}
}

func handleTokenError(ctx echo.Context, err error) error {

	if reflect.TypeOf(err) == reflect.TypeOf(apiError.NewExpiredAccessTokenError("")) {
		petErr := Error{
			Code:    int32(404),
			Message: err.Error(),
		}
		return ctx.JSON(404, petErr)
	}

	return ctx.JSON(400, err)
}

func sendCreateTokenError(ctx echo.Context, code int, message string) error {
	petErr := Error{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, petErr)
	return err
}