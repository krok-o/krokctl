package pkg

import (
	"net/http"

	"github.com/rs/zerolog"

	"github.com/krok-o/krokctl/pkg/clients"
	"github.com/krok-o/krokctl/pkg/clients/command"
	"github.com/krok-o/krokctl/pkg/clients/event"
	"github.com/krok-o/krokctl/pkg/clients/platform"
	"github.com/krok-o/krokctl/pkg/clients/repository"
	"github.com/krok-o/krokctl/pkg/clients/runs"
	"github.com/krok-o/krokctl/pkg/clients/setting"
	"github.com/krok-o/krokctl/pkg/clients/user"
	"github.com/krok-o/krokctl/pkg/clients/vault"
	"github.com/krok-o/krokctl/pkg/clients/vcs"
)

// Config defines configuration for the Krok server client.
type Config struct {
	Address      string
	APIKeyID     string
	APIKeySecret string
	Email        string
}

// KrokClient is the main client for the Krok server.
type KrokClient struct {
	RepositoryClient *repository.Client
	CommandClient    *command.Client
	CommandRunClient *runs.Client
	VcsClient        *vcs.Client
	PlatformClient   *platform.Client
	SettingsClient   *setting.Client
	EventClient      *event.Client
	VaultClient      *vault.Client
	UserClient       *user.Client
}

// NewKrokClient creates a new Krok server client.
func NewKrokClient(cfg Config, log zerolog.Logger) *KrokClient {
	handler := clients.NewHandler(clients.Config{
		Client:       &http.Client{},
		Address:      cfg.Address,
		APIKeyID:     cfg.APIKeyID,
		APIKeySecret: cfg.APIKeySecret,
		Email:        cfg.Email,
		Logger:       log,
	})
	repoClient := repository.NewClient(cfg.Address, log, handler)
	commandClient := command.NewClient(cfg.Address, log, handler)
	vcsClient := vcs.NewClient(cfg.Address, log, handler)
	platformClient := platform.NewClient(cfg.Address, log, handler)
	settingsClient := setting.NewClient(cfg.Address, log, handler)
	eventsClient := event.NewClient(cfg.Address, log, handler)
	vaultClient := vault.NewClient(cfg.Address, log, handler)
	commandRunClient := runs.NewClient(cfg.Address, log, handler)
	userClient := user.NewClient(cfg.Address, log, handler)
	return &KrokClient{
		RepositoryClient: repoClient,
		VcsClient:        vcsClient,
		CommandClient:    commandClient,
		CommandRunClient: commandRunClient,
		PlatformClient:   platformClient,
		SettingsClient:   settingsClient,
		EventClient:      eventsClient,
		VaultClient:      vaultClient,
		UserClient:       userClient,
	}
}
