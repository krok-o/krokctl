package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/krok-o/krok/pkg/models"
	"github.com/rs/zerolog"
)

const getTokenURI = "/rest/api/1/get-token"

// Handler makes requests to krok server.
type Handler interface {
	MakeRequest(ctx context.Context, method string, url string, opts ...MakeRequestOptions) (int, error)
}

// Config contains configuration entries for the handler.
type Config struct {
	Client       *http.Client
	Address      string
	APIKeyID     string
	APIKeySecret string
	Email        string
	Logger       zerolog.Logger
}

// NewHandler creates a new handler with a given client.
func NewHandler(cfg Config) *KrokHandler {
	return &KrokHandler{
		Config: cfg,
	}
}

// KrokHandler has methods which can deal with talking to REST endpoints.
type KrokHandler struct {
	Config

	tokenCache string
}

// MakeRequestOption defines options for MakeRequest call.
type MakeRequestOption struct {
	data        []byte
	output      interface{}
	contentType string
}

// MakeRequestOptions defines functional options for optional parameters.
type MakeRequestOptions func(option *MakeRequestOption)

// WithPayload adds a payload options for MakeRequest to send if it's a POST request for example.
func WithPayload(payload []byte) MakeRequestOptions {
	return func(option *MakeRequestOption) {
		option.data = payload
	}
}

// WithOutput adds an output if there is something to get back from MakeRequest like a Get call.
func WithOutput(out interface{}) MakeRequestOptions {
	return func(option *MakeRequestOption) {
		option.output = out
	}
}

// WithContentType adds a specific content type to the request. By default it's application/json.
func WithContentType(ct string) MakeRequestOptions {
	return func(option *MakeRequestOption) {
		option.contentType = ct
	}
}

// MakeRequest sends a request to the designated URL.
// @data - optional data to send along if it is a POST request.
// @url - defines the destination.
// @output - optional output if the body contains a request to parse.
func (p *KrokHandler) MakeRequest(ctx context.Context, method string, url string, opts ...MakeRequestOptions) (int, error) {
	mos := &MakeRequestOption{
		data:        nil,
		output:      nil,
		contentType: "application/json",
	}

	for _, o := range opts {
		o(mos)
	}
	payload := bytes.NewReader(mos.data)
	return p.prepare(ctx, method, url, payload, mos.output, mos.contentType)
}

// prepare the request. Any possible result will be put into the parseTo variable.
func (p *KrokHandler) prepare(ctx context.Context, method, url string, payload io.Reader, parseTo interface{}, contentType string) (int, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, payload)
	if err != nil {
		p.Logger.Error().Err(err).Msg("Failed to create HTTP request.")
		return http.StatusInternalServerError, err
	}
	req.Header.Add("Content-Type", contentType)
	token := p.tokenCache
	if token == "" {
		if token, err = p.authenticate(); err != nil {
			return http.StatusInternalServerError, err
		}
		// any subsequent calls to the api during this run instance should come from a cached
		// record instead of constantly calling out to authenticate.
		p.tokenCache = token
	}
	req.Header.Add("Authorization", "Bearer "+token)
	response, err := p.Send(req, parseTo)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return response.StatusCode, nil
}

// authenticate call the API to get token.
func (p *KrokHandler) authenticate() (string, error) {
	u, err := url.Parse(p.Address)
	if err != nil {
		return "", err
	}
	u.Path = path.Join(u.Path, getTokenURI)
	apiKeyRequest := models.APIKeyAuthRequest{
		Email:        p.Email,
		APIKeyID:     p.APIKeyID,
		APIKeySecret: p.APIKeySecret,
	}
	b, err := json.Marshal(apiKeyRequest)
	if err != nil {
		p.Logger.Debug().Err(err).Msg("Failed to parse repository")
		return "", err
	}
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, u.String(), bytes.NewReader(b))
	if err != nil {
		p.Logger.Error().Err(err).Msg("Failed to create HTTP request.")
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	var result models.TokenResponse
	resp, err := p.Send(req, &result)
	if err != nil {
		p.Logger.Debug().Err(err).Int("code", resp.StatusCode).Msg("failed to authenticate")
		return "", err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		p.Logger.Debug().Err(err).Int("code", resp.StatusCode).Msg("failed to authenticate")
		return "", errors.New("failed to authenticate")
	}
	return result.Token, nil
}

// Send extracts the common operation to send over the wire.
func (p *KrokHandler) Send(req *http.Request, parseTo interface{}) (*http.Response, error) {
	// Send the request
	resp, err := p.Client.Do(req)
	if err != nil {
		p.Logger.Error().Err(err).Msg("Failed to call endpoint.")
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			p.Logger.Debug().Err(err).Msg("Failed to close response body reader.")
		}
	}()
	if parseTo != nil {
		if err := p.parseBody(resp.Body, parseTo); err != nil {
			return nil, err
		}
	}
	return resp, nil
}

// parseBody is a convenient wrapper around a common set of logical operations to get something out of
// the return body of an http response.
func (p *KrokHandler) parseBody(respBody io.Reader, v interface{}) error {
	if err := json.NewDecoder(respBody).Decode(&v); err != nil {
		p.Logger.Error().Err(err).Msg("Failed to unmarshal body.")
		return err
	}
	return nil
}
