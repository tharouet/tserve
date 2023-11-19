package handler

import (
	"net/http"

	"github.com/tharouet/tserve/logger"
)

type Handle struct {
	Logger logger.Logit
}

type Handler interface {
	func(*Handle) (w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

func NewHandle() Handle {
	return Handle{
		Logger: logger.New(),
	}
}
