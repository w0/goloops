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

// Create the out file
func OutFile(path string, filename string) os.File {
	u, _ := user.Current()
	fp := filepath.Join(u.HomeDir, "Downloads/goloops", filename)
	dir := filepath.Dir(fp)

	if _, ok := os.Stat(dir); os.IsNotExist(ok) {
		os.Mkdir(dir, 0755)
	}

	out, err := os.Create(fp)

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
