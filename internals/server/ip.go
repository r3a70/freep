package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	IpAddr string `json:"ip_addr"`
}

func Ip(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Header)

	ipAddr := response{
		IpAddr: r.RemoteAddr,
	}

	w.Header().Add("application", "json")
	resp, _ := json.Marshal(ipAddr)

	w.Write(resp)

}
