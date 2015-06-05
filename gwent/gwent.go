package gwent

import (
	"strings"
	"net/http"
	"encoding/json"
)

const (
	GWENT_VERSION = "0.1a"
	GWENT_PATH_QUIT = "/quit"
	GWENT_PATH_API = "/api/"
	
	GWENT_PATH_API_LEN = len(GWENT_PATH_API)
)

func ProcessRestApi(w http.ResponseWriter, r *http.Request) {	
	segments := strings.Split(r.URL.Path[GWENT_PATH_API_LEN:],"/")
	
	if len(segments) == 0 {
		//Handle nothiiing
		w.Write([]byte(":)"))
	} else {
		switch(segments[0]){
			case "cards":
				b, _ := json.Marshal(AllTemplates)
				w.Write(b)
				
			default:
				w.Write([]byte("Unknown command"))						
		}
	}
}