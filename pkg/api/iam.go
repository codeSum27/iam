//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=api --generate server -o server.gen.go spec.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=api --generate types -o type.gen.go spec.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=api --generate spec -o spec.gen.go spec.yaml

package api

import (
	"sync"
)

type IamServer struct {
	Lock   sync.Mutex
}


func NewIamServer()  *IamServer {
	return &IamServer{}
}