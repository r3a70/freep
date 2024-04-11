package server

import (
	"encoding/json"
	"net/http"
)

type response struct {
	IpAddr string `json:"ip_addr"`
}

func Ip(w http.ResponseWriter, r *http.Request) {

	ipAddr := response{
		IpAddr: r.Header.Get("X-Real-Ip"),
	}

	w.Header().Add("application", "json")
	resp, _ := json.Marshal(ipAddr)

	w.Write(resp)

}
