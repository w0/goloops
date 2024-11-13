package client

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type Client struct {
	httpClient http.Client
	inProgress int
	remaining  int
}

// Create the out file
func OutFile(path string, filename string) os.File {
	u, _ := user.Current()
	filepath := filepath.Join(u.HomeDir, filename)
	out, err := os.Create(filepath)

	if err != nil {
		panic(err)
	}
	return *out
}

// Download File
func DownloadFile(url string) error {
	filename := url[strings.LastIndex(url, "/")+1:]

	outfile := OutFile("~/Downloads", filename)

	defer outfile.Close()

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("oop: %s", resp.Status)
	}

	_, err = io.Copy(&outfile, resp.Body)

	if err != nil {
		return err
	}

	return nil
}
