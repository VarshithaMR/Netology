package server

import (
	"net/http"

	"Netology/pythagoreancalculation/handler"
)

func HandlePythogoreanCalculation(rw http.ResponseWriter, r *http.Request, calculate handler.Calculate) {

	switch r.Method {
	case http.MethodPost:
		calculate.StartCalculateDistances(r, rw)
	}
}
