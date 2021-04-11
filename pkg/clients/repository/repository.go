package repository

import (
	"net/http"

	"github.com/rs/zerolog"

	"github.com/krok-o/krok/pkg/models"

	"github.com/krok-o/krokctl/pkg/clients"
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
		Logger:  log,
		Handler: clients.NewHandler(*client, log),
	}
}

// Client contains methods for repository related resource actions.
type Client struct {
	Address string
	Logger  zerolog.Logger
	Handler *clients.Handler
}

// Create creates a repository resource.
func (c *Client) Create(repo *models.Repository) (*models.Repository, error) {

	return repo, nil
}
