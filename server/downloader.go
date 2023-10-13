package server

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"freep.space/fsp/telegram"
)

func Download(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	fileID := strings.ReplaceAll(r.URL.String(), "/download/", "")

	downloadUrl := telegram.DownloadFromTelegram(fileID)

	resp, err := http.Get(downloadUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Disposition", "attachment; filename="+strings.Split(downloadUrl, "/")[6])
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", fmt.Sprint(resp.ContentLength))

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
