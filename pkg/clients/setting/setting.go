package setting

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"

	"github.com/rs/zerolog"

	"github.com/krok-o/krok/pkg/models"
	"github.com/krok-o/krokctl/pkg/clients"
)

const (
	timeOutInSeconds = 10
	settingURI       = "/rest/api/1/krok/command/setting"
	settingsURI      = "/rest/api/1/krok/command/"
)

// NewClient creates a new settings provider.
func NewClient(address string, client *http.Client, token string, log zerolog.Logger) *Client {
	return &Client{
		Address: address,
		Logger:  log,
		Handler: clients.NewHandler(*client, token, log),
	}
}

// Client contains methods for settings related resource actions.
type Client struct {
	Address string
	Logger  zerolog.Logger
	Handler *clients.Handler
}

// Create will create settings.
func (c *Client) Create(setting *models.CommandSetting) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutInSeconds)*time.Second)
	defer cancel()

	b, err := json.Marshal(setting)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse repository")
		return err
	}

	u, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse address")
		return err
	}

	u.Path = path.Join(u.Path, settingURI)
	code, err := c.Handler.MakeRequest(ctx, http.MethodGet, b, u.String(), nil)
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

// List settings.
func (c *Client) List(id int) ([]*models.CommandSetting, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutInSeconds)*time.Second)
	defer cancel()

	u, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse address")
		return nil, err
	}

	var result []*models.CommandSetting
	u.Path = path.Join(u.Path, settingsURI, strconv.Itoa(id), "settings")
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
