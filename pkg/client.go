package pkg

import (
	"net/http"

	"github.com/rs/zerolog"

	"github.com/krok-o/krokctl/pkg/clients/repository"
)

// Config defines configuration for the Krok server client.
type Config struct {
	Address string
}

// KrokClient is the main client for the Krok server.
type KrokClient struct {
	RepositoryClient *repository.Client
}

// NewKrokClient creates a new Krok server client.
func NewKrokClient(cfg Config, log zerolog.Logger) *KrokClient {
	repoClient := repository.NewClient(cfg.Address, &http.Client{}, log)
	return &KrokClient{
		RepositoryClient: repoClient,
	}
}
