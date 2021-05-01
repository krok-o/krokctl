package pkg

import (
	"net/http"

	"github.com/krok-o/krokctl/pkg/clients/event"
	"github.com/krok-o/krokctl/pkg/clients/vault"
	"github.com/rs/zerolog"

	"github.com/krok-o/krokctl/pkg/clients/command"
	"github.com/krok-o/krokctl/pkg/clients/platform"
	"github.com/krok-o/krokctl/pkg/clients/repository"
	"github.com/krok-o/krokctl/pkg/clients/setting"
	"github.com/krok-o/krokctl/pkg/clients/vcs"
)

// Config defines configuration for the Krok server client.
type Config struct {
	Address string
	Token   string
}

// KrokClient is the main client for the Krok server.
type KrokClient struct {
	RepositoryClient *repository.Client
	CommandClient    *command.Client
	VcsClient        *vcs.Client
	PlatformClient   *platform.Client
	SettingsClient   *setting.Client
	EventClient      *event.Client
	VaultClient      *vault.Client
	Token            string
}

// NewKrokClient creates a new Krok server client.
func NewKrokClient(cfg Config, log zerolog.Logger) *KrokClient {
	repoClient := repository.NewClient(cfg.Address, &http.Client{}, cfg.Token, log)
	commandClient := command.NewClient(cfg.Address, &http.Client{}, cfg.Token, log)
	vcsClient := vcs.NewClient(cfg.Address, &http.Client{}, cfg.Token, log)
	platformClient := platform.NewClient(cfg.Address, &http.Client{}, "", log)
	settingsClient := setting.NewClient(cfg.Address, &http.Client{}, cfg.Token, log)
	eventsClient := event.NewClient(cfg.Address, &http.Client{}, cfg.Token, log)
	vaultClient := vault.NewClient(cfg.Address, &http.Client{}, cfg.Token, log)
	return &KrokClient{
		RepositoryClient: repoClient,
		VcsClient:        vcsClient,
		CommandClient:    commandClient,
		PlatformClient:   platformClient,
		SettingsClient:   settingsClient,
		EventClient:      eventsClient,
		VaultClient:      vaultClient,
		Token:            cfg.Token,
	}
}
