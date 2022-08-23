package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type FlightPath [2]string
type FlightPaths = []FlightPath

// Info is used for health checking as well as outputting current app version
func (h Handler) Path() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var input FlightPaths
		params, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(params, &input)
		if err != nil {
			writeResponse(w, http.StatusBadRequest, "invalid input data")
			return
		}

		var flag = map[string]int{}
		for _, path := range input {
			flag[path[0]] += 1
			flag[path[1]] -= 1
		}

		var result [2]string

		for key, value := range flag {
			if value == 1 {
				if result[0] != "" {
					writeResponse(w, http.StatusBadRequest, "invalid input data")
					return
				}
				result[0] = key
			} else if value == -1 {
				if result[1] != "" {
					writeResponse(w, http.StatusBadRequest, "invalid input data")
					return
				}
				result[1] = key
			}
		}

		writeResponse(w, http.StatusOK, result)
	}
}
