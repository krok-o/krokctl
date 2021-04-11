package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog"
)

// NewHandler creates a new handler with a given client.
func NewHandler(client http.Client, token string, log zerolog.Logger) *Handler {
	return &Handler{
		Logger: log,
		Client: client,
		Token:  token,
	}
}

// Handler has methods which can deal with talking to REST endpoints.
type Handler struct {
	Logger zerolog.Logger
	Client http.Client
	Token  string
}

// Post extracts operation regarding posting actions.
func (p *Handler) Post(ctx context.Context, data []byte, url string, a interface{}) (int, error) {
	// Create the request
	payload := bytes.NewReader(data)
	return p.prepare(ctx, "POST", url, payload, a)

}

// Delete extracts operation regarding delete actions.
func (p *Handler) Delete(ctx context.Context, url string) (int, error) {
	return p.prepare(ctx, "DELETE", url, nil, nil)
}

// Get extracts operation regarding get actions.
func (p *Handler) Get(ctx context.Context, url string, a interface{}) (int, error) {
	return p.prepare(ctx, "GET", url, nil, a)
}

// prepare the request. Any possible result will be put into the parseTo variable.
func (p *Handler) prepare(ctx context.Context, method, url string, payload io.Reader, parseTo interface{}) (int, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, payload)
	if err != nil {
		p.Logger.Error().Err(err).Msg("Failed to create HTTP request.")
		return http.StatusInternalServerError, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+p.Token)

	response, err := p.Send(req, parseTo)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return response.StatusCode, nil
}

// Send extracts the common operation to send over the wire.
func (p *Handler) Send(req *http.Request, parseTo interface{}) (*http.Response, error) {
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
		if err := p.parseBody(resp.Body, &parseTo); err != nil {
			return nil, err
		}
	}
	return resp, nil
}

// parseBody is a convenient wrapper around a common set of logical operations to get something out of
// the return body of an http response.
func (p *Handler) parseBody(respBody io.Reader, v interface{}) error {
	if err := json.NewDecoder(respBody).Decode(&v); err != nil {
		p.Logger.Error().Err(err).Msg("Failed to unmarshal body.")
		return err
	}
	return nil
}
