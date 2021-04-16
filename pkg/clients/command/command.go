package command

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
	commandURI       = "/rest/api/1/krok/command"
	commandsURI      = "/rest/api/1/krok/commands"
)

// NewClient creates a new command provider.
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

// Update updates a command resource.
func (c *Client) Update(repo *models.Command) (*models.Command, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutInSeconds)*time.Second)
	defer cancel()

	b, err := json.Marshal(repo)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse command")
		return nil, err
	}

	u, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse address")
		return nil, err
	}

	result := models.Command{}
	u.Path = path.Join(u.Path, commandURI, "update")
	code, err := c.Handler.Post(ctx, b, u.String(), &result)
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

// Delete deletes a command resource.
func (c *Client) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutInSeconds)*time.Second)
	defer cancel()

	u, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse address")
		return err
	}
	u.Path = path.Join(u.Path, commandURI, strconv.Itoa(id))
	code, err := c.Handler.Delete(ctx, u.String())
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

// List repositories.
func (c *Client) List(opts *models.ListOptions) ([]*models.Command, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutInSeconds)*time.Second)
	defer cancel()

	b, err := json.Marshal(opts)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse options")
		return nil, err
	}

	u, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse address")
		return nil, err
	}

	var result []*models.Command
	u.Path = path.Join(u.Path, commandsURI)
	code, err := c.Handler.Post(ctx, b, u.String(), &result)
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

// Get returns a command resource.
func (c *Client) Get(id int) (*models.Command, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutInSeconds)*time.Second)
	defer cancel()

	u, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse address")
		return nil, err
	}

	result := models.Command{}
	u.Path = path.Join(u.Path, commandURI, strconv.Itoa(id))
	code, err := c.Handler.Get(ctx, u.String(), &result)
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

// AddRelationshipToRepository adds a relationship to a repository.
func (c *Client) AddRelationshipToRepository(commandID int, repositoryID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutInSeconds)*time.Second)
	defer cancel()

	u, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse address")
		return fmt.Errorf("failed to parse address: %w", err)
	}

	u.Path = path.Join(u.Path, commandURI, "add-command-rel-for-repository", strconv.Itoa(commandID), strconv.Itoa(repositoryID))
	code, err := c.Handler.Post(ctx, nil, u.String(), nil)
	if err != nil {
		c.Logger.Debug().Err(err).Int("code", code).Msg("Failed to get result.")
		return fmt.Errorf("failed to call POST handler: %w", err)
	}
	if code > 299 || code < 200 {
		c.Logger.Error().Str("url", u.String()).Int("code", code).Msg("Return code was not OK")
		return fmt.Errorf("return code was not OK %d", code)
	}
	return nil
}

// RemoveRelationshipToRepository adds a relationship to a repository.
func (c *Client) RemoveRelationshipToRepository(commandID int, repositoryID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutInSeconds)*time.Second)
	defer cancel()

	u, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse address")
		return fmt.Errorf("failed to parse address: %w", err)
	}

	u.Path = path.Join(u.Path, commandURI, "remove-command-rel-for-repository", strconv.Itoa(commandID), strconv.Itoa(repositoryID))
	code, err := c.Handler.Post(ctx, nil, u.String(), nil)
	if err != nil {
		c.Logger.Debug().Err(err).Int("code", code).Msg("Failed to get result.")
		return fmt.Errorf("failed to call POST handler: %w", err)
	}
	if code > 299 || code < 200 {
		c.Logger.Error().Str("url", u.String()).Int("code", code).Msg("Return code was not OK")
		return fmt.Errorf("return code was not OK %d", code)
	}
	return nil
}

// RemoveRelationshipToPlatform adds a relationship to a platform.
func (c *Client) RemoveRelationshipToPlatform(commandID int, platformID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutInSeconds)*time.Second)
	defer cancel()

	u, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse address")
		return fmt.Errorf("failed to parse address: %w", err)
	}

	u.Path = path.Join(u.Path, commandURI, "remove-command-rel-for-platform", strconv.Itoa(commandID), strconv.Itoa(platformID))
	code, err := c.Handler.Post(ctx, nil, u.String(), nil)
	if err != nil {
		c.Logger.Debug().Err(err).Int("code", code).Msg("Failed to get result.")
		return fmt.Errorf("failed to call POST handler: %w", err)
	}
	if code > 299 || code < 200 {
		c.Logger.Error().Str("url", u.String()).Int("code", code).Msg("Return code was not OK")
		return fmt.Errorf("return code was not OK %d", code)
	}
	return nil
}

// AddRelationshipToPlatform adds a relationship to a platform.
func (c *Client) AddRelationshipToPlatform(commandID int, platformID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutInSeconds)*time.Second)
	defer cancel()

	u, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.Debug().Err(err).Msg("Failed to parse address")
		return fmt.Errorf("failed to parse address: %w", err)
	}

	u.Path = path.Join(u.Path, commandURI, "add-command-rel-for-platform", strconv.Itoa(commandID), strconv.Itoa(platformID))
	code, err := c.Handler.Post(ctx, nil, u.String(), nil)
	if err != nil {
		c.Logger.Debug().Err(err).Int("code", code).Msg("Failed to get result.")
		return fmt.Errorf("failed to call POST handler: %w", err)
	}
	if code > 299 || code < 200 {
		c.Logger.Error().Str("url", u.String()).Int("code", code).Msg("Return code was not OK")
		return fmt.Errorf("return code was not OK %d", code)
	}
	return nil
}
