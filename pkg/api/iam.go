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