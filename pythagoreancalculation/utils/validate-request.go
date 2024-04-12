package utils

import (
	"encoding/json"
	"io"

	"github.com/opentracing/opentracing-go/log"

	"Netology/pythagoreancalculation/models"
)

func GetRequestBody(body io.ReadCloser) (requestBody []models.Coordinates, err error) {
	decoder := json.NewDecoder(body)

	if err = decoder.Decode(&requestBody); err != nil {
		log.Error(err)
		return
	}

	return
}
