package user

import (
	"context"
	"encoding/json"
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
	userURI          = "/rest/api/1/krok/user"
	usersURI         = "/rest/api/1/krok/users"
)

// NewClient creates a new user provider.
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

// Create creates a user resource.
func (c *Client) Create(user *models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutInSeconds)*time.Second)
	defer cancel()

	b, err := json.Marshal(user)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse user")
		return nil, err
	}

	u, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse address")
		return nil, err
	}

	result := models.User{}
	u.Path = path.Join(u.Path, userURI)
	code, err := c.Handler.MakeRequest(ctx, http.MethodPost, u.String(), clients.WithPayload(b), clients.WithOutput(&result))
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

// Get returns a user resource.
func (c *Client) Get(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutInSeconds)*time.Second)
	defer cancel()

	u, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse address")
		return nil, err
	}

	result := models.User{}
	u.Path = path.Join(u.Path, userURI, strconv.Itoa(id))
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

// Update updates a user resource.
func (c *Client) Update(user *models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutInSeconds)*time.Second)
	defer cancel()

	b, err := json.Marshal(user)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse user")
		return nil, err
	}

	u, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse address")
		return nil, err
	}

	result := models.User{}
	u.Path = path.Join(u.Path, userURI, "update")
	code, err := c.Handler.MakeRequest(ctx, http.MethodPost, u.String(), clients.WithPayload(b), clients.WithOutput(&result))
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

// Delete deletes a user resource.
func (c *Client) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutInSeconds)*time.Second)
	defer cancel()

	u, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse address")
		return err
	}
	u.Path = path.Join(u.Path, userURI, strconv.Itoa(id))
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

// List users.
func (c *Client) List() ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutInSeconds)*time.Second)
	defer cancel()

	u, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse address")
		return nil, err
	}

	var result []*models.User
	u.Path = path.Join(u.Path, usersURI)
	code, err := c.Handler.MakeRequest(ctx, http.MethodPost, u.String(), clients.WithOutput(&result))
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

// Generate creates a new login token for the current user.
func (c *Client) Generate() (map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutInSeconds)*time.Second)
	defer cancel()

	u, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse address")
		return nil, err
	}

	var result map[string]string
	u.Path = path.Join(u.Path, userURI, "token", "generate")
	code, err := c.Handler.MakeRequest(ctx, http.MethodPost, u.String(), clients.WithOutput(&result))
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
