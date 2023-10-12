package server

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

// Max file size is 10 MB
const ALLOW_FILE_SZIE = 50 * 1024 * 1024

func Upload(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	if r.ContentLength > ALLOW_FILE_SZIE {
		w.WriteHeader(http.StatusForbidden)
		if _, err := w.Write([]byte("Allowed size is 50MB\n")); err != nil {
			log.Println(err)
		}
		return
	}

	if err := r.ParseMultipartForm(50 << 20); err != nil {
		log.Println(err)
	}

	var formFile string
	for k := range r.MultipartForm.File {

		formFile = k

	}

	file, headers, err := r.FormFile(formFile)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	var nBytes, nChunks int = 0, 0

	fileName := "./downloads/" + uuid.New().String() + "_" + headers.Filename
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

	if _, err := w.Write([]byte("File SuccessFully Uploaded\n")); err != nil {
		log.Println(err)
	}

	return

}
