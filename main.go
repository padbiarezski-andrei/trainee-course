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

// ///////////////////////////////////////////////////////////////
type ErrInvalidURL struct {
	url string
	err error
}

func newErrInvalidURL(url string, err error) *ErrInvalidURL {
	return &ErrInvalidURL{
		url: url,
		err: err,
	}
}

func (e *ErrInvalidURL) Error() string {
	return fmt.Sprintf("invalid URL %q", e.url)
}

func (e *ErrInvalidURL) Unwrap() error {
	return e.err
}

// ///////////////////////////////////////////////////////////////
type ErrConnectionFailed struct {
	url string
	err error
}

func newErrConnectionFailed(url string, err error) *ErrConnectionFailed {
	return &ErrConnectionFailed{
		url: url,
		err: err,
	}
}

func (e *ErrConnectionFailed) Error() string {
	return fmt.Sprintf("cannot connect to %q", e.url)
}

func (e *ErrConnectionFailed) Unwrap() error {
	return e.err
}

// ///////////////////////////////////////////////////////////////
type ErrDownloadFailed struct {
	url string
	err error
}

func newErrDownloadFailed(url string, err error) *ErrDownloadFailed {
	return &ErrDownloadFailed{
		url: url,
		err: err,
	}
}

func (e *ErrDownloadFailed) Error() string {
	return fmt.Sprintf("cannot download file %q", e.url)
}

func (e *ErrDownloadFailed) Unwrap() error {
	return e.err
}

// ///////////////////////////////////////////////////////////////
type ErrFileNotFound struct {
	url string
	err error
}

func newErrFileNotFound(url string, err error) *ErrFileNotFound {
	return &ErrFileNotFound{
		url: url,
		err: err,
	}
}

func (e *ErrFileNotFound) Error() string {
	return fmt.Sprintf("not such file %q", e.url)
}

func (e *ErrFileNotFound) Unwrap() error {
	return e.err
}

// ///////////////////////////////////////////////////////////////

func downloadFile(fileURL, filename string) (err error) {
	u, err := url.ParseRequestURI(fileURL) // 1
	if err != nil {
		return newErrInvalidURL(fileURL, err)
	}

	resp, err := http.Get(u.String())
	var urlErr *url.Error        // ?
	if errors.As(err, &urlErr) { // 2
		return newErrConnectionFailed(fileURL, err)
	}

	if resp.StatusCode == http.StatusInternalServerError { // 3
		return newErrDownloadFailed(fileURL, err)
	}

	if resp.StatusCode == http.StatusNotFound { // 4
		return newErrFileNotFound(fileURL, err)
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
		var errInvalidURL *ErrInvalidURL
		var errConnectionFailed *ErrConnectionFailed
		var errDownloadFailed *ErrDownloadFailed
		var errNotFOund *ErrFileNotFound

		switch {
		case errors.As(err, &errInvalidURL):
			fmt.Printf("Invalide URL. Please check the URL correctness. err:%s\n", err)

		case errors.As(err, &errConnectionFailed):
			fmt.Printf("Cannot connect to %q. Please check your connection settings or try to again later. err:%s\n", *url, err)

		case errors.As(err, &errDownloadFailed):
			fmt.Printf("Error occured while downloading file %q. Please check if file is available. err:%s\n", *url, err)

		case errors.As(err, &errNotFOund):
			fmt.Printf("No such file found. Please check if the file available or URL %q is correct.\n", *url)

		default:
			fmt.Println("error occurred: ", err)
		}

		os.Exit(1)
	}
}
