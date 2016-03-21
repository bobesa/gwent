package gwent

import (
	"encoding/json"
	"net/http"
	"strings"
)

const (
	// Version of Gwent server
	Version = "0.1a"
	// PathQuit defines http path to stop the Gwent server
	PathQuit = "/quit"
	// PathAPI defines http path to api of Gwent server
	PathAPI = "/api/"

	pathAPILen = len(PathAPI)
)

// ProcessRestAPI is our http call
func ProcessRestAPI(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(r.URL.Path[pathAPILen:], "/")

	if len(segments) == 0 {
		//Handle nothing
		w.Write([]byte(":)"))
	} else {
		switch segments[0] {
		case "cards":
			b, _ := json.Marshal(AllTemplates)
			w.Write(b)

		default:
			w.Write([]byte("Unknown command"))
		}
	}
}
