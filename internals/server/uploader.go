package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"freep.space/fsp/internals/constant"
	"freep.space/fsp/internals/telegram"
	"github.com/google/uuid"
)

func Upload(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		if _, err := w.Write([]byte("Method Not Allowed\n")); err != nil {
			log.Println(err)
		}
		return
	}

	if r.ContentLength > constant.ALLOW_FILE_SZIE {
		w.WriteHeader(http.StatusForbidden)
		if _, err := w.Write([]byte("Allowed size is " + fmt.Sprint(constant.ALLOW_FILE_SZIE/1024/1024, "MB\n"))); err != nil {
			log.Println(err)
		}
		return
	}

	if err := r.ParseMultipartForm(constant.MULTY_PART_MAX_SIZE); err != nil {
	}

	var formFile string
	for k := range r.MultipartForm.File {

		formFile = k

	}

	if formFile == "" {
		w.WriteHeader(http.StatusForbidden)
		if _, err := w.Write([]byte("File not found..")); err != nil {
			log.Println(err)
		}
		return

	}

	file, headers, err := r.FormFile(formFile)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	var nBytes, nChunks int = 0, 0
	fmt.Println("============================")

	rg := regexp.MustCompile(`[^A-Za-z0-9.-_]`)
	fileName := "./downloads/" + uuid.New().String() + "_" + rg.ReplaceAllString(headers.Filename, "_")
	createdFile, err := os.Create(fileName)
	if err != nil {
		log.Println(err)
	}
	defer createdFile.Close()

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(createdFile)

	buf := make([]byte, 0, 4*1024)
	for {
		n, err := reader.Read(buf[:cap(buf)])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		nChunks++
		nBytes += len(buf)

		if _, err := writer.Write(buf); err != nil {
			log.Println(err)
		}

		// process buf
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

	}
	writer.Flush()

	downloadUrl := telegram.UploadToTelegram(fileName)

	if _, err := w.Write([]byte(downloadUrl + "\n")); err != nil {
		log.Println(err)
	}

	defer func() {
		if err := os.Remove(fileName); err != nil {
			log.Println(err)
		}
	}()

}
