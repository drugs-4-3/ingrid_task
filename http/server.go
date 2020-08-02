package http

import (
	"git.limango.tech/osp/logger"
	"github.com/drugs-4-3/ingrid_task/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func NewServer() (*http.Server, error) {
	router := mux.NewRouter()
	listenOn, err := utils.GetEnvOrErr("LISTEN")
	if err != nil {
		return nil, err
	}

	sWriteTimeoutStr, err := utils.GetEnvOrErr("SERVER_WRITE_TIMEOUT")
	if err != nil {
		return nil, err
	}
	sWriteTimeout, err := strconv.Atoi(sWriteTimeoutStr)
	if err != nil {
		return nil, err
	}

	sReadTimeoutStr, err := utils.GetEnvOrErr("SERVER_READ_TIMEOUT")
	if err != nil {
		return nil, err
	}
	sReadTimeout, err := strconv.Atoi(sReadTimeoutStr)
	if err != nil {
		logger.Error(err)
	}

	SetUpEndpoints(router)
	return &http.Server{
		Addr:         listenOn,
		Handler:      router,
		WriteTimeout: time.Duration(sWriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(sReadTimeout) * time.Second,
	}, nil
}

func SetUpEndpoints(router *mux.Router) {
	router.Methods("GET").Path("/routes").HandlerFunc(GetRoutes)
}
