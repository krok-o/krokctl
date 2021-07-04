package auth

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"

	"github.com/krok-o/krok/pkg/models"
	"github.com/rs/zerolog"

	"github.com/krok-o/krokctl/pkg/clients"
)

const (
	timeOutInSeconds = 10
	apiKeyURI        = "/rest/api/1/krok/user/apikey"
	apiKeysURI       = "/rest/api/1/krok/user/apikeys"
)

// NewClient creates a new api key provider.
func NewClient(address string, log zerolog.Logger, handler clients.Handler) *Client {
	return &Client{
		Address: address,
		Logger:  log,
		Handler: handler,
	}
}

// Client contains methods for command related resource actions.
type Client struct {
	Address string
	Logger  zerolog.Logger
	Handler clients.Handler
}

// Create creates a repository resource.
func (c *Client) Create(name string) (*models.APIKey, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutInSeconds)*time.Second)
	defer cancel()

	u, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse address")
		return nil, err
	}

	result := models.APIKey{}
	u.Path = path.Join(u.Path, apiKeyURI, "generate", name)
	code, err := c.Handler.MakeRequest(ctx, http.MethodPost, u.String(), clients.WithOutput(&result))
	if err != nil {
		c.Logger.Debug().Err(err).Int("code", code).Msg("Failed to get result.")
		return nil, err
	}
	if code > 299 || code < 200 {
		c.Logger.Error().Str("url", u.String()).Int("code", code).Msg("Return code was not OK")
		return nil, fmt.Errorf("return code was not OK %d", code)
	}
	return &result, nil
}

// List api keys.
func (c *Client) List() ([]*models.APIKey, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutInSeconds)*time.Second)
	defer cancel()

	u, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse address")
		return nil, err
	}

	var result []*models.APIKey
	u.Path = path.Join(u.Path, apiKeysURI)
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

// Delete deletes an api key resource.
func (c *Client) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutInSeconds)*time.Second)
	defer cancel()

	u, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse address")
		return err
	}
	u.Path = path.Join(u.Path, apiKeyURI, "delete", strconv.Itoa(id))
	code, err := c.Handler.MakeRequest(ctx, http.MethodDelete, u.String())
	if err != nil {
		c.Logger.Debug().Err(err).Int("code", code).Msg("Failed to get result.")
		return err
	}
	if code > 299 || code < 200 {
		c.Logger.Error().Str("url", u.String()).Int("code", code).Msg("Return code was not OK")
		return fmt.Errorf("return code was not OK %d", code)
	}
	return nil
}

// Get returns a api key resource.
func (c *Client) Get(id int) (*models.APIKey, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutInSeconds)*time.Second)
	defer cancel()

	u, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse address")
		return nil, err
	}

	result := models.APIKey{}
	u.Path = path.Join(u.Path, apiKeyURI, strconv.Itoa(id))
	code, err := c.Handler.MakeRequest(ctx, http.MethodGet, u.String(), clients.WithOutput(&result))
	if err != nil {
		c.Logger.Debug().Err(err).Int("code", code).Msg("Failed to get result.")
		return nil, err
	}
	if code > 299 || code < 200 {
		c.Logger.Error().Str("url", u.String()).Int("code", code).Msg("Return code was not OK")
		return nil, fmt.Errorf("return code was not OK %d", code)
	}
	return &result, nil
}
