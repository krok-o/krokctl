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
func (p *Handler) MakeRequest(ctx context.Context, method string, url string, opts ...MakeRequestOptions) (int, error) {
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
func (p *Handler) prepare(ctx context.Context, method, url string, payload io.Reader, parseTo interface{}, contentType string) (int, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, payload)
	if err != nil {
		p.Logger.Error().Err(err).Msg("Failed to create HTTP request.")
		return http.StatusInternalServerError, err
	}
	req.Header.Add("Content-Type", contentType)
	if p.Token != "" {
		req.Header.Add("Authorization", "Bearer "+p.Token)
	}

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
		if err := p.parseBody(resp.Body, parseTo); err != nil {
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
