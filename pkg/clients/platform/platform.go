package platform

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/krok-o/krok/pkg/models"
	"github.com/rs/zerolog"

	"github.com/krok-o/krokctl/pkg/clients"
)

const (
	timeOutInSeconds = 10
	platformURI      = "/supported-platforms"
)

// NewClient creates a new platform provider.
func NewClient(address string, client *http.Client, token string, log zerolog.Logger) *Client {
	return &Client{
		Address: address,
		Logger:  log,
		Handler: clients.NewHandler(*client, token, log),
	}
}

// Client contains methods for command related resource actions.
type Client struct {
	Address string
	Logger  zerolog.Logger
	Handler *clients.Handler
}

// List repositories.
func (c *Client) List() ([]models.Platform, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutInSeconds)*time.Second)
	defer cancel()

	u, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse address")
		return nil, err
	}

	var result []models.Platform
	u.Path = path.Join(u.Path, platformURI)
	code, err := c.Handler.MakeRequest(ctx, http.MethodGet, u.String(), clients.WithOutput(&result))
	if err != nil {
		c.Logger.Debug().Err(err).Int("code", code).Msg("Failed to get result.")
		return nil, err
	}
	if code > 299 || code < 200 {
		c.Logger.Error().Str("url", u.String()).Int("code", code).Msg("Return code was not OK")
		return nil, fmt.Errorf("return code was not OK %d", code)
	}
	return result, nil
}
