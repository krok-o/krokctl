package runs

import (
	"context"
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
	runURI           = "/rest/api/1/krok/command/run"
)

// NewClient creates a new command run provider.
func NewClient(address string, log zerolog.Logger, handler clients.Handler) *Client {
	return &Client{
		Address: address,
		Logger:  log,
		Handler: handler,
	}
}

// Client contains methods for command run related resource actions.
type Client struct {
	Address string
	Logger  zerolog.Logger
	Handler clients.Handler
}

// Get returns a command resource.
func (c *Client) Get(id int) (*models.CommandRun, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutInSeconds)*time.Second)
	defer cancel()

	u, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse address")
		return nil, err
	}

	result := models.CommandRun{}
	u.Path = path.Join(u.Path, runURI, strconv.Itoa(id))
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
