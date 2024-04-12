package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-resty/resty/v2"

	"Netology/pythagoreancalculation"
	"Netology/pythagoreancalculation/utils"
)

const (
	contentTypeKey   = "Content-Type"
	contentTypeValue = "application/json; charset=utf-8"
)

type Calculate interface {
	StartCalculateDistances(*http.Request, http.ResponseWriter)
}

type calculateDomainHandlerService struct {
	httpClient *resty.Client
}

func (c *calculateDomainHandlerService) StartCalculateDistances(request *http.Request, response http.ResponseWriter) {
	points, err := utils.GetRequestBody(request.Body)
	if err != nil {
		WriteResponse(response, "Request coordinates improper", 400)
		return
	}

	// check for minimum 2 coordinate sets to be found
	if len(points) < 2 {
		WriteResponse(response, "Cannot build Proper response - improper request body", 400)
		return
	}

	distancePairsResponse, err := pythagoreancalculation.CalculateDistance(points)
	if err != nil {
		WriteResponse(response, "Cannot build proper response", 400)
		return
	}

	WriteResponse(response, distancePairsResponse, http.StatusOK)

}

func WriteResponse(rw http.ResponseWriter, resp interface{}, responseCode int) {
	rw.WriteHeader(responseCode)
	rw.Header().Set(contentTypeKey, contentTypeValue)
	bytes, err := json.Marshal(resp)
	if err != nil {
		//TODO logging
	}
	rw.Write(bytes)
}

func NewCalculateDomainHandler() Calculate {
	return &calculateDomainHandlerService{
		httpClient: resty.New(),
	}
}
