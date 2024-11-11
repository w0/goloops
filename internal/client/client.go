package client

import (
	"net/http"
	"os"
)

type Client struct {
	httpClient http.Client
	inProgress int
	remaining  int
}

// Create the out file
func OutFile(filepath string) os.File {
	return os.File{}
}

// Download File
func (c *Client) DownloadFile()

//
