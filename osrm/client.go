package osrm

import (
	"encoding/json"
	"fmt"
	"github.com/drugs-4-3/ingrid_task/models"
	"net/http"
)

type Client struct {
	Config Config
	Client http.Client
}

type Config struct {
	Host string
}

type RouteResponse struct {
	Routes []Route
	Code string
}

type Route struct {
	Duration float32
	Distance float32
}

func (c Client) GetRoute(src, dst models.Coordinates) (*RouteResponse, error){
	u := c.Config.Host + "/route/v1/driving/" + src.ToString() + ";" + dst.ToString()
	request, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	response, err := c.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("osrm response status != 200: got %d. message: %s", response.StatusCode)
	}
	responseStruct := RouteResponse{}
	if err = json.NewDecoder(response.Body).Decode(&responseStruct); err != nil {
		return nil, err
	}
	if responseStruct.Code != "Ok" {
		return nil, fmt.Errorf("OSRM response code != \"Ok\"")
	}

	return &responseStruct, nil
}
