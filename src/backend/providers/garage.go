package providers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/danmurphy1217/invoice-generator/custom_errors"
	"github.com/danmurphy1217/invoice-generator/models"
)


const baseUrl = "https://garage-backend.onrender.com"

type Garage struct {
	Timeout time.Duration;
	TruckId string;
	Endpoint string;
	Headers map[string]string;
}

func (g *Garage) GetListing() (*models.GetListingResponse, error) {
    reqGetListingBytes, err := json.Marshal(map[string]string{
		"id": g.TruckId,
	})
	if err != nil {
		return &models.GetListingResponse{}, fmt.Errorf("failed to unmarshal request")
	}


    httpClient := NewHttpClient(g.Timeout)
    getListingResponseBytes, err := httpClient.Post(
        baseUrl + fmt.Sprintf("/%s", g.Endpoint),
        "application/json",
        reqGetListingBytes,
        g.Headers,
    )
	if err != nil {
		return &models.GetListingResponse{}, fmt.Errorf("failed to get listings from garage")
	}

    var respGetListing models.GetListingResponse
    err = json.Unmarshal(getListingResponseBytes, &respGetListing)
	if err != nil {
		return &models.GetListingResponse{}, fmt.Errorf("failed to unmarshal response to get listings request")
	}

    if (respGetListing.Result == nil) || (respGetListing.Result.Listing == nil) {
        // return not found error if no valid result is returned
        // NOTE: this is important since garage API is returning 200 OK if
        // nothing is found
        return &models.GetListingResponse{}, custom_errors.NewTruckNotFoundError()
    }

	return &respGetListing, nil
}