package server

import (
	"net/http"

	"github.com/tharouet/tserve/logger"
)

func Init(serverName, port string, r Router) (err error) {
	l := logger.New()
	log := l.Log
	log.Debug("new tserve server: %s", serverName)

	// set port
	if port == "" {
		port = "8080"
	}

	log.Debug("tserve server (%s)- port:%s", serverName, port)
	err = http.ListenAndServe(":"+port, r.Router)
	if err != nil {
		log.Fatal("tserve - ListenAndServe error: ", err)
	}
	return
}
