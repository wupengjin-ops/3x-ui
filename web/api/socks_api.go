
package api

import (
	"encoding/json"
	"net/http"
	"your_project/service/socks"
)

func SocksList(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(socks.List())
}
