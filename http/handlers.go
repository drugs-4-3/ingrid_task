package http

import (
	"encoding/json"
	"fmt"
	"github.com/drugs-4-3/ingrid_task/models"
	"github.com/drugs-4-3/ingrid_task/service"
	"log"
	"net/http"
)

func GetRoutes(w http.ResponseWriter, r *http.Request) {
	source, ok := r.URL.Query()["src"]
	if !ok || len(source) == 0 {
		respond(w, http.StatusBadRequest, "missing \"src\" parameter")
		return
	}
	destination, ok := r.URL.Query()["dst"]
	if !ok || len(destination) == 0 {
		respond(w, http.StatusBadRequest, "missing \"dst\" parameter")
		return
	}
	srcCoordinates, err := models.NewCoordinateFromString(source[0])
	if err != nil {
		respond(w, http.StatusBadRequest, "incorrect \"src\" parameter: " + err.Error())
		return
	}

	dstCoordinates, err := models.NewCoordinatesCollectionFromString(destination)
	if err != nil {
		respond(w, http.StatusBadRequest, "incorrect \"dst\" parameter: " + err.Error())
		return
	}

	routesResp, err := service.GetService().GetRoutes(service.RoutesRequestParams{
		Ctx:         r.Context(),
		Source:      srcCoordinates,
		Destinations: dstCoordinates,
	})
	if err != nil {
		respond(w, http.StatusInternalServerError, "could not load routes: " + err.Error())
		return
	}
	fmt.Printf("%+v\n", *routesResp)


	w.WriteHeader(http.StatusOK)
	// todo: write json response
	//if _, err := w.Write([]byte("not implemented")); err != nil {
	//	log.Println("error writing response")
	//}
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func respond(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	responseObj := Response{
		Status:  status,
		Message: message,
	}
	payload, err := json.Marshal(responseObj)
	if err != nil {
		log.Println("cannot marshal response: " + err.Error())
		return
	}
	_, err = w.Write(payload)
	if err != nil {
		log.Println("cannot write response: " + err.Error())
	}
}