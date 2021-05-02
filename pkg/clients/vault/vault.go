package vault

import (
	"context"
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
	vaultURI         = "/rest/api/1/krok/vault/secret"
	vaultListURI     = "/rest/api/1/krok/vault/secrets"
)

// NewClient creates a new vault provider.
func NewClient(address string, client *http.Client, token string, log zerolog.Logger) *Client {
	return &Client{
		Address: address,
		Logger:  log,
		Handler: clients.NewHandler(*client, token, log),
	}
}

// Client contains methods for vault related resource actions.
type Client struct {
	Address string
	Logger  zerolog.Logger
	Handler *clients.Handler
}

// Get returns a secret resource.
func (c *Client) Get(name string) (*models.VaultSetting, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutInSeconds)*time.Second)
	defer cancel()

	u, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse address")
		return nil, err
	}

	result := models.VaultSetting{}
	u.Path = path.Join(u.Path, vaultURI, name)
	code, err := c.Handler.MakeRequest(ctx, http.MethodGet, nil, u.String(), &result)
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

// List secrets.
func (c *Client) List() ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutInSeconds)*time.Second)
	defer cancel()

	u, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse address")
		return nil, err
	}

	var result []string
	u.Path = path.Join(u.Path, vaultListURI)
	code, err := c.Handler.MakeRequest(ctx, http.MethodPost, nil, u.String(), &result)
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
