package osrm

import (
	"github.com/drugs-4-3/ingrid_task/utils"
	"log"
	"net/http"
	"sync"
	"time"
)

var clientOnce sync.Once
var client *Client

func NewClient() *Client {
	clientOnce.Do(func() {
		osrmApiHost, err := utils.GetEnvOrErr("OSRM_API_HOST")
		if err != nil {
			log.Fatal(err)
		}
		client = &Client{
			Config: Config{
				Host: osrmApiHost,
			},
			Client: http.Client{
				Timeout: 3 * time.Second,
			},
		}
	})
	return client
}
