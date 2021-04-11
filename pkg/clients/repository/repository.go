package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/rs/zerolog"

	"github.com/krok-o/krok/pkg/models"

	"github.com/krok-o/krokctl/pkg/clients"
)

const (
	timeOutInSeconds = 10
	repositoryURI    = "/rest/api/1/krok/repository"
	repositoriesURI  = "/rest/api/1/krok/repositories"
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutInSeconds)*time.Second)
	defer cancel()

	b, err := json.Marshal(repo)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse repository")
		return nil, err
	}

	u, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse address")
		return nil, err
	}

	result := &models.Repository{}
	u.Path = path.Join(u.Path, repositoryURI)
	code, err := c.Handler.Post(ctx, b, u.String(), result)
	if err != nil {
		c.Logger.Debug().Err(err).Int("code", code).Msg("Failed to get result.")
		return nil, err
	}
	if code > 299 || code < 200 {
		c.Logger.Error().Str("url", u.String()).Int("code", code).Msg("Return code was not OK")
		return nil, fmt.Errorf("return code was not OK %d", code)
	}
	return repo, nil
}
