package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"freep.space/fsp/telegram"
)

func Download(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	fileID := strings.ReplaceAll(r.URL.String(), "/download/", "")

	downloadPath := telegram.DownloadFromTelegram(fileID)

	resp, err := os.Open(downloadPath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Close()

	info, err := resp.Stat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+info.Name())
	w.Header().Set("Content-Length", fmt.Sprint(info.Size()))

	_, err = io.Copy(w, resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer os.Remove(downloadPath)
}
