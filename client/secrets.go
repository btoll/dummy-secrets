// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "dummy-secrets": Secrets Resource Client
//
// Command:
// $ goagen
// --design=github.com/btoll/dummy-secrets/design
// --out=$(GOPATH)/src/github.com/btoll/dummy-secrets
// --version=v1.3.0

package client

import (
	"bytes"
	"context"
	"fmt"
	uuid "github.com/goadesign/goa/uuid"
	"net/http"
	"net/url"
)

// CreateSecretsPath computes a request path to the create action of Secrets.
func CreateSecretsPath() string {

	return fmt.Sprintf("/v1/secrets/")
}

// Store a new Secret
func (c *Client) CreateSecrets(ctx context.Context, path string, payload *SecretPayload) (*http.Response, error) {
	req, err := c.NewCreateSecretsRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateSecretsRequest create the request corresponding to the create action endpoint of the Secrets resource.
func (c *Client) NewCreateSecretsRequest(ctx context.Context, path string, payload *SecretPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return req, nil
}

// ShowSecretsPath computes a request path to the show action of Secrets.
func ShowSecretsPath(id uuid.UUID) string {
	param0 := id.String()

	return fmt.Sprintf("/v1/secrets/%s", param0)
}

// Get a Secret by its ID
func (c *Client) ShowSecrets(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowSecretsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowSecretsRequest create the request corresponding to the show action endpoint of the Secrets resource.
func (c *Client) NewShowSecretsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
