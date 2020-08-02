package service

import (
	"context"
	"github.com/drugs-4-3/ingrid_task/models"
	"github.com/drugs-4-3/ingrid_task/osrm"
	"sort"
)

type Service struct {
	client *osrm.Client
}

type RoutesRequestParams struct {
	Ctx          context.Context
	Source       models.Coordinates
	Destinations []models.Coordinates
}

type RoutesResponse struct {
	Source string
	Routes []Route
}

type Route struct {
	Destination string
	Duration    float32
	Distance    float32
}

type clientResponse struct {
	Error error
	Response *osrm.RouteResponse
	Destination models.Coordinates
}

func (s Service) GetRoutes(params RoutesRequestParams) (*RoutesResponse, error) {
	response := RoutesResponse{}
	responseChan := make(chan clientResponse, len(params.Destinations))
	defer close(responseChan)

	for _, dest := range params.Destinations {
		dest := dest
		go func() {
			resp, err := s.client.GetRoute(params.Source, dest)
			responseChan <- clientResponse{
				Error:    err,
				Response: resp,
				Destination: dest,
			}
		}()
	}

	var resultErr error = nil
	for range params.Destinations {
		clResp := <- responseChan
		if clResp.Error != nil {
			resultErr = clResp.Error
		} else {
			response.Routes = append(response.Routes, Route{
				Destination: clResp.Destination.ToString(),
				Duration:    clResp.Response.Routes[0].Duration,
				Distance:    clResp.Response.Routes[0].Distance,
			})
		}
	}
	if resultErr != nil {
		return nil, resultErr
	}

	response.Source = params.Source.ToString()
	response.Routes = sortByEffort(response.Routes)
	return &response, nil
}

func sortByEffort(routes []Route) []Route {
	sort.Sort(ByEffort(routes))
	return routes
}
