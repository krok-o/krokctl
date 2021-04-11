package repository

import (
	"net/http"

	"github.com/rs/zerolog"
)

const (
	timeOutInSeconds = 10
	repositoryURI    = "/rest/api/1/repository"
	repositoriesURI  = "/rest/api/1/repositories"
)

// NewClient creates a new repository provider.
func NewClient(address string, client *http.Client, log zerolog.Logger) *Client {
	return &Client{
		Address: address,
		Client:  client,
		Logger:  log,
	}
}

// Client contains methods for repository related resource actions.
type Client struct {
	Address string
	Client  *http.Client
	Logger  zerolog.Logger
}
