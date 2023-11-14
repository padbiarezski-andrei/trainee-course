package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

var (
	ErrInvalidURL       = errors.New("invalid URL")
	ErrConnectionFailed = errors.New("connection failed")
	ErrDownloadFailed   = errors.New("cannot download file")
	ErrFileNotFound     = errors.New("file not found")
)

func downloadFile(fileURL, filename string) (err error) {
	u, err := url.ParseRequestURI(fileURL) // 1
	if err != nil {
		return errors.Join(ErrInvalidURL, err)
	}

	resp, err := http.Get(u.String())
	if err != nil { // 2
		return errors.Join(ErrConnectionFailed, err)
	}

	if resp.StatusCode == http.StatusInternalServerError { // 3
		return ErrDownloadFailed
	}

	if resp.StatusCode == http.StatusNotFound { // 4
		return ErrFileNotFound
	}

	outFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot create file %q, error %w", filename, err)
	}

	defer func() {
		err = outFile.Close()
	}()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("cannot write to file %q, error %w", filename, err)
	}

	return nil
}

func main() {
	var url = flag.String("url", "", "URL to the file")
	var path = flag.String("output", "", "path to the file")

	flag.Parse()

	err := downloadFile(*url, *path)
	if err != nil {

		switch {
		case errors.Is(err, ErrInvalidURL):
			fmt.Printf("Invalide URL %q. Please check the URL correctness. err:%s\n", *url, err)

		case errors.Is(err, ErrConnectionFailed):
			fmt.Printf("Cannot connect to %q. Please check your connection settings or try to again later. err:%s\n", *url, err)

		case errors.Is(err, ErrDownloadFailed):
			fmt.Printf("Error occured while downloading file %q. Please check if file is available. err:%s\n", *url, err)

		case errors.Is(err, ErrFileNotFound):
			fmt.Printf("No such file found. Please check if the file available or URL %q is correct. err:%s\n", *url, err)

		default:
			fmt.Println("error occurred: ", err)
		}

		os.Exit(1)
	}
}
